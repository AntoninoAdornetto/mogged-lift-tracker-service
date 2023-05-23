package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateSession(t *testing.T) {
	user := GenRandUser(t)
	GenRandSession(t, user.ID)
}

func TestGetSession(t *testing.T) {
	user := GenRandUser(t)
	GenRandSession(t, user.ID)
}

func TestDeleteSession(t *testing.T) {
	user := GenRandUser(t)
	GenRandSession(t, user.ID)

	err := testQueries.DeleteSession(context.Background(), user.ID)
	require.NoError(t, err)

	session, err := testQueries.GetSession(context.Background(), user.ID)
	require.Error(t, err)
	require.Empty(t, session)
}

// asserts both CREATE & GET queries
func GenRandSession(t *testing.T, userID string) Session {
	args := CreateSessionParams{
		RefreshToken: util.RandomStr(20),
		UserID:       userID,
		ClientIp: fmt.Sprintf("%d.%d.%d.%d",
			util.RandomInt(100, 192),
			util.RandomInt(100, 170),
			util.RandomInt(1, 10),
			util.RandomInt(1, 12)),
		UserAgent: util.RandomStr(10),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	err := testQueries.CreateSession(context.Background(), args)
	require.NoError(t, err)

	session, err := testQueries.GetSession(context.Background(), userID)
	require.NoError(t, err)

	parsedUUID, err := uuid.FromBytes(session.UserID)
	require.NoError(t, err)
	require.Equal(t, args.RefreshToken, session.RefreshToken)
	require.Equal(t, args.UserID, parsedUUID.String())
	require.Equal(t, args.ClientIp, session.ClientIp)
	require.Equal(t, args.UserAgent, session.UserAgent)
	require.WithinDuration(t, args.ExpiresAt, session.ExpiresAt, time.Minute)

	return session
}

package db

import (
	"context"
	"testing"
	"time"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/stretchr/testify/require"
)

func TestCreateSession(t *testing.T) {
	user := GenRandUser(t)

	args := CreateSessionParams{
		RefreshToken: util.RandomStr(20),
		UserID:       user.ID,
		ClientIp:     "0:0:0:0:8080",
		UserAgent:    "Chrome",
		ExpiresAt:    time.Now().Add(time.Hour * 24),
	}

	err := testQueries.CreateSession(context.Background(), args)
	require.NoError(t, err)
}

func TestGetSession(t *testing.T) {
	user := GenRandUser(t)
}

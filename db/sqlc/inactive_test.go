package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueueUserDeletion(t *testing.T) {
	user := GenRandUser(t)

	err := testQueries.InsertInactiveUser(context.Background(), user.ID)
	require.NoError(t, err)

	inactiveUserID, err := testQueries.GetInactiveUser(context.Background(), user.ID)
	require.NoError(t, err)

	require.Equal(t, user.ID, inactiveUserID)
}

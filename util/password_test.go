package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	n := 10
	password := RandomStr(int64(n))
	require.NotEmpty(t, password)
	require.Len(t, password, n)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)

	err = ValidatePassword(password, hashedPassword)
	require.NoError(t, err)
}

func TestIncorrectHashedPassword(t *testing.T) {
	n := 10
	password := RandomStr(int64(n))
	require.NotEmpty(t, password)
	require.Len(t, password, n)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)

	incorrectPassword := password + "incorrect"
	err = ValidatePassword(incorrectPassword, hashedPassword)
	require.Error(t, err)
}

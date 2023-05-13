package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	n := 10
	password := RandomStr(int64(n))
	require.NotEmpty(t, password)
	require.Len(t, password, n)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(incorrectPassword))
	require.Error(t, err)
}

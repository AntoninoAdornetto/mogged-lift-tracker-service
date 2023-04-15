package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	GenRandUser(t)
}

func TestUpdateUserFirstName(t *testing.T) {
	user := GenRandUser(t)
	newFirstName := util.RandomStr(5)

	_, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		FirstName: newFirstName,
		UUIDTOBIN: user.ID,
	})
	require.NoError(t, err)

	updatedUser, err := testQueries.GetUser(context.Background(), user.EmailAddress)
	require.NoError(t, err)
	require.NotNil(t, updateUser)
	require.NotEqual(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, updatedUser.FirstName, newFirstName)
	require.Equal(t, user.EmailAddress, updatedUser.EmailAddress)
	require.Equal(t, user.ID, updatedUser.ID)
}

func TestUpdateUserLastName(t *testing.T) {
	user := GenRandUser(t)
	newLastName := util.RandomStr(5)

	_, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		LastName:  newLastName,
		UUIDTOBIN: user.ID,
	})
	require.NoError(t, err)

	updatedUser, err := testQueries.GetUser(context.Background(), user.EmailAddress)
	require.NoError(t, err)
	require.NotNil(t, updateUser)
	require.NotEqual(t, user.LastName, updatedUser.LastName)
	require.Equal(t, updatedUser.LastName, newLastName)
	require.Equal(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, user.EmailAddress, updatedUser.EmailAddress)
	require.Equal(t, user.ID, updatedUser.ID)
}

func TestUpdateEmail(t *testing.T) {
	user := GenRandUser(t)
	newEmailAddress := util.RandomStr(5) + "@gmail.com"

	_, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		EmailAddress: newEmailAddress,
		UUIDTOBIN:    user.ID,
	})
	require.NoError(t, err)

	updatedUser, err := testQueries.GetUser(context.Background(), newEmailAddress)
	require.NoError(t, err)
	require.NotNil(t, updateUser)
	require.NotEqual(t, user.EmailAddress, updatedUser.EmailAddress)
	require.Equal(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, user.LastName, updatedUser.LastName)
	require.Equal(t, user.ID, updatedUser.ID)
	require.Equal(t, updatedUser.EmailAddress, newEmailAddress)
}

func TestDeleteUser(t *testing.T) {
	user := GenRandUser(t)

	_, err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	_, err = testQueries.GetUser(context.Background(), user.EmailAddress)
	require.Error(t, err)
}

// tests both Create and Get
func GenRandUser(t *testing.T) GetUserRow {
	args := CreateUserParams{
		FirstName:    util.RandomStr(10),
		LastName:     util.RandomStr(10),
		Password:     util.RandomStr(20),
		EmailAddress: util.RandomStr(7) + "@gmail.com",
	}

	insert, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotNil(t, insert)

	user, err := testQueries.GetUser(context.Background(), args.EmailAddress)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, args.FirstName, user.FirstName)
	require.Equal(t, args.LastName, user.LastName)
	require.Equal(t, args.EmailAddress, user.EmailAddress)
	require.NotNil(t, user.ID)

	return user
}

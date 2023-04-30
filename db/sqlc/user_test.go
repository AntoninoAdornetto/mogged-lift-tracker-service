package db

import (
	"context"
	"testing"
	"time"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	GenRandUser(t)
}

func TestGetUser(t *testing.T) {
	user := GenRandUser(t)
	query, err := testQueries.GetUser(context.Background(), user.EmailAddress)
	require.NoError(t, err)
	require.Equal(t, user.EmailAddress, query.EmailAddress)
	require.Equal(t, user.Password, query.Password)
	require.Equal(t, user.FirstName, query.FirstName)
	require.Equal(t, user.LastName, query.LastName)
	require.Equal(t, user.ID, query.ID)
	require.Equal(t, user.PasswordChangedAt, query.PasswordChangedAt)
}

func TestUpdateUserFirstName(t *testing.T) {
	user := GenRandUser(t)
	newFirstName := util.RandomStr(5)

	err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		FirstName: newFirstName,
		UserID:    user.ID,
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

	err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		LastName: newLastName,
		UserID:   user.ID,
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

	err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		EmailAddress: newEmailAddress,
		UserID:       user.ID,
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

func TestChangePassword(t *testing.T) {
	user := GenRandUser(t)
	require.Equal(t, user.PasswordChangedAt.Year(), 1970) // indicates password has never been changed

	newPassword := util.RandomStr(12)
	err := testQueries.ChangePassword(context.Background(), ChangePasswordParams{
		Password: newPassword,
		UserID:   user.ID,
	})

	updatedUser, err := testQueries.GetUser(context.Background(), user.EmailAddress)
	require.NoError(t, err)
	require.WithinDuration(t, updatedUser.PasswordChangedAt, time.Now(), time.Minute)
	require.NoError(t, err)
	require.NotEqual(t, user.Password, updatedUser.Password)
}

func TestDeleteUser(t *testing.T) {
	user := GenRandUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	_, err = testQueries.GetUser(context.Background(), user.EmailAddress)
	require.Error(t, err)
}

func GenRandUser(t *testing.T) GetUserRow {
	args := CreateUserParams{
		FirstName:    util.RandomStr(10),
		LastName:     util.RandomStr(10),
		Password:     util.RandomStr(20),
		EmailAddress: util.RandomStr(7) + "@gmail.com",
	}

	err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)

	user, err := testQueries.GetUser(context.Background(), args.EmailAddress)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, args.FirstName, user.FirstName)
	require.Equal(t, args.LastName, user.LastName)
	require.Equal(t, args.EmailAddress, user.EmailAddress)
	require.Equal(t, user.PasswordChangedAt.Year(), 1970) // indicates the password has never been changed
	require.GreaterOrEqual(t, len(args.Password), 10)
	require.NotNil(t, user.ID)

	return user
}

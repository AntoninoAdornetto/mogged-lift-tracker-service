package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	userId := getNewUserId(t)
	GenRanProfile(t, userId)
}

func TestUpdateCountry(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRanProfile(t, userId)

	newCountry := util.RandomStr(3)
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		Country: newCountry,
		UserID:  userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newCountry, profile.Country)
	require.Equal(t, query.Country, newCountry)
}

func TestUpdateBodyFat(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRanProfile(t, userId)

	newBodyFat := float64(util.RandomInt(10, 20))
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		BodyFat: newBodyFat,
		UserID:  userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newBodyFat, profile.BodyFat)
	require.Equal(t, query.BodyFat, newBodyFat)
}

func TestUpdateBodyWeight(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRanProfile(t, userId)

	newBodyWeight := float64(util.RandomInt(10, 20))
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		BodyWeight: newBodyWeight,
		UserID:     userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newBodyWeight, profile.BodyWeight)
	require.Equal(t, query.BodyWeight, newBodyWeight)
}

func TestUpdateTimezone(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRanProfile(t, userId)

	newTimezone := util.RandomStr(10)
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		Timezone: newTimezone,
		UserID:   userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newTimezone, profile.Timezone)
	require.Equal(t, query.Timezone, newTimezone)
}

func TestUpdateMeasurementSystem(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRanProfile(t, userId)

	newMeasurementSystem := util.RandomStr(10)
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		MeasurementSystem: newMeasurementSystem,
		UserID:            userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newMeasurementSystem, profile.MeasurementSystem)
	require.Equal(t, query.MeasurementSystem, newMeasurementSystem)
}

func TestDeleteProfile(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	GenRanProfile(t, userId)

	_, err = testQueries.DeleteProfile(context.Background(), userId.String())
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.Error(t, err)
	require.Zero(t, query.ID)
}

func getNewUserId(t *testing.T) uuid.UUID {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	return userId
}

// asserts both create and get one queries
func GenRanProfile(t *testing.T, userId uuid.UUID) Profile {
	p := &CreateProfileParams{
		Country:           util.RandomStr(3),
		BodyFat:           float64(util.RandomInt(8, 20)),
		BodyWeight:        float64(util.RandomInt(150, 220)),
		Timezone:          util.RandomStr(10),
		MeasurementSystem: "Imperial",
		UserID:            userId.String(),
	}

	_, err := testQueries.CreateProfile(context.Background(), *p)
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.TODO(), userId.String())
	userIDFromBytes, _ := uuid.FromBytes(query.UserID)

	require.NoError(t, err)
	require.NotNil(t, query)
	require.Equal(t, p.Country, query.Country)
	require.Equal(t, p.BodyFat, query.BodyFat)
	require.Equal(t, p.BodyWeight, query.BodyWeight)
	require.Equal(t, p.Timezone, query.Timezone)
	require.Equal(t, p.MeasurementSystem, query.MeasurementSystem)
	require.Equal(t, userId, userIDFromBytes)
	return query
}

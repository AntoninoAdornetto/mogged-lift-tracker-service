package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	userId := getNewUserId(t)
	GenRandProfile(t, userId)
}

func TestUpdateCountry(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRandProfile(t, userId)

	newCountry := util.RandomStr(3)
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		Country: sql.NullString{
			Valid:  true,
			String: newCountry,
		},
		UserID: userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newCountry, profile.Country)
	require.Equal(t, query.Country, newCountry)
}

func TestUpdateBodyFat(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRandProfile(t, userId)

	newBodyFat := float64(util.RandomInt(50, 100))
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		BodyFat: sql.NullFloat64{
			Valid:   true,
			Float64: newBodyFat,
		},
		UserID: userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newBodyFat, profile.BodyFat)
	require.Equal(t, query.BodyFat, newBodyFat)
}

func TestUpdateBodyWeight(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRandProfile(t, userId)

	newBodyWeight := float64(util.RandomInt(300, 500))
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		BodyWeight: sql.NullFloat64{
			Valid:   true,
			Float64: newBodyWeight,
		},
		UserID: userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newBodyWeight, profile.BodyWeight)
	require.Equal(t, query.BodyWeight, newBodyWeight)
}

func TestUpdateTimezoneOffset(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRandProfile(t, userId)

	// invalid range. There are constraints on the column
	invalidTimezoneOffset := []int32{-900, 1000}

	for _, tOffset := range invalidTimezoneOffset {
		_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
			TimezoneOffset: sql.NullInt32{
				Valid: true,
				Int32: tOffset,
			},
		})
		require.Error(t, err)
	}

	newTimezoneOffset := int32(util.RandomInt(-700, 800))
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		TimezoneOffset: sql.NullInt32{
			Valid: true,
			Int32: newTimezoneOffset,
		},
		UserID: userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetProfile(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEqual(t, newTimezoneOffset, profile.TimezoneOffset)
	require.Equal(t, query.TimezoneOffset, newTimezoneOffset)
}

func TestUpdateMeasurementSystem(t *testing.T) {
	userId := getNewUserId(t)
	profile := GenRandProfile(t, userId)

	newMeasurementSystem := util.RandomStr(10)
	_, err := testQueries.UpdateProfile(context.Background(), UpdateProfileParams{
		MeasurementSystem: sql.NullString{
			Valid:  true,
			String: newMeasurementSystem,
		},
		UserID: userId.String(),
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
	GenRandProfile(t, userId)

	err = testQueries.DeleteProfile(context.Background(), userId.String())
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
func GenRandProfile(t *testing.T, userId uuid.UUID) Profile {
	p := &CreateProfileParams{
		Country:           util.RandomStr(3),
		BodyFat:           float64(util.RandomInt(8, 20)),
		BodyWeight:        float64(util.RandomInt(150, 220)),
		TimezoneOffset:    0,
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
	require.Equal(t, p.TimezoneOffset, query.TimezoneOffset)
	require.Equal(t, p.MeasurementSystem, query.MeasurementSystem)
	require.Equal(t, userId, userIDFromBytes)
	return query
}

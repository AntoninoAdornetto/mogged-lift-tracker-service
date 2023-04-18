package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// asserts both create and get one queries
func TestCreateLift(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	GenRandLift(t, userId.String())
}

func GenRandLift(t *testing.T, userId string) Lift {
	lift := &Lift{}
	exercise := GenRandExercise(t, userId)
	workout := GenRandWorkout(t, userId)

	args := CreateLiftParams{
		ExerciseName: exercise.Name,
		Reps:         int32(util.RandomInt(3, 12)),
		WeightLifted: float64(util.RandomInt(100, 300)),
		UserID:       userId,
		WorkoutID:    workout.ID,
	}
	record, err := testQueries.CreateLift(context.Background(), args)
	require.NoError(t, err)
	lastId, err := record.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, lastId)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{UserID: userId, ID: lastId})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ExerciseName, args.ExerciseName)
	require.Equal(t, query.Reps, args.Reps)
	require.Equal(t, query.WeightLifted, args.WeightLifted)
	require.Equal(t, query.WorkoutID, args.WorkoutID)

	queryUserId, err := uuid.FromBytes(query.UserID)
	require.NoError(t, err)
	require.Equal(t, queryUserId.String(), args.UserID)

	lift = &query
	return *lift
}

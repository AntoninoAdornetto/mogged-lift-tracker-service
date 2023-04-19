package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type NewLiftArgs struct {
	UserID    string
	WorkoutID int32
}

// asserts both create and get one queries
func TestCreateLift(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	workout := GenRandWorkout(t, userId.String())
	GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
}

func TestListLiftsFromWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())

	query, err := testQueries.ListLiftsFromWorkout(context.Background(), ListLiftsFromWorkoutParams{
		UserID:    userId.String(),
		WorkoutID: workout.ID,
	})
	require.NoError(t, err)
	require.Empty(t, query)

	n := 5
	lifts := make([]Lift, n)
	for i := 0; i < n; i++ {
		lifts[i] = GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
	}

	query, err = testQueries.ListLiftsFromWorkout(context.Background(), ListLiftsFromWorkoutParams{
		UserID:    userId.String(),
		WorkoutID: workout.ID,
	})
	require.NoError(t, err)
	require.Len(t, query, n)
}

func TestListMaxWeightPrs(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	n := 5
	workout := GenRandWorkout(t, userId.String())

	query, err := testQueries.ListMaxWeightPrs(context.Background(), ListMaxWeightPrsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Empty(t, query)

	for i := 0; i < n; i++ {
		GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
	}

	query, err = testQueries.ListMaxWeightPrs(context.Background(), ListMaxWeightPrsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Len(t, query, n)

	weight := query[0].WeightLifted
	for i := 1; i < n; i++ {
		require.LessOrEqual(t, query[i].WeightLifted, weight)
		weight = query[i].WeightLifted
	}
}

func TestListMaxRepPrs(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	n := 5
	workout := GenRandWorkout(t, userId.String())

	query, err := testQueries.ListMaxRepPrs(context.Background(), ListMaxRepPrsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Empty(t, query)

	for i := 0; i < n; i++ {
		GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
	}

	query, err = testQueries.ListMaxRepPrs(context.Background(), ListMaxRepPrsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Len(t, query, n)

	reps := query[0].Reps
	for i := 1; i < n; i++ {
		require.LessOrEqual(t, query[i].Reps, reps)
		reps = query[i].Reps
	}
}

func GenRandLift(t *testing.T, args NewLiftArgs) Lift {
	lift := &Lift{}
	exercise := GenRandExercise(t, args.UserID)

	createLiftArgs := CreateLiftParams{
		ExerciseName: exercise.Name,
		Reps:         int32(util.RandomInt(3, 12)),
		WeightLifted: float64(util.RandomInt(100, 300)),
		UserID:       args.UserID,
		WorkoutID:    args.WorkoutID,
	}
	record, err := testQueries.CreateLift(context.Background(), createLiftArgs)
	require.NoError(t, err)
	lastId, err := record.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, lastId)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{UserID: args.UserID, ID: lastId})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ExerciseName, createLiftArgs.ExerciseName)
	require.Equal(t, query.Reps, createLiftArgs.Reps)
	require.Equal(t, query.WeightLifted, createLiftArgs.WeightLifted)
	require.Equal(t, query.WorkoutID, args.WorkoutID)

	queryUserId, err := uuid.FromBytes(query.UserID)
	require.NoError(t, err)
	require.Equal(t, queryUserId.String(), args.UserID)

	lift = &query
	return *lift
}

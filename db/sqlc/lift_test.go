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

// @todo
// func TestGetMaxLiftByExercise(t *testing.T) {
// 	testQueries.GetMaxLiftByExercise()
// }

// func TestListMaxWeightPrsByCategory(t *testing.T) {
// 	testQueries.GetMaxLiftsByMuscleGroup(context.Background())
// }

func TestUpdateLift(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())
	lift := GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})

	newExerciseName := GenRandExercise(t, userId.String()).Name
	_, err = testQueries.UpdateLift(context.Background(), UpdateLiftParams{
		ExerciseName: newExerciseName,
		UserID:       userId.String(),
		ID:           lift.ID,
	})
	require.NoError(t, err)
	require.NoError(t, err)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)
	require.Equal(t, query.ID, lift.ID)
	require.Equal(t, query.ExerciseName, newExerciseName)

	newWeightLifted := float64(util.RandomInt(100, 500))
	_, err = testQueries.UpdateLift(context.Background(), UpdateLiftParams{
		WeightLifted: newWeightLifted,
		UserID:       userId.String(),
		ID:           lift.ID,
	})
	require.NoError(t, err)

	query, err = testQueries.GetLift(context.Background(), GetLiftParams{
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)
	require.Equal(t, query.ID, lift.ID)
	require.Equal(t, query.WeightLifted, newWeightLifted)

	newReps := int32(util.RandomInt(100, 500))
	_, err = testQueries.UpdateLift(context.Background(), UpdateLiftParams{
		Reps:   newReps,
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)

	query, err = testQueries.GetLift(context.Background(), GetLiftParams{
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)
	require.Equal(t, query.ID, lift.ID)
	require.Equal(t, query.Reps, newReps)
}

func TestDeleteLift(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())
	lift := GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})

	_, err = testQueries.DeleteLift(context.Background(), DeleteLiftParams{
		ID:     lift.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{
		ID:     lift.ID,
		UserID: userId.String(),
	})
	require.Error(t, err)
	require.Zero(t, query.ID)
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
		SetType:      "Working",
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
	require.Equal(t, query.SetType, createLiftArgs.SetType)

	queryUserId, err := uuid.FromBytes(query.UserID)
	require.NoError(t, err)
	require.Equal(t, queryUserId.String(), args.UserID)

	lift = &query
	return *lift
}

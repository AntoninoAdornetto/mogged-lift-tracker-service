package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// asserts both create and get one queries
func TestCreateExercise(t *testing.T) {
	user := GenRandUser(t)
	GenRandExercise(t, user.ID)
}

func TestListExercises(t *testing.T) {
	n := 5
	user := GenRandUser(t)
	otherUser := GenRandUser(t) // this user should not see the newly created exercises
	exercises := make([]Exercise, n)

	for i := 0; i < n; i++ {
		exercises[i] = GenRandExercise(t, user.ID)
	}

	query, err := testQueries.ListExercises(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, query, n)

	creatorUserId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	otherUserId, err := uuid.Parse(otherUser.ID)
	require.NoError(t, err)

	for _, v := range query {
		queryUserId, err := uuid.FromBytes(v.UserID)
		require.NoError(t, err)
		require.NotEmpty(t, v.ID)
		require.NotEmpty(t, v.Name)
		require.NotEmpty(t, v.Category)
		require.NotEmpty(t, v.MuscleGroup)
		require.NotEmpty(t, v.RestTimer)
		require.False(t, v.Isstock)
		require.Equal(t, queryUserId, creatorUserId)
		require.Zero(t, v.MostRepsLifted)
		require.Zero(t, v.MostWeightLifted)

		// other users will not see the exercises
		require.NotEqual(t, queryUserId, otherUserId)
	}

	otherUserQuery, err := testQueries.ListExercises(context.Background(), otherUser.ID)
	require.NoError(t, err)
	require.Empty(t, otherUserQuery)
}

func TestUpdateExerciseName(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newExerciseName := util.RandomStr(5)

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		Name:   newExerciseName,
		UserID: userId.String(),
		ID:     exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), exercise.ID)
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.Name, newExerciseName)
	require.NotEqual(t, query.Name, exercise.Name)
}

func TestUpdateMuscleGroup(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newMuscleGroup := GenRandMuscleGroup(t).Name

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		MuscleGroup: newMuscleGroup,
		UserID:      userId.String(),
		ID:          exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), exercise.ID)
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.MuscleGroup, newMuscleGroup)
	require.NotEqual(t, query.MuscleGroup, exercise.MuscleGroup)
}

func TestUpdateCategory(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newCategory := GenRandCategory(t).Name

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		Category: newCategory,
		UserID:   userId.String(),
		ID:       exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), exercise.ID)
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.Category, newCategory)
	require.NotEqual(t, query.Category, exercise.Category)
}

func TestUpdateMostWeightLifted(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newMostWeightLifted := float64(util.RandomInt(150, 500))

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		MostWeightLifted: newMostWeightLifted,
		UserID:           userId.String(),
		ID:               exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), exercise.ID)
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.MostWeightLifted, newMostWeightLifted)
	require.NotEqual(t, query.MostWeightLifted, exercise.MostWeightLifted)
}

func GenRandExercise(t *testing.T, userID string) Exercise {
	exercise := &Exercise{}

	exerciseName := util.RandomStr(10)
	category := GenRandCategory(t)
	muscleGroup := GenRandMuscleGroup(t)
	userId, err := uuid.Parse(userID)
	require.NoError(t, err)

	args := CreateExerciseParams{
		Name:        exerciseName,
		Category:    category.Name,
		MuscleGroup: muscleGroup.Name,
		UserID:      userId.String(),
	}

	record, err := testQueries.CreateExercise(context.Background(), args)
	require.NoError(t, err)

	exerciseId, err := record.LastInsertId()
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), int32(exerciseId))
	require.NoError(t, err)
	require.NotEmpty(t, query.ID)

	queryUserId, err := uuid.FromBytes(query.UserID)
	require.NoError(t, err)

	require.Equal(t, query.Name, args.Name)
	require.Equal(t, query.Category, args.Category)
	require.Equal(t, query.MuscleGroup, args.MuscleGroup)
	require.Equal(t, queryUserId, userId)

	exercise = &query
	return *exercise
}

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

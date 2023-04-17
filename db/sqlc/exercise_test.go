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
	GenRandExercise(t)
}

// @todo
// func TestListExercises(t *testing.T) {
// }

func GenRandExercise(t *testing.T) Exercise {
	exercise := &Exercise{}

	exerciseName := util.RandomStr(10)
	category := GenRandCategory(t)
	muscleGroup := GenRandMuscleGroup(t)
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
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

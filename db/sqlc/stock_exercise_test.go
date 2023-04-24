package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/stretchr/testify/require"
)

// asserts both Create & GetOne queries
func TestCreateStockExercise(t *testing.T) {
	GenRandStockExercise(t)
}

func TestListStockExercises(t *testing.T) {
	GenRandStockExercise(t)

	exercises, err := testQueries.ListStockExercies(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(exercises), 1)
}

func TestUpdateStockExercise(t *testing.T) {
	exercise := GenRandStockExercise(t)

	newExerciseName := util.RandomStr(5)
	newCategory := GenRandCategory(t).Name
	newMuscleGroup := GenRandMuscleGroup(t).Name

	_, err := testQueries.UpdateStockExercise(context.Background(), UpdateStockExerciseParams{
		ID: exercise.ID,
		Name: newExerciseName,
		MuscleGroup: newMuscleGroup,
		Category: newCategory,
	})
	require.NoError(t, err)

	query, err := testQueries.GetStockExercise(context.Background(), exercise.ID)
	require.NoError(t, err)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.Category, newCategory)
	require.Equal(t, query.MuscleGroup, newMuscleGroup)
	require.Equal(t, query.Name, newExerciseName)
}

func TestDeleteStockExercise(t *testing.T) {
	exercise := GenRandStockExercise(t)

	_, err := testQueries.DeleteStockExercise(context.Background(), exercise.ID)
	require.NoError(t, err)

	query, err := testQueries.GetStockExercise(context.Background(), exercise.ID)
	require.Error(t, err)
	require.Zero(t, query.ID)
}

func GenRandStockExercise(t *testing.T) StockExercise {
	muscleGroup := GenRandMuscleGroup(t)
	category := GenRandCategory(t)

	args := CreateStockExerciseParams{
		Name: util.RandomStr(5),
		MuscleGroup: muscleGroup.Name,
		Category: category.Name,
	}

	record, err := testQueries.CreateStockExercise(context.Background(), args)
	require.NoError(t, err)

	id, err := record.LastInsertId()
	require.NoError(t, err)

	query, err := testQueries.GetStockExercise(context.Background(), int32(id))
	require.NoError(t, err)

	require.Equal(t, query.ID, int32(id))
	require.Equal(t, query.Category, args.Category)
	require.Equal(t, query.Name, args.Name)
	require.Equal(t, query.MuscleGroup, args.MuscleGroup)
	return query
}
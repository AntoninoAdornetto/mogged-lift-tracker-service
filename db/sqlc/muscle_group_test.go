package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/stretchr/testify/require"
)

func TestCreateMuscleGroup(t *testing.T) {
	GenRandMuscleGroup(t)
}

// Real Muscle Group, not rand generated. uses muscle group from seed migration
func TestGetMuscleGroupByName(t *testing.T) {
	chest := "Chest"
	muscleGroup, err := testQueries.GetMuscleGroupByName(context.Background(), chest)
	require.NoError(t, err)
	require.NotZero(t, muscleGroup.ID)
	require.Equal(t, muscleGroup.Name, chest)
}

func TestListMuscleGroups(t *testing.T) {
	n := 5
	muscleGroups := make([]MuscleGroup, n)

	for i := 0; i < n; i++ {
		muscleGroups[i] = GenRandMuscleGroup(t)
	}

	query, err := testQueries.ListMuscleGroups(context.Background())
	require.NoError(t, err)

	for _, v := range query {
		require.NotEmpty(t, v.ID)
		require.NotEmpty(t, v.Name)
	}
}

func TestUpdateMuscleGroupName(t *testing.T) {
	muscleGroup := GenRandMuscleGroup(t)

	newMuscleGroupName := util.RandomStr(5)

	_, err := testQueries.UpdateMuscleGroup(context.Background(), UpdateMuscleGroupParams{
		ID:   muscleGroup.ID,
		Name: newMuscleGroupName,
	})
	require.NoError(t, err)

	query, err := testQueries.GetMuscleGroup(context.Background(), muscleGroup.ID)
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.NotEqual(t, query.Name, muscleGroup.Name)
	require.Equal(t, query.Name, newMuscleGroupName)
	require.Equal(t, query.ID, muscleGroup.ID)
}

func TestDeleteMuscleGroup(t *testing.T) {
	muscleGroup := GenRandMuscleGroup(t)

	_, err := testQueries.DeleteMuscleGroup(context.Background(), muscleGroup.ID)
	require.NoError(t, err)

	query, err := testQueries.GetMuscleGroup(context.Background(), muscleGroup.ID)
	require.Error(t, err)
	require.Empty(t, query.ID)
	require.Empty(t, query.Name)
}

// asserts both create and get one queries
func GenRandMuscleGroup(t *testing.T) MuscleGroup {
	muscleGroupName := util.RandomStr(5)
	record, err := testQueries.CreateMuscleGroup(context.Background(), muscleGroupName)
	require.NoError(t, err)

	id, err := record.LastInsertId()
	require.NoError(t, err)

	muscleGroup, err := testQueries.GetMuscleGroup(context.Background(), int32(id))
	require.NoError(t, err)
	require.Equal(t, muscleGroup.Name, muscleGroupName)
	require.Equal(t, muscleGroup.ID, int32(id))
	return muscleGroup
}

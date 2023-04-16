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

func TestListMuscleGroups(t *testing.T) {
	n := 5
	muscleGroups := make([]MuscleGroup, n)

	for i := 0; i < n; i++ {
		muscleGroups[i] = GenRandMuscleGroup(t)
	}

	query, err := testQueries.ListMuscleGroups(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(query), n)
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

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

package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateExercise(t *testing.T) {
	user := GenRandUser(t)
	GenRandExercise(t, user.ID)
}

func TestGetExerciseByName(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	exercise := GenRandExercise(t, user.ID)

	query, err := testQueries.GetExerciseByName(context.Background(), GetExerciseByNameParams{
		Name:   exercise.Name,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	queryUserId, err := uuid.FromBytes(query.UserID)
	require.NoError(t, err)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, queryUserId, userId)
	require.Equal(t, query.Name, exercise.Name)
	require.Equal(t, query.Isstock, exercise.Isstock)
	require.Equal(t, query.Category, exercise.Category)
	require.Equal(t, query.RestTimer, exercise.RestTimer)
	require.Equal(t, query.MuscleGroup, exercise.MuscleGroup)
	require.Equal(t, query.MostRepsLifted, exercise.MostRepsLifted)
	require.Equal(t, query.MostWeightLifted, exercise.MostWeightLifted)
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
	newExerciseName := sql.NullString{
		String: util.RandomStr(5),
		Valid:  true,
	}

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		Name:   newExerciseName,
		UserID: userId.String(),
		ID:     exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.Name, newExerciseName.String)
	require.NotEqual(t, query.Name, exercise.Name)
}

func TestUpdateMuscleGroup(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newMuscleGroup := sql.NullString{
		String: GenRandMuscleGroup(t).Name,
		Valid:  true,
	}

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		MuscleGroup: newMuscleGroup,
		UserID:      userId.String(),
		ID:          exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.MuscleGroup, newMuscleGroup.String)
	require.NotEqual(t, query.MuscleGroup, exercise.MuscleGroup)
}

func TestUpdateCategory(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newCategory := sql.NullString{
		String: GenRandCategory(t).Name,
		Valid:  true,
	}

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		Category: newCategory,
		UserID:   userId.String(),
		ID:       exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.Category, newCategory.String)
	require.NotEqual(t, query.Category, exercise.Category)
}

func TestUpdateMostWeightLifted(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newMostWeightLifted := sql.NullFloat64{
		Float64: float64(util.RandomInt(150, 500)),
		Valid:   true,
	}

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		MostWeightLifted: newMostWeightLifted,
		UserID:           userId.String(),
		ID:               exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.MostWeightLifted, newMostWeightLifted.Float64)
	require.NotEqual(t, query.MostWeightLifted, exercise.MostWeightLifted)
}

func TestUpdateMostRepsLifted(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newMostRepsLifted := sql.NullInt32{
		Int32: int32(util.RandomInt(150, 500)),
		Valid: true,
	}

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		MostRepsLifted: newMostRepsLifted,
		UserID:         userId.String(),
		ID:             exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.MostRepsLifted, newMostRepsLifted.Int32)
	require.NotEqual(t, query.MostRepsLifted, exercise.MostRepsLifted)
}

func TestUpdateRestTimer(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())
	newRestTimer := sql.NullString{
		String: "01:30:00s",
		Valid:  true,
	}

	_, err = testQueries.UpdateExercise(context.Background(), UpdateExerciseParams{
		RestTimer: newRestTimer,
		UserID:    userId.String(),
		ID:        exercise.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, exercise.ID)
	require.Equal(t, query.RestTimer, newRestTimer.String)
	require.NotEqual(t, query.RestTimer, exercise.RestTimer)
}

func TestDeleteExercise(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	exercise := GenRandExercise(t, userId.String())

	_, err = testQueries.DeleteExercise(context.Background(), DeleteExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})

	require.NoError(t, err)

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     exercise.ID,
		UserID: userId.String(),
	})
	require.Error(t, err)
	require.Empty(t, query)
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

	query, err := testQueries.GetExercise(context.Background(), GetExerciseParams{
		ID:     int32(exerciseId),
		UserID: userId.String(),
	})
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

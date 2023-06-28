package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type BuildLiftMapParams struct {
	ExerciseName string
	UserID       string
	WorkoutID    int32
	WorkoutMap   map[string][]Lift
}

func TestCreateWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	GenRandWorkout(t, userId.String())
}

func TestGetWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	workoutMap := make(map[string][]Lift)
	workout := GenRandWorkout(t, userId.String())
	err = json.Unmarshal(workout.Lifts, &workoutMap)
	require.NoError(t, err)

	qWorkoutMap := make(map[string][]Lift)
	query, err := testQueries.GetWorkout(context.Background(), GetWorkoutParams{
		ID:     workout.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)

	err = json.Unmarshal(query.Lifts, &qWorkoutMap)
	require.NoError(t, err)

	userIdFromBytes, err := uuid.FromBytes(query.UserID)
	require.NoError(t, err)
	require.Equal(t, query.Duration, workout.Duration)
	require.Equal(t, query.ID, workout.ID)
	require.Equal(t, query.Lifts, workout.Lifts)
	require.Equal(t, userIdFromBytes, userId)

	for k, lifts := range workoutMap {
		for i, lift := range lifts {
			require.Equal(t, lift.ExerciseName, qWorkoutMap[k][i].ExerciseName)
			require.Equal(t, lift.ID, qWorkoutMap[k][i].ID)
			require.Equal(t, lift.Reps, qWorkoutMap[k][i].Reps)
			require.Equal(t, lift.WeightLifted, qWorkoutMap[k][i].WeightLifted)
			require.Equal(t, lift.SetType, qWorkoutMap[k][i].SetType)
			require.Equal(t, lift.WorkoutID, qWorkoutMap[k][i].WorkoutID)
		}
	}
}

func TestListWorkouts(t *testing.T) {
	n := 5
	user := GenRandUser(t)

	workouts := make([]Workout, n)
	for i := range workouts {
		workouts[i] = GenRandWorkout(t, user.ID)
	}

	query, err := testQueries.ListWorkouts(context.Background(), user.ID)
	require.NoError(t, err)

	for i, v := range query {
		require.Equal(t, workouts[i].CompletedDate, v.CompletedDate)
		require.Equal(t, workouts[i].ID, v.ID)
		require.Equal(t, workouts[i].Lifts, v.Lifts)
		require.Equal(t, workouts[i].Duration, v.Duration)
		require.Equal(t, workouts[i].UserID, v.UserID)
	}
}

func TestUpdateWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	workout := GenRandWorkout(t, userId.String())
	newWorkout := GenRandWorkout(t, userId.String())
	newDuration := sql.NullString{
		String: "01:00:0s",
		Valid:  true,
	}

	err = testQueries.UpdateWorkout(context.Background(), UpdateWorkoutParams{
		Duration: newDuration,
		Lifts:    newWorkout.Lifts,
		ID:       workout.ID,
		UserID:   userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetWorkout(context.Background(), GetWorkoutParams{
		ID:     workout.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)
	require.Equal(t, query.Duration, newDuration.String)
	require.Equal(t, query.ID, workout.ID)
	require.Equal(t, query.Lifts, newWorkout.Lifts)
	require.NotEqual(t, query.Lifts, workout.Lifts)
}

func TestDeleteWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())

	err = testQueries.DeleteWorkout(context.Background(), DeleteWorkoutParams{
		ID:     workout.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetWorkout(context.Background(), GetWorkoutParams{
		ID:     workout.ID,
		UserID: userId.String(),
	})
	require.Error(t, err)
	require.Zero(t, query.ID)
}

// using Lift struct but not actually creating an entry into Lift table
// that is what the lift table is for
func BuildLiftsMap(t *testing.T, args BuildLiftMapParams) {
	sets := make([]Lift, util.RandomInt(1, 3))

	for i := range sets {
		sets[i] = Lift{
			ExerciseName: args.ExerciseName,
			WeightLifted: float64(util.RandomInt(100, 200)),
			Reps:         int32(util.RandomInt(6, 12)),
			SetType:      "Working",
		}
	}

	args.WorkoutMap[args.ExerciseName] = sets
}

func GenRandWorkout(t *testing.T, userId string) Workout {
	n := 3
	liftsMap := make(map[string][]Lift)
	exercises := make([]Exercise, n)

	record, err := testQueries.CreateWorkout(context.Background(), CreateWorkoutParams{
		UserID: userId,
	})
	require.NoError(t, err)
	workoutId, err := record.LastInsertId()
	require.NoError(t, err)

	for i := range exercises {
		exercises[i] = GenRandExercise(t, userId)
		BuildLiftsMap(t, BuildLiftMapParams{
			ExerciseName: exercises[i].Name,
			UserID:       userId,
			WorkoutID:    int32(workoutId),
			WorkoutMap:   liftsMap,
		})
	}

	rawJson, err := json.Marshal(liftsMap)
	require.NoError(t, err)
	err = testQueries.UpdateWorkout(context.Background(), UpdateWorkoutParams{
		Lifts:  rawJson,
		ID:     int32(workoutId),
		UserID: userId,
	})
	require.NoError(t, err)

	query, err := testQueries.GetWorkout(context.Background(), GetWorkoutParams{
		ID:     int32(workoutId),
		UserID: userId,
	})
	require.NoError(t, err)
	require.WithinDuration(t, time.Now().UTC(), query.CompletedDate.UTC(), time.Hour*24)

	return query
}

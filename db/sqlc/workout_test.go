package db

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type LiftJSON struct {
	ExerciseName string
	Reps         int32
	WeightLifted float64
	SetType      string
}

type WorkoutJSON struct {
	Lifts map[string][]LiftJSON
}

// asserts both create and get one queries
func TestCreateWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	GenRandWorkout(t, userId.String())
}

func TestListWorkouts(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	// empty
	query, err := testQueries.ListWorkouts(context.Background(), userId.String())
	require.NoError(t, err)
	require.Empty(t, query)

	n := 5
	workouts := make([]Workout, n)
	for i := 0; i < n; i++ {
		workouts[i] = GenRandWorkout(t, userId.String())
	}

	// not empty
	query, err = testQueries.ListWorkouts(context.Background(), userId.String())
	require.NoError(t, err)
	require.NotEmpty(t, query)
	require.Len(t, query, n)

	for i, v := range query {
		userIdFromQuery, err := uuid.FromBytes(v.UserID)
		require.NoError(t, err)
		require.Equal(t, userIdFromQuery, userId)
		require.Equal(t, v.Duration, workouts[i].Duration)
		require.Equal(t, v.Lifts, workouts[i].Lifts)
	}
}

func TestUpdateWorkout(t *testing.T) {
	n := 5
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	workoutJson := &WorkoutJSON{}
	workout := GenRandWorkout(t, userId.String())
	err = json.Unmarshal(workout.Lifts, workoutJson)

	newLifts := &WorkoutJSON{Lifts: make(map[string][]LiftJSON)}
	newDuration := "00:30:00s"
	exercises := make([]Exercise, n)
	for i := 0; i < n; i++ {
		exercises[i] = GenRandExercise(t, userId.String())
		insertIntoWorkoutMap(exercises[i].Name, newLifts)
	}

	newLiftsRaw, err := json.Marshal(*newLifts)
	require.NoError(t, err)

	_, err = testQueries.UpdateWorkout(context.Background(), UpdateWorkoutParams{
		ID:       workout.ID,
		Duration: newDuration,
		UserID:   userId.String(),
		Lifts:    newLiftsRaw,
	})
	require.NoError(t, err)

	query, err := testQueries.GetWorkout(context.Background(), GetWorkoutParams{
		ID:     workout.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)

	// new lifts
	queryJson := &WorkoutJSON{}
	err = json.Unmarshal(query.Lifts, queryJson)
	require.Equal(t, query.ID, workout.ID)
	require.Equal(t, *queryJson, *newLifts)
	require.NotEqual(t, *queryJson, *workoutJson)

	// new Duration
	require.Equal(t, query.Duration, newDuration)
	require.NotEqual(t, query.Duration, workout.Duration)
}

func TestDeleteWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())

	_, err = testQueries.DeleteWorkout(context.Background(), DeleteWorkoutParams{
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
	require.Empty(t, query.UserID)
	require.Empty(t, query.Lifts)
}

func insertIntoWorkoutMap(exerciseName string, mp *WorkoutJSON) *WorkoutJSON {
	liftSetRange := int(util.RandomInt(1, 3))

	lift := &LiftJSON{
		ExerciseName: exerciseName,
		Reps:         int32(util.RandomInt(5, 20)),
		SetType:      "Working Set",
		WeightLifted: float64(util.RandomInt(100, 250)),
	}

	for i := 0; i < liftSetRange; i++ {
		mp.Lifts[exerciseName] = append(mp.Lifts[exerciseName], *lift)
	}

	return mp
}

func GenRandWorkout(t *testing.T, userId string) Workout {
	n := 5
	buildWorkout := &WorkoutJSON{make(map[string][]LiftJSON)}

	workout := &Workout{}

	exercises := make([]Exercise, n)

	for i := 0; i < n; i++ {
		exercises[i] = GenRandExercise(t, userId)
		insertIntoWorkoutMap(exercises[i].Name, buildWorkout)
	}

	rawJson, err := json.Marshal(*buildWorkout)
	require.NoError(t, err)

	record, err := testQueries.CreateWorkout(context.Background(), CreateWorkoutParams{
		Lifts:    rawJson,
		UserID:   userId,
		Duration: "01:05:02s",
	})

	workoutId, err := record.LastInsertId()
	require.NoError(t, err)

	query, err := testQueries.GetWorkout(context.Background(), GetWorkoutParams{
		ID:     int32(workoutId),
		UserID: userId,
	})

	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ID, int32(workoutId))

	queryJsonLifts := &WorkoutJSON{}
	err = json.Unmarshal(query.Lifts, queryJsonLifts)
	require.NoError(t, err)

	require.Equal(t, *queryJsonLifts, *buildWorkout)

	workout = &query
	return *workout
}

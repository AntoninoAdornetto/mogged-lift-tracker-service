package db

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type NewLiftArgs struct {
	UserID    string
	WorkoutID int32
}

// asserts both create and get one queries
func TestCreateLift(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	workout := GenRandWorkout(t, userId.String())
	GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
}

func TestListLiftsFromWorkout(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())

	query, err := testQueries.ListLiftsFromWorkout(context.Background(), ListLiftsFromWorkoutParams{
		UserID:    userId.String(),
		WorkoutID: workout.ID,
	})
	require.NoError(t, err)
	require.Empty(t, query)

	n := 5
	lifts := make([]Lift, n)
	for i := 0; i < n; i++ {
		lifts[i] = GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
	}

	query, err = testQueries.ListLiftsFromWorkout(context.Background(), ListLiftsFromWorkoutParams{
		UserID:    userId.String(),
		WorkoutID: workout.ID,
	})
	require.NoError(t, err)
	require.Len(t, query, n)
}

func TestGetMaxLifts(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	n := 5
	workout := GenRandWorkout(t, userId.String())

	query, err := testQueries.GetMaxLifts(context.Background(), GetMaxLiftsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Empty(t, query)

	for i := 0; i < n; i++ {
		GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
	}

	query, err = testQueries.GetMaxLifts(context.Background(), GetMaxLiftsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Len(t, query, n)

	weight := query[0].WeightLifted
	for i := 1; i < n; i++ {
		require.LessOrEqual(t, query[i].WeightLifted, weight)
		weight = query[i].WeightLifted
	}
}

func TestGetMaxRepLifts(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	n := 5
	workout := GenRandWorkout(t, userId.String())

	query, err := testQueries.GetMaxRepLifts(context.Background(), GetMaxRepLiftsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Empty(t, query)

	for i := 0; i < n; i++ {
		GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})
	}

	query, err = testQueries.GetMaxRepLifts(context.Background(), GetMaxRepLiftsParams{
		UserID: userId.String(),
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Len(t, query, n)

	reps := query[0].Reps
	for i := 1; i < n; i++ {
		require.LessOrEqual(t, query[i].Reps, reps)
		reps = query[i].Reps
	}
}

func TestGetMaxLiftsByExercise(t *testing.T) {
	user := GenRandUser(t) // fresh user no lifts tracked as of yet
	exerciseName := util.RandomStr(10)
	amountOfLifts := 10

	lifts := GenRandLiftsSameExercise(t, createExerciseLifts{
		userID:        user.ID,
		AmountOfLifts: amountOfLifts,
		exerciseName:  exerciseName,
	})

	// sorting manually to ensure that the maxLifts query match the sorted lifts array
	sort.Slice(lifts, func(i, j int) bool {
		return lifts[i].WeightLifted > lifts[j].WeightLifted
	})

	maxLifts, err := testQueries.GetMaxLiftsByExercise(context.Background(), GetMaxLiftsByExerciseParams{
		ExerciseName: exerciseName,
		UserID:       user.ID,
	})
	require.NoError(t, err)
	fmt.Println(maxLifts)
	require.Len(t, maxLifts, amountOfLifts)
	require.Equal(t, len(maxLifts), len(lifts))

	for i, v := range lifts {
		require.Equal(t, maxLifts[i].WeightLifted, v.WeightLifted)
	}
}

func TestGetMaxLiftsByMuscleGroup(t *testing.T) {
	user := GenRandUser(t)
	amountOfLifts := 3
	muscleGroup := "Chest"
	exercises := []string{"Barbell Bench Press", "Incline Barbell Bench Press"}

	for _, v := range exercises {
		for i := 0; i < amountOfLifts; i++ {
			GenRandLiftsByMuscleGroup(t, genRandLiftByMuscleGroup{ExerciseName: v, MuscleGroup: muscleGroup, UserID: user.ID})
		}
	}

	maxLifts, err := testQueries.GetMaxLiftsByMuscleGroup(context.Background(), GetMaxLiftsByMuscleGroupParams{
		MuscleGroup: muscleGroup,
		UserID:      user.ID,
	})
	require.NoError(t, err)

	n := len(maxLifts) - 1
	for n > 0 {
		require.LessOrEqual(t, maxLifts[n].WeightLifted, maxLifts[n-1].WeightLifted)
		n--
	}
	require.Zero(t, n)
}

func TestUpdateLift(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())
	lift := GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})

	newExerciseName := GenRandExercise(t, userId.String()).Name
	err = testQueries.UpdateLift(context.Background(), UpdateLiftParams{
		ExerciseName: sql.NullString{
			String: newExerciseName,
			Valid:  true,
		},
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)
	require.Equal(t, query.ID, lift.ID)
	require.Equal(t, query.ExerciseName, newExerciseName)

	newWeightLifted := float64(util.RandomInt(100, 500))
	err = testQueries.UpdateLift(context.Background(), UpdateLiftParams{
		WeightLifted: sql.NullFloat64{
			Float64: newWeightLifted,
			Valid:   true,
		},
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)

	query, err = testQueries.GetLift(context.Background(), GetLiftParams{
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)
	require.Equal(t, query.ID, lift.ID)
	require.Equal(t, query.WeightLifted, newWeightLifted)

	newReps := int32(util.RandomInt(100, 500))
	err = testQueries.UpdateLift(context.Background(), UpdateLiftParams{
		Reps: sql.NullInt32{
			Int32: newReps,
			Valid: true,
		},
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)

	query, err = testQueries.GetLift(context.Background(), GetLiftParams{
		UserID: userId.String(),
		ID:     lift.ID,
	})
	require.NoError(t, err)
	require.Equal(t, query.ID, lift.ID)
	require.Equal(t, query.Reps, newReps)
}

func TestUpdateSetType(t *testing.T) {
	user := GenRandUser(t)
	workout := GenRandWorkout(t, user.ID)
	lift := GenRandLift(t, NewLiftArgs{UserID: user.ID, WorkoutID: workout.ID})

	newSetType := sql.NullString{String: "warmup set", Valid: true}

	err := testQueries.UpdateLift(context.Background(), UpdateLiftParams{
		ID:      lift.ID,
		SetType: newSetType,
		UserID:  user.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{ID: lift.ID, UserID: user.ID})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.SetType, newSetType.String)
	require.NotEqual(t, query.SetType, lift.SetType)
}

func TestDeleteLift(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	workout := GenRandWorkout(t, userId.String())
	lift := GenRandLift(t, NewLiftArgs{UserID: userId.String(), WorkoutID: workout.ID})

	err = testQueries.DeleteLift(context.Background(), DeleteLiftParams{
		ID:     lift.ID,
		UserID: userId.String(),
	})
	require.NoError(t, err)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{
		ID:     lift.ID,
		UserID: userId.String(),
	})
	require.Error(t, err)
	require.Zero(t, query.ID)
}

func GenRandLift(t *testing.T, args NewLiftArgs) Lift {
	lift := &Lift{}
	exercise := GenRandExercise(t, args.UserID)

	createLiftArgs := CreateLiftParams{
		ExerciseName: exercise.Name,
		Reps:         int32(util.RandomInt(3, 12)),
		WeightLifted: float64(util.RandomInt(100, 300)),
		UserID:       args.UserID,
		WorkoutID:    args.WorkoutID,
		SetType:      "Working",
	}
	record, err := testQueries.CreateLift(context.Background(), createLiftArgs)
	require.NoError(t, err)
	lastId, err := record.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, lastId)

	query, err := testQueries.GetLift(context.Background(), GetLiftParams{UserID: args.UserID, ID: lastId})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.ExerciseName, createLiftArgs.ExerciseName)
	require.Equal(t, query.Reps, createLiftArgs.Reps)
	require.Equal(t, query.WeightLifted, createLiftArgs.WeightLifted)
	require.Equal(t, query.WorkoutID, args.WorkoutID)
	require.Equal(t, query.SetType, createLiftArgs.SetType)

	queryUserId, err := uuid.FromBytes(query.UserID)
	require.NoError(t, err)
	require.Equal(t, queryUserId.String(), args.UserID)

	lift = &query
	return *lift
}

type genRandLiftByMuscleGroup struct {
	ExerciseName string
	MuscleGroup  string
	UserID       string
}

// use real muscle groups & exercises that are part of seed migration
func GenRandLiftsByMuscleGroup(t *testing.T, args genRandLiftByMuscleGroup) {
	muscleGroup, err := testQueries.GetMuscleGroupByName(context.Background(), args.MuscleGroup)
	require.NoError(t, err)
	require.NotZero(t, muscleGroup.ID)

	_, err = testQueries.GetExerciseByName(context.Background(), GetExerciseByNameParams{Name: args.ExerciseName, UserID: args.UserID})
	if err != nil {
		testQueries.CreateExercise(context.Background(), CreateExerciseParams{
			Name:        args.ExerciseName,
			MuscleGroup: muscleGroup.Name,
			Category:    "Barbell",
			UserID:      args.UserID,
		})
	}
	exercise, err := testQueries.GetExerciseByName(context.Background(), GetExerciseByNameParams{Name: args.ExerciseName, UserID: args.UserID})
	require.NotZero(t, exercise.ID)

	workout, err := testQueries.CreateWorkout(context.Background(), CreateWorkoutParams{UserID: args.UserID})
	require.NoError(t, err)
	workoutId, err := workout.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, workoutId)

	_, err = testQueries.CreateLift(context.Background(), CreateLiftParams{
		ExerciseName: exercise.Name,
		UserID:       args.UserID,
		Reps:         int32(util.RandomInt(6, 20)),
		WeightLifted: float64(util.RandomInt(100, 285)),
		SetType:      "working",
		WorkoutID:    int32(workoutId),
	})
	require.NoError(t, err)
}

type createExerciseLifts struct {
	exerciseName  string
	userID        string
	AmountOfLifts int
}

func GenRandLiftsSameExercise(t *testing.T, args createExerciseLifts) []Lift {
	muscleGroup := GenRandMuscleGroup(t)
	category := GenRandCategory(t)

	id, err := testQueries.CreateExercise(context.Background(), CreateExerciseParams{
		Name:        args.exerciseName,
		Category:    category.Name,
		MuscleGroup: muscleGroup.Name,
		UserID:      args.userID,
	})
	require.NoError(t, err)

	res, err := testQueries.CreateWorkout(context.Background(), CreateWorkoutParams{UserID: args.userID})
	require.NoError(t, err)
	workoutId, err := res.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, workoutId)

	exercise, err := testQueries.GetExercise(context.Background(), GetExerciseParams{ID: int32(id), UserID: args.userID})
	require.NoError(t, err)
	require.NotZero(t, exercise.ID)

	lifts := make([]Lift, args.AmountOfLifts)
	for i := range lifts {
		res, err := testQueries.CreateLift(context.Background(), CreateLiftParams{
			WeightLifted: float64(util.RandomInt(100, 285)),
			UserID:       args.userID,
			ExerciseName: exercise.Name,
			Reps:         int32(util.RandomInt(6, 30)),
			SetType:      "working",
			WorkoutID:    int32(workoutId),
		})
		require.NoError(t, err)
		liftId, err := res.LastInsertId()
		require.NoError(t, err)
		require.NotZero(t, liftId)

		lift, err := testQueries.GetLift(context.Background(), GetLiftParams{ID: liftId, UserID: args.userID})
		lifts[i] = lift
	}

	return lifts
}

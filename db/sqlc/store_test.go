package db

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestWorkoutTx(t *testing.T) {
	store := NewStore(testDB)
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	lifts := buildLiftsMap(t, userId.String())
	liftsMap := make(map[string][]Lift)
	err = json.Unmarshal(lifts, &liftsMap)
	require.NoError(t, err)

	args := WorkoutTxParams{
		UserID:   userId.String(),
		Duration: "01:25:00s",
		LiftsMap: lifts,
	}

	workout, err := store.WorkoutTx(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, workout)
	require.NotZero(t, workout.ID)
	require.Equal(t, workout.Duration, args.Duration)

	workoutUserId, err := uuid.FromBytes(workout.UserID)
	require.NoError(t, err)
	require.Equal(t, workoutUserId, userId)

	workoutLiftsMap := make(map[string][]Lift)
	err = json.Unmarshal(workout.Lifts, &workoutLiftsMap)

	for key, lifts := range workoutLiftsMap {
		require.NotEmpty(t, liftsMap[key])

		for i, lift := range lifts {
			require.Equal(t, lift.ID, liftsMap[key][i].ID)
			require.Equal(t, lift.WorkoutID, liftsMap[key][i].WorkoutID)
			require.Equal(t, lift.WeightLifted, liftsMap[key][i].WeightLifted)
			require.Equal(t, lift.Reps, liftsMap[key][i].Reps)
			require.Equal(t, lift.ExerciseName, liftsMap[key][i].ExerciseName)
			require.Equal(t, lift.SetType, liftsMap[key][i].SetType)
		}
	}
}

func buildLiftsMap(t *testing.T, userId string) []byte {
	n := 5
	sets := 3

	exercises := make([]Exercise, n)
	liftsMap := make(map[string][]CreateLiftParams)

	for i := 0; i < n; i++ {
		exercises[i] = GenRandExercise(t, userId)
		for j := 0; j < sets; j++ {
			lift := CreateLiftParams{
				ExerciseName: exercises[i].Name,
				Reps:         int32(util.RandomInt(6, 12)),
				SetType:      "Working",
				WeightLifted: float64(util.RandomInt(100, 250)),
			}
			liftsMap[exercises[i].Name] = append(liftsMap[exercises[i].Name], lift)
		}
	}

	rawJson, err := json.Marshal(liftsMap)
	require.NoError(t, err)
	return rawJson
}

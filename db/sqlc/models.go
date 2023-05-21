// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"encoding/json"
	"time"
)

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Exercise struct {
	ID               int32   `json:"id"`
	Name             string  `json:"name"`
	MuscleGroup      string  `json:"muscle_group"`
	Category         string  `json:"category"`
	Isstock          bool    `json:"isstock"`
	MostWeightLifted float64 `json:"most_weight_lifted"`
	MostRepsLifted   int32   `json:"most_reps_lifted"`
	RestTimer        string  `json:"rest_timer"`
	UserID           []byte  `json:"user_id"`
}

type Lift struct {
	ID           int64   `json:"id"`
	ExerciseName string  `json:"exercise_name"`
	WeightLifted float64 `json:"weight_lifted"`
	Reps         int32   `json:"reps"`
	SetType      string  `json:"set_type"`
	UserID       []byte  `json:"user_id"`
	WorkoutID    int32   `json:"workout_id"`
}

type MuscleGroup struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Profile struct {
	ID                int32   `json:"id"`
	Country           string  `json:"country"`
	MeasurementSystem string  `json:"measurement_system"`
	BodyWeight        float64 `json:"body_weight"`
	BodyFat           float64 `json:"body_fat"`
	TimezoneOffset    int32   `json:"timezone_offset"`
	UserID            []byte  `json:"user_id"`
}

type StockExercise struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
	Category    string `json:"category"`
}

type Template struct {
	ID           int32           `json:"id"`
	Name         string          `json:"name"`
	Exercises    json.RawMessage `json:"exercises"`
	DateLastUsed time.Time       `json:"date_last_used"`
	CreatedBy    []byte          `json:"created_by"`
}

type User struct {
	ID                []byte    `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	EmailAddress      string    `json:"email_address"`
	Password          string    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

type Workout struct {
	ID       int32           `json:"id"`
	Duration string          `json:"duration"`
	Lifts    json.RawMessage `json:"lifts"`
	UserID   []byte          `json:"user_id"`
}

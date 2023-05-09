// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	ChangePassword(ctx context.Context, arg ChangePasswordParams) error
	CreateCategory(ctx context.Context, name string) (sql.Result, error)
	CreateExercise(ctx context.Context, arg CreateExerciseParams) (sql.Result, error)
	CreateLift(ctx context.Context, arg CreateLiftParams) (sql.Result, error)
	CreateMuscleGroup(ctx context.Context, name string) (sql.Result, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (int64, error)
	CreateStockExercise(ctx context.Context, arg CreateStockExerciseParams) (sql.Result, error)
	CreateTemplate(ctx context.Context, arg CreateTemplateParams) (sql.Result, error)
	CreateUser(ctx context.Context, arg CreateUserParams) error
	CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (sql.Result, error)
	DeleteCategory(ctx context.Context, id int32) (sql.Result, error)
	DeleteExercise(ctx context.Context, arg DeleteExerciseParams) error
	DeleteLift(ctx context.Context, arg DeleteLiftParams) (sql.Result, error)
	DeleteMuscleGroup(ctx context.Context, id int32) (sql.Result, error)
	DeleteProfile(ctx context.Context, userID string) error
	DeleteStockExercise(ctx context.Context, id int32) (sql.Result, error)
	DeleteTemplate(ctx context.Context, arg DeleteTemplateParams) (sql.Result, error)
	DeleteUser(ctx context.Context, userID string) error
	DeleteWorkout(ctx context.Context, arg DeleteWorkoutParams) (sql.Result, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetExercise(ctx context.Context, arg GetExerciseParams) (Exercise, error)
	GetExerciseByName(ctx context.Context, arg GetExerciseByNameParams) (Exercise, error)
	GetLift(ctx context.Context, arg GetLiftParams) (Lift, error)
	GetMuscleGroup(ctx context.Context, id int32) (MuscleGroup, error)
	GetProfile(ctx context.Context, userID string) (Profile, error)
	GetStockExercise(ctx context.Context, id int32) (StockExercise, error)
	GetTemplate(ctx context.Context, arg GetTemplateParams) (Template, error)
	GetUserByEmail(ctx context.Context, emailAddress string) (GetUserByEmailRow, error)
	GetUserById(ctx context.Context, userID string) (GetUserByIdRow, error)
	GetWorkout(ctx context.Context, arg GetWorkoutParams) (Workout, error)
	ListCategories(ctx context.Context) ([]Category, error)
	ListExercises(ctx context.Context, userID string) ([]Exercise, error)
	ListLiftsFromWorkout(ctx context.Context, arg ListLiftsFromWorkoutParams) ([]Lift, error)
	ListMaxRepPrs(ctx context.Context, arg ListMaxRepPrsParams) ([]Lift, error)
	ListMaxWeightPrs(ctx context.Context, arg ListMaxWeightPrsParams) ([]Lift, error)
	ListMuscleGroups(ctx context.Context) ([]MuscleGroup, error)
	ListStockExercies(ctx context.Context) ([]StockExercise, error)
	ListTemplates(ctx context.Context, createdBy string) ([]Template, error)
	ListWorkouts(ctx context.Context, userID string) ([]Workout, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (sql.Result, error)
	UpdateExercise(ctx context.Context, arg UpdateExerciseParams) error
	UpdateLift(ctx context.Context, arg UpdateLiftParams) (sql.Result, error)
	UpdateMuscleGroup(ctx context.Context, arg UpdateMuscleGroupParams) (sql.Result, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (sql.Result, error)
	UpdateStockExercise(ctx context.Context, arg UpdateStockExerciseParams) (sql.Result, error)
	UpdateTemplate(ctx context.Context, arg UpdateTemplateParams) (int64, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)

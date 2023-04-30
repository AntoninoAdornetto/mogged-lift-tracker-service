// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.changePasswordStmt, err = db.PrepareContext(ctx, changePassword); err != nil {
		return nil, fmt.Errorf("error preparing query ChangePassword: %w", err)
	}
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createExerciseStmt, err = db.PrepareContext(ctx, createExercise); err != nil {
		return nil, fmt.Errorf("error preparing query CreateExercise: %w", err)
	}
	if q.createLiftStmt, err = db.PrepareContext(ctx, createLift); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLift: %w", err)
	}
	if q.createMuscleGroupStmt, err = db.PrepareContext(ctx, createMuscleGroup); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMuscleGroup: %w", err)
	}
	if q.createProfileStmt, err = db.PrepareContext(ctx, createProfile); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProfile: %w", err)
	}
	if q.createStockExerciseStmt, err = db.PrepareContext(ctx, createStockExercise); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStockExercise: %w", err)
	}
	if q.createTemplateStmt, err = db.PrepareContext(ctx, createTemplate); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTemplate: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.createWorkoutStmt, err = db.PrepareContext(ctx, createWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWorkout: %w", err)
	}
	if q.deleteCategoryStmt, err = db.PrepareContext(ctx, deleteCategory); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCategory: %w", err)
	}
	if q.deleteExerciseStmt, err = db.PrepareContext(ctx, deleteExercise); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExercise: %w", err)
	}
	if q.deleteLiftStmt, err = db.PrepareContext(ctx, deleteLift); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteLift: %w", err)
	}
	if q.deleteMuscleGroupStmt, err = db.PrepareContext(ctx, deleteMuscleGroup); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMuscleGroup: %w", err)
	}
	if q.deleteProfileStmt, err = db.PrepareContext(ctx, deleteProfile); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProfile: %w", err)
	}
	if q.deleteStockExerciseStmt, err = db.PrepareContext(ctx, deleteStockExercise); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteStockExercise: %w", err)
	}
	if q.deleteTemplateStmt, err = db.PrepareContext(ctx, deleteTemplate); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTemplate: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.deleteWorkoutStmt, err = db.PrepareContext(ctx, deleteWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteWorkout: %w", err)
	}
	if q.getCategoryStmt, err = db.PrepareContext(ctx, getCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategory: %w", err)
	}
	if q.getExerciseStmt, err = db.PrepareContext(ctx, getExercise); err != nil {
		return nil, fmt.Errorf("error preparing query GetExercise: %w", err)
	}
	if q.getExerciseByNameStmt, err = db.PrepareContext(ctx, getExerciseByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetExerciseByName: %w", err)
	}
	if q.getLiftStmt, err = db.PrepareContext(ctx, getLift); err != nil {
		return nil, fmt.Errorf("error preparing query GetLift: %w", err)
	}
	if q.getMuscleGroupStmt, err = db.PrepareContext(ctx, getMuscleGroup); err != nil {
		return nil, fmt.Errorf("error preparing query GetMuscleGroup: %w", err)
	}
	if q.getProfileStmt, err = db.PrepareContext(ctx, getProfile); err != nil {
		return nil, fmt.Errorf("error preparing query GetProfile: %w", err)
	}
	if q.getStockExerciseStmt, err = db.PrepareContext(ctx, getStockExercise); err != nil {
		return nil, fmt.Errorf("error preparing query GetStockExercise: %w", err)
	}
	if q.getTemplateStmt, err = db.PrepareContext(ctx, getTemplate); err != nil {
		return nil, fmt.Errorf("error preparing query GetTemplate: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getWorkoutStmt, err = db.PrepareContext(ctx, getWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query GetWorkout: %w", err)
	}
	if q.listCategoriesStmt, err = db.PrepareContext(ctx, listCategories); err != nil {
		return nil, fmt.Errorf("error preparing query ListCategories: %w", err)
	}
	if q.listExercisesStmt, err = db.PrepareContext(ctx, listExercises); err != nil {
		return nil, fmt.Errorf("error preparing query ListExercises: %w", err)
	}
	if q.listLiftsFromWorkoutStmt, err = db.PrepareContext(ctx, listLiftsFromWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query ListLiftsFromWorkout: %w", err)
	}
	if q.listMaxRepPrsStmt, err = db.PrepareContext(ctx, listMaxRepPrs); err != nil {
		return nil, fmt.Errorf("error preparing query ListMaxRepPrs: %w", err)
	}
	if q.listMaxWeightPrsStmt, err = db.PrepareContext(ctx, listMaxWeightPrs); err != nil {
		return nil, fmt.Errorf("error preparing query ListMaxWeightPrs: %w", err)
	}
	if q.listMuscleGroupsStmt, err = db.PrepareContext(ctx, listMuscleGroups); err != nil {
		return nil, fmt.Errorf("error preparing query ListMuscleGroups: %w", err)
	}
	if q.listStockExerciesStmt, err = db.PrepareContext(ctx, listStockExercies); err != nil {
		return nil, fmt.Errorf("error preparing query ListStockExercies: %w", err)
	}
	if q.listTemplatesStmt, err = db.PrepareContext(ctx, listTemplates); err != nil {
		return nil, fmt.Errorf("error preparing query ListTemplates: %w", err)
	}
	if q.listWorkoutsStmt, err = db.PrepareContext(ctx, listWorkouts); err != nil {
		return nil, fmt.Errorf("error preparing query ListWorkouts: %w", err)
	}
	if q.updateCategoryStmt, err = db.PrepareContext(ctx, updateCategory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCategory: %w", err)
	}
	if q.updateExerciseStmt, err = db.PrepareContext(ctx, updateExercise); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateExercise: %w", err)
	}
	if q.updateLiftStmt, err = db.PrepareContext(ctx, updateLift); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateLift: %w", err)
	}
	if q.updateMuscleGroupStmt, err = db.PrepareContext(ctx, updateMuscleGroup); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMuscleGroup: %w", err)
	}
	if q.updateProfileStmt, err = db.PrepareContext(ctx, updateProfile); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProfile: %w", err)
	}
	if q.updateStockExerciseStmt, err = db.PrepareContext(ctx, updateStockExercise); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateStockExercise: %w", err)
	}
	if q.updateTemplateStmt, err = db.PrepareContext(ctx, updateTemplate); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTemplate: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	if q.updateWorkoutStmt, err = db.PrepareContext(ctx, updateWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateWorkout: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.changePasswordStmt != nil {
		if cerr := q.changePasswordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing changePasswordStmt: %w", cerr)
		}
	}
	if q.createCategoryStmt != nil {
		if cerr := q.createCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCategoryStmt: %w", cerr)
		}
	}
	if q.createExerciseStmt != nil {
		if cerr := q.createExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createExerciseStmt: %w", cerr)
		}
	}
	if q.createLiftStmt != nil {
		if cerr := q.createLiftStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLiftStmt: %w", cerr)
		}
	}
	if q.createMuscleGroupStmt != nil {
		if cerr := q.createMuscleGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMuscleGroupStmt: %w", cerr)
		}
	}
	if q.createProfileStmt != nil {
		if cerr := q.createProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProfileStmt: %w", cerr)
		}
	}
	if q.createStockExerciseStmt != nil {
		if cerr := q.createStockExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStockExerciseStmt: %w", cerr)
		}
	}
	if q.createTemplateStmt != nil {
		if cerr := q.createTemplateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTemplateStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.createWorkoutStmt != nil {
		if cerr := q.createWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWorkoutStmt: %w", cerr)
		}
	}
	if q.deleteCategoryStmt != nil {
		if cerr := q.deleteCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCategoryStmt: %w", cerr)
		}
	}
	if q.deleteExerciseStmt != nil {
		if cerr := q.deleteExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExerciseStmt: %w", cerr)
		}
	}
	if q.deleteLiftStmt != nil {
		if cerr := q.deleteLiftStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteLiftStmt: %w", cerr)
		}
	}
	if q.deleteMuscleGroupStmt != nil {
		if cerr := q.deleteMuscleGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMuscleGroupStmt: %w", cerr)
		}
	}
	if q.deleteProfileStmt != nil {
		if cerr := q.deleteProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProfileStmt: %w", cerr)
		}
	}
	if q.deleteStockExerciseStmt != nil {
		if cerr := q.deleteStockExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteStockExerciseStmt: %w", cerr)
		}
	}
	if q.deleteTemplateStmt != nil {
		if cerr := q.deleteTemplateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTemplateStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.deleteWorkoutStmt != nil {
		if cerr := q.deleteWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteWorkoutStmt: %w", cerr)
		}
	}
	if q.getCategoryStmt != nil {
		if cerr := q.getCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryStmt: %w", cerr)
		}
	}
	if q.getExerciseStmt != nil {
		if cerr := q.getExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getExerciseStmt: %w", cerr)
		}
	}
	if q.getExerciseByNameStmt != nil {
		if cerr := q.getExerciseByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getExerciseByNameStmt: %w", cerr)
		}
	}
	if q.getLiftStmt != nil {
		if cerr := q.getLiftStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLiftStmt: %w", cerr)
		}
	}
	if q.getMuscleGroupStmt != nil {
		if cerr := q.getMuscleGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMuscleGroupStmt: %w", cerr)
		}
	}
	if q.getProfileStmt != nil {
		if cerr := q.getProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProfileStmt: %w", cerr)
		}
	}
	if q.getStockExerciseStmt != nil {
		if cerr := q.getStockExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStockExerciseStmt: %w", cerr)
		}
	}
	if q.getTemplateStmt != nil {
		if cerr := q.getTemplateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTemplateStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.getWorkoutStmt != nil {
		if cerr := q.getWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getWorkoutStmt: %w", cerr)
		}
	}
	if q.listCategoriesStmt != nil {
		if cerr := q.listCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCategoriesStmt: %w", cerr)
		}
	}
	if q.listExercisesStmt != nil {
		if cerr := q.listExercisesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listExercisesStmt: %w", cerr)
		}
	}
	if q.listLiftsFromWorkoutStmt != nil {
		if cerr := q.listLiftsFromWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listLiftsFromWorkoutStmt: %w", cerr)
		}
	}
	if q.listMaxRepPrsStmt != nil {
		if cerr := q.listMaxRepPrsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMaxRepPrsStmt: %w", cerr)
		}
	}
	if q.listMaxWeightPrsStmt != nil {
		if cerr := q.listMaxWeightPrsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMaxWeightPrsStmt: %w", cerr)
		}
	}
	if q.listMuscleGroupsStmt != nil {
		if cerr := q.listMuscleGroupsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMuscleGroupsStmt: %w", cerr)
		}
	}
	if q.listStockExerciesStmt != nil {
		if cerr := q.listStockExerciesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listStockExerciesStmt: %w", cerr)
		}
	}
	if q.listTemplatesStmt != nil {
		if cerr := q.listTemplatesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTemplatesStmt: %w", cerr)
		}
	}
	if q.listWorkoutsStmt != nil {
		if cerr := q.listWorkoutsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listWorkoutsStmt: %w", cerr)
		}
	}
	if q.updateCategoryStmt != nil {
		if cerr := q.updateCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCategoryStmt: %w", cerr)
		}
	}
	if q.updateExerciseStmt != nil {
		if cerr := q.updateExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateExerciseStmt: %w", cerr)
		}
	}
	if q.updateLiftStmt != nil {
		if cerr := q.updateLiftStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateLiftStmt: %w", cerr)
		}
	}
	if q.updateMuscleGroupStmt != nil {
		if cerr := q.updateMuscleGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMuscleGroupStmt: %w", cerr)
		}
	}
	if q.updateProfileStmt != nil {
		if cerr := q.updateProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProfileStmt: %w", cerr)
		}
	}
	if q.updateStockExerciseStmt != nil {
		if cerr := q.updateStockExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateStockExerciseStmt: %w", cerr)
		}
	}
	if q.updateTemplateStmt != nil {
		if cerr := q.updateTemplateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTemplateStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	if q.updateWorkoutStmt != nil {
		if cerr := q.updateWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateWorkoutStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                       DBTX
	tx                       *sql.Tx
	changePasswordStmt       *sql.Stmt
	createCategoryStmt       *sql.Stmt
	createExerciseStmt       *sql.Stmt
	createLiftStmt           *sql.Stmt
	createMuscleGroupStmt    *sql.Stmt
	createProfileStmt        *sql.Stmt
	createStockExerciseStmt  *sql.Stmt
	createTemplateStmt       *sql.Stmt
	createUserStmt           *sql.Stmt
	createWorkoutStmt        *sql.Stmt
	deleteCategoryStmt       *sql.Stmt
	deleteExerciseStmt       *sql.Stmt
	deleteLiftStmt           *sql.Stmt
	deleteMuscleGroupStmt    *sql.Stmt
	deleteProfileStmt        *sql.Stmt
	deleteStockExerciseStmt  *sql.Stmt
	deleteTemplateStmt       *sql.Stmt
	deleteUserStmt           *sql.Stmt
	deleteWorkoutStmt        *sql.Stmt
	getCategoryStmt          *sql.Stmt
	getExerciseStmt          *sql.Stmt
	getExerciseByNameStmt    *sql.Stmt
	getLiftStmt              *sql.Stmt
	getMuscleGroupStmt       *sql.Stmt
	getProfileStmt           *sql.Stmt
	getStockExerciseStmt     *sql.Stmt
	getTemplateStmt          *sql.Stmt
	getUserStmt              *sql.Stmt
	getUserByIdStmt          *sql.Stmt
	getWorkoutStmt           *sql.Stmt
	listCategoriesStmt       *sql.Stmt
	listExercisesStmt        *sql.Stmt
	listLiftsFromWorkoutStmt *sql.Stmt
	listMaxRepPrsStmt        *sql.Stmt
	listMaxWeightPrsStmt     *sql.Stmt
	listMuscleGroupsStmt     *sql.Stmt
	listStockExerciesStmt    *sql.Stmt
	listTemplatesStmt        *sql.Stmt
	listWorkoutsStmt         *sql.Stmt
	updateCategoryStmt       *sql.Stmt
	updateExerciseStmt       *sql.Stmt
	updateLiftStmt           *sql.Stmt
	updateMuscleGroupStmt    *sql.Stmt
	updateProfileStmt        *sql.Stmt
	updateStockExerciseStmt  *sql.Stmt
	updateTemplateStmt       *sql.Stmt
	updateUserStmt           *sql.Stmt
	updateWorkoutStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		changePasswordStmt:       q.changePasswordStmt,
		createCategoryStmt:       q.createCategoryStmt,
		createExerciseStmt:       q.createExerciseStmt,
		createLiftStmt:           q.createLiftStmt,
		createMuscleGroupStmt:    q.createMuscleGroupStmt,
		createProfileStmt:        q.createProfileStmt,
		createStockExerciseStmt:  q.createStockExerciseStmt,
		createTemplateStmt:       q.createTemplateStmt,
		createUserStmt:           q.createUserStmt,
		createWorkoutStmt:        q.createWorkoutStmt,
		deleteCategoryStmt:       q.deleteCategoryStmt,
		deleteExerciseStmt:       q.deleteExerciseStmt,
		deleteLiftStmt:           q.deleteLiftStmt,
		deleteMuscleGroupStmt:    q.deleteMuscleGroupStmt,
		deleteProfileStmt:        q.deleteProfileStmt,
		deleteStockExerciseStmt:  q.deleteStockExerciseStmt,
		deleteTemplateStmt:       q.deleteTemplateStmt,
		deleteUserStmt:           q.deleteUserStmt,
		deleteWorkoutStmt:        q.deleteWorkoutStmt,
		getCategoryStmt:          q.getCategoryStmt,
		getExerciseStmt:          q.getExerciseStmt,
		getExerciseByNameStmt:    q.getExerciseByNameStmt,
		getLiftStmt:              q.getLiftStmt,
		getMuscleGroupStmt:       q.getMuscleGroupStmt,
		getProfileStmt:           q.getProfileStmt,
		getStockExerciseStmt:     q.getStockExerciseStmt,
		getTemplateStmt:          q.getTemplateStmt,
		getUserStmt:              q.getUserStmt,
		getUserByIdStmt:          q.getUserByIdStmt,
		getWorkoutStmt:           q.getWorkoutStmt,
		listCategoriesStmt:       q.listCategoriesStmt,
		listExercisesStmt:        q.listExercisesStmt,
		listLiftsFromWorkoutStmt: q.listLiftsFromWorkoutStmt,
		listMaxRepPrsStmt:        q.listMaxRepPrsStmt,
		listMaxWeightPrsStmt:     q.listMaxWeightPrsStmt,
		listMuscleGroupsStmt:     q.listMuscleGroupsStmt,
		listStockExerciesStmt:    q.listStockExerciesStmt,
		listTemplatesStmt:        q.listTemplatesStmt,
		listWorkoutsStmt:         q.listWorkoutsStmt,
		updateCategoryStmt:       q.updateCategoryStmt,
		updateExerciseStmt:       q.updateExerciseStmt,
		updateLiftStmt:           q.updateLiftStmt,
		updateMuscleGroupStmt:    q.updateMuscleGroupStmt,
		updateProfileStmt:        q.updateProfileStmt,
		updateStockExerciseStmt:  q.updateStockExerciseStmt,
		updateTemplateStmt:       q.updateTemplateStmt,
		updateUserStmt:           q.updateUserStmt,
		updateWorkoutStmt:        q.updateWorkoutStmt,
	}
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

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
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createExerciseStmt, err = db.PrepareContext(ctx, createExercise); err != nil {
		return nil, fmt.Errorf("error preparing query CreateExercise: %w", err)
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
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteCategoryStmt, err = db.PrepareContext(ctx, deleteCategory); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCategory: %w", err)
	}
	if q.deleteMuscleGroupStmt, err = db.PrepareContext(ctx, deleteMuscleGroup); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMuscleGroup: %w", err)
	}
	if q.deleteProfileStmt, err = db.PrepareContext(ctx, deleteProfile); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProfile: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getCategoryStmt, err = db.PrepareContext(ctx, getCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategory: %w", err)
	}
	if q.getExerciseStmt, err = db.PrepareContext(ctx, getExercise); err != nil {
		return nil, fmt.Errorf("error preparing query GetExercise: %w", err)
	}
	if q.getMuscleGroupStmt, err = db.PrepareContext(ctx, getMuscleGroup); err != nil {
		return nil, fmt.Errorf("error preparing query GetMuscleGroup: %w", err)
	}
	if q.getProfileStmt, err = db.PrepareContext(ctx, getProfile); err != nil {
		return nil, fmt.Errorf("error preparing query GetProfile: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.listCategoriesStmt, err = db.PrepareContext(ctx, listCategories); err != nil {
		return nil, fmt.Errorf("error preparing query ListCategories: %w", err)
	}
	if q.listExercisesStmt, err = db.PrepareContext(ctx, listExercises); err != nil {
		return nil, fmt.Errorf("error preparing query ListExercises: %w", err)
	}
	if q.listMuscleGroupsStmt, err = db.PrepareContext(ctx, listMuscleGroups); err != nil {
		return nil, fmt.Errorf("error preparing query ListMuscleGroups: %w", err)
	}
	if q.updateCategoryStmt, err = db.PrepareContext(ctx, updateCategory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCategory: %w", err)
	}
	if q.updateMuscleGroupStmt, err = db.PrepareContext(ctx, updateMuscleGroup); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMuscleGroup: %w", err)
	}
	if q.updateProfileStmt, err = db.PrepareContext(ctx, updateProfile); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProfile: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
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
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteCategoryStmt != nil {
		if cerr := q.deleteCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCategoryStmt: %w", cerr)
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
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
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
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
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
	if q.listMuscleGroupsStmt != nil {
		if cerr := q.listMuscleGroupsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMuscleGroupsStmt: %w", cerr)
		}
	}
	if q.updateCategoryStmt != nil {
		if cerr := q.updateCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCategoryStmt: %w", cerr)
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
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
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
	db                      DBTX
	tx                      *sql.Tx
	createCategoryStmt      *sql.Stmt
	createExerciseStmt      *sql.Stmt
	createMuscleGroupStmt   *sql.Stmt
	createProfileStmt       *sql.Stmt
	createStockExerciseStmt *sql.Stmt
	createUserStmt          *sql.Stmt
	deleteCategoryStmt      *sql.Stmt
	deleteMuscleGroupStmt   *sql.Stmt
	deleteProfileStmt       *sql.Stmt
	deleteUserStmt          *sql.Stmt
	getCategoryStmt         *sql.Stmt
	getExerciseStmt         *sql.Stmt
	getMuscleGroupStmt      *sql.Stmt
	getProfileStmt          *sql.Stmt
	getUserStmt             *sql.Stmt
	listCategoriesStmt      *sql.Stmt
	listExercisesStmt       *sql.Stmt
	listMuscleGroupsStmt    *sql.Stmt
	updateCategoryStmt      *sql.Stmt
	updateMuscleGroupStmt   *sql.Stmt
	updateProfileStmt       *sql.Stmt
	updateUserStmt          *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                      tx,
		tx:                      tx,
		createCategoryStmt:      q.createCategoryStmt,
		createExerciseStmt:      q.createExerciseStmt,
		createMuscleGroupStmt:   q.createMuscleGroupStmt,
		createProfileStmt:       q.createProfileStmt,
		createStockExerciseStmt: q.createStockExerciseStmt,
		createUserStmt:          q.createUserStmt,
		deleteCategoryStmt:      q.deleteCategoryStmt,
		deleteMuscleGroupStmt:   q.deleteMuscleGroupStmt,
		deleteProfileStmt:       q.deleteProfileStmt,
		deleteUserStmt:          q.deleteUserStmt,
		getCategoryStmt:         q.getCategoryStmt,
		getExerciseStmt:         q.getExerciseStmt,
		getMuscleGroupStmt:      q.getMuscleGroupStmt,
		getProfileStmt:          q.getProfileStmt,
		getUserStmt:             q.getUserStmt,
		listCategoriesStmt:      q.listCategoriesStmt,
		listExercisesStmt:       q.listExercisesStmt,
		listMuscleGroupsStmt:    q.listMuscleGroupsStmt,
		updateCategoryStmt:      q.updateCategoryStmt,
		updateMuscleGroupStmt:   q.updateMuscleGroupStmt,
		updateProfileStmt:       q.updateProfileStmt,
		updateUserStmt:          q.updateUserStmt,
	}
}

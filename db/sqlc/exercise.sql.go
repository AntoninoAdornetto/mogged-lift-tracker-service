// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: exercise.sql

package db

import (
	"context"
	"database/sql"
)

const createExercise = `-- name: CreateExercise :execresult
INSERT INTO exercise (
	name,
	muscle_group,
	category,
	user_id
) VALUES (
	?,
	?,
	?,
	UUID_TO_BIN(?)
)
`

type CreateExerciseParams struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
	Category    string `json:"category"`
	UserID      string `json:"user_id"`
}

func (q *Queries) CreateExercise(ctx context.Context, arg CreateExerciseParams) (sql.Result, error) {
	return q.exec(ctx, q.createExerciseStmt, createExercise,
		arg.Name,
		arg.MuscleGroup,
		arg.Category,
		arg.UserID,
	)
}

const getExercise = `-- name: GetExercise :one
SELECT id, name, muscle_group, category, isstock, most_weight_lifted, most_reps_lifted, rest_timer, user_id FROM exercise
WHERE id = ? LIMIT 1
`

func (q *Queries) GetExercise(ctx context.Context, id int32) (Exercise, error) {
	row := q.queryRow(ctx, q.getExerciseStmt, getExercise, id)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.MuscleGroup,
		&i.Category,
		&i.Isstock,
		&i.MostWeightLifted,
		&i.MostRepsLifted,
		&i.RestTimer,
		&i.UserID,
	)
	return i, err
}

const listExercises = `-- name: ListExercises :many
SELECT id, name, muscle_group, category, isstock, most_weight_lifted, most_reps_lifted, rest_timer, user_id FROM exercise
WHERE user_id = UUID_TO_BIN(?)
`

func (q *Queries) ListExercises(ctx context.Context, userID string) ([]Exercise, error) {
	rows, err := q.query(ctx, q.listExercisesStmt, listExercises, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exercise
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.MuscleGroup,
			&i.Category,
			&i.Isstock,
			&i.MostWeightLifted,
			&i.MostRepsLifted,
			&i.RestTimer,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

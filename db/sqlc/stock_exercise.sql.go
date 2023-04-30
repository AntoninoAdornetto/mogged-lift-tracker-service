// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stock_exercise.sql

package db

import (
	"context"
	"database/sql"
)

const createStockExercise = `-- name: CreateStockExercise :execresult
INSERT INTO stock_exercise (
	name,
	muscle_group,
	category
) VALUES (
	?,
	?,
	?
)
`

type CreateStockExerciseParams struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
	Category    string `json:"category"`
}

func (q *Queries) CreateStockExercise(ctx context.Context, arg CreateStockExerciseParams) (sql.Result, error) {
	return q.exec(ctx, q.createStockExerciseStmt, createStockExercise, arg.Name, arg.MuscleGroup, arg.Category)
}

const deleteStockExercise = `-- name: DeleteStockExercise :execresult
DELETE FROM stock_exercise WHERE
id = ?
`

func (q *Queries) DeleteStockExercise(ctx context.Context, id int32) (sql.Result, error) {
	return q.exec(ctx, q.deleteStockExerciseStmt, deleteStockExercise, id)
}

const getStockExercise = `-- name: GetStockExercise :one
SELECT id, name, muscle_group, category FROM stock_exercise
WHERE id = ?
`

func (q *Queries) GetStockExercise(ctx context.Context, id int32) (StockExercise, error) {
	row := q.queryRow(ctx, q.getStockExerciseStmt, getStockExercise, id)
	var i StockExercise
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.MuscleGroup,
		&i.Category,
	)
	return i, err
}

const listStockExercies = `-- name: ListStockExercies :many
SELECT id, name, muscle_group, category FROM stock_exercise
`

func (q *Queries) ListStockExercies(ctx context.Context) ([]StockExercise, error) {
	rows, err := q.query(ctx, q.listStockExerciesStmt, listStockExercies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StockExercise
	for rows.Next() {
		var i StockExercise
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.MuscleGroup,
			&i.Category,
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

const updateStockExercise = `-- name: UpdateStockExercise :execresult
UPDATE stock_exercise SET 
name = IFNULL(?, name),
muscle_group = IFNULL(?, muscle_group),
category = IFNULL(?, category)
WHERE id = ?
`

type UpdateStockExerciseParams struct {
	Name        interface{} `json:"name"`
	MuscleGroup interface{} `json:"muscle_group"`
	Category    interface{} `json:"category"`
	ID          int32       `json:"id"`
}

func (q *Queries) UpdateStockExercise(ctx context.Context, arg UpdateStockExerciseParams) (sql.Result, error) {
	return q.exec(ctx, q.updateStockExerciseStmt, updateStockExercise,
		arg.Name,
		arg.MuscleGroup,
		arg.Category,
		arg.ID,
	)
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: workout.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createWorkout = `-- name: CreateWorkout :execresult
INSERT INTO workout (
	duration,
	lifts,
	user_id
) VALUES (
	?,
	?,
	UUID_TO_BIN(?)
)
`

type CreateWorkoutParams struct {
	Duration string          `json:"duration"`
	Lifts    json.RawMessage `json:"lifts"`
	UserID   string          `json:"user_id"`
}

func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (sql.Result, error) {
	return q.exec(ctx, q.createWorkoutStmt, createWorkout, arg.Duration, arg.Lifts, arg.UserID)
}

const deleteWorkout = `-- name: DeleteWorkout :execresult
DELETE FROM workout 
WHERE id = ? AND user_id = UUID_TO_BIN(?)
`

type DeleteWorkoutParams struct {
	ID     int32  `json:"id"`
	UserID string `json:"user_id"`
}

func (q *Queries) DeleteWorkout(ctx context.Context, arg DeleteWorkoutParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteWorkoutStmt, deleteWorkout, arg.ID, arg.UserID)
}

const getWorkout = `-- name: GetWorkout :one
SELECT id, duration, lifts, user_id FROM workout
WHERE id = ? AND user_id = UUID_TO_BIN(?)
`

type GetWorkoutParams struct {
	ID     int32  `json:"id"`
	UserID string `json:"user_id"`
}

func (q *Queries) GetWorkout(ctx context.Context, arg GetWorkoutParams) (Workout, error) {
	row := q.queryRow(ctx, q.getWorkoutStmt, getWorkout, arg.ID, arg.UserID)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.Duration,
		&i.Lifts,
		&i.UserID,
	)
	return i, err
}

const listWorkouts = `-- name: ListWorkouts :many
SELECT id, duration, lifts, user_id FROM workout
WHERE user_id = UUID_TO_BIN(?)
`

func (q *Queries) ListWorkouts(ctx context.Context, userID string) ([]Workout, error) {
	rows, err := q.query(ctx, q.listWorkoutsStmt, listWorkouts, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Workout{}
	for rows.Next() {
		var i Workout
		if err := rows.Scan(
			&i.ID,
			&i.Duration,
			&i.Lifts,
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

const updateWorkout = `-- name: UpdateWorkout :execresult
UPDATE workout SET
duration = IFNULL(?, duration),
lifts = IFNULL(?, lifts)
WHERE id = ? AND user_id = UUID_TO_BIN(?)
`

type UpdateWorkoutParams struct {
	Duration interface{} `json:"duration"`
	Lifts    interface{} `json:"lifts"`
	ID       int32       `json:"id"`
	UserID   string      `json:"user_id"`
}

func (q *Queries) UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (sql.Result, error) {
	return q.exec(ctx, q.updateWorkoutStmt, updateWorkout,
		arg.Duration,
		arg.Lifts,
		arg.ID,
		arg.UserID,
	)
}

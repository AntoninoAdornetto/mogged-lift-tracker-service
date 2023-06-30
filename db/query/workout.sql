-- name: CreateWorkout :execresult
INSERT INTO workout (
	duration,
	lifts,
	user_id
) VALUES (
	?,
	?,
	UUID_TO_BIN(sqlc.arg('user_id'))
);

-- name: GetWorkout :one
SELECT * FROM workout
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: ListWorkouts :many
SELECT * FROM workout
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: UpdateWorkout :exec
UPDATE workout SET
	duration = COALESCE(sqlc.narg('duration'), duration),
	lifts = COALESCE(sqlc.narg('lifts'), lifts),
	completed_date = COALESCE(sqlc.narg('completed_date'), completed_date)
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: DeleteWorkout :exec
DELETE FROM workout 
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: GetTotalWorkouts :one
SELECT COUNT(*) FROM workout
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: GetLastWorkout :one
SELECT * FROM workout
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'))
ORDER BY completed_date DESC
LIMIT 1;

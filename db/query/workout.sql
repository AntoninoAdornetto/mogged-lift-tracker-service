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

-- name: UpdateWorkout :execresult
UPDATE workout SET
duration = IFNULL(sqlc.arg('duration'), duration),
lifts = IFNULL(sqlc.arg('lifts'), lifts)
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: DeleteWorkout :execresult
DELETE FROM workout 
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));
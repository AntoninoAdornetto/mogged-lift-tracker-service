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

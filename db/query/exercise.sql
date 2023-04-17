-- name: CreateExercise :execresult
INSERT INTO exercise (
	name,
	muscle_group,
	category,
	user_id
) VALUES (
	?,
	?,
	?,
	UUID_TO_BIN(sqlc.arg('user_id'))
);

-- name: GetExercise :one
SELECT * FROM exercise
WHERE id = ? LIMIT 1;

-- name: ListExercises :many
SELECT * FROM exercise
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

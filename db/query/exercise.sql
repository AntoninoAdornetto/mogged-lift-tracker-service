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
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id')) LIMIT 1;

-- name: GetExerciseByName :one
SELECT * FROM exercise
WHERE name = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: ListExercises :many
SELECT * FROM exercise
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: UpdateExercise :execresult
UPDATE exercise SET
	name = COALESCE(sqlc.narg('name'), name),
	muscle_group = COALESCE(sqlc.narg('muscle_group'), muscle_group),
	category = COALESCE(sqlc.narg('category'), category),
	most_weight_lifted = COALESCE(sqlc.narg('most_weight_lifted'), most_weight_lifted),
	most_reps_lifted = COALESCE(sqlc.narg('most_reps_lifted'), most_reps_lifted),
	rest_timer = COALESCE(sqlc.narg('rest_timer'), rest_timer)
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id')) AND id = sqlc.arg('id');

-- name: DeleteExercise :execresult
DELETE FROM exercise
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

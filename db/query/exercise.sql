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

-- name: UpdateExercise :execresult
UPDATE exercise SET
name = IFNULL(sqlc.arg('name'), name),
muscle_group = IFNULL(sqlc.arg('muscle_group'), muscle_group),
category = IFNULL(sqlc.arg('category'), category),
most_weight_lifted = IFNULL(sqlc.arg('most_weight_lifted'), most_weight_lifted),
most_reps_lifted = IFNULL(sqlc.arg('most_reps_lifted'), most_reps_lifted),
rest_timer = IFNULL(sqlc.arg('rest_timer'), rest_timer)
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id')) AND id = sqlc.arg('id');

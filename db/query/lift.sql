-- name: CreateLift :execresult
INSERT INTO lift (
	exercise_name,
	weight_lifted,
	reps,
	user_id,
	workout_id
) VALUES (
	?,
	?,
	?,
	UUID_TO_BIN(sqlc.arg('user_id')),
	?
);

-- name: GetLift :one
SELECT * FROM lift
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: ListLiftsFromWorkout :many
SELECT * FROM lift
WHERE workout_id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: ListMaxWeightPrs :many
SELECT * FROM lift
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'))
ORDER BY weight_lifted DESC LIMIT ?;

-- name: ListMaxRepPrs :many
SELECT * FROM lift
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'))
ORDER BY reps DESC LIMIT ?;

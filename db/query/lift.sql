-- name: CreateLift :execresult
INSERT INTO lift (
	exercise_name,
	weight_lifted,
	reps,
	set_type,
	user_id,
	workout_id
) VALUES (
	?,
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

-- @todo fix this, args are broken when using query for user_id
-- name: GetMaxLiftByExercise :one
SELECT MAX(weight_lifted) FROM lift
WHERE exercise_name = ? AND user_id = UUID_TO_BIN(sql.arg('user_id'));

-- @todo fix this, args are broken when using query and it should allow them to query by muscle group, not exercise name
-- name: GetMaxLiftsByMuscleGroup :many
SELECT muscle_group, exercise_name, weight_lifted, reps FROM lift
JOIN exercise ON lift.exercise_name = exercise.name
WHERE lift.user_id = UUID_TO_BIN(sql.arg('user_id'))
ORDER BY weight_lifted DESC;

-- name: ListMaxRepPrs :many
SELECT * FROM lift
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'))
ORDER BY reps DESC LIMIT ?;

-- name: UpdateLift :execresult
UPDATE lift set
exercise_name = IFNULL(sqlc.arg('exercise_name'), exercise_name),
weight_lifted = IFNULL(sqlc.arg('weight_lifted'), weight_lifted),
reps = IFNULL(sqlc.arg('reps'), reps)
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: DeleteLift :execresult
DELETE FROM lift
WHERE id = ? AND user_id = UUID_TO_BIN(sqlc.arg('user_id'));

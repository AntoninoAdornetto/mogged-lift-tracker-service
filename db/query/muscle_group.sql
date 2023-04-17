-- name: CreateMuscleGroup :execresult
INSERT INTO muscle_group (
	name
) VALUES (?);

-- name: GetMuscleGroup :one
SELECT * FROM muscle_group
WHERE id = ?;

-- name: ListMuscleGroups :many
SELECT * FROM muscle_group;

-- name: UpdateMuscleGroup :execresult
UPDATE muscle_group SET
name = ?
WHERE id = ?;

-- name: DeleteMuscleGroup :execresult
DELETE FROM muscle_group
WHERE id = ?;

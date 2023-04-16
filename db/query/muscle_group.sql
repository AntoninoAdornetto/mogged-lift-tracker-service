-- name: CreateMuscleGroup :execresult
INSERT INTO muscle_group (
	name
) VALUES (?);

-- name: GetMuscleGroup :one
SELECT * FROM muscle_group
WHERE id = ?;

-- name: ListMuscleGroups :many
SELECT * FROM muscle_group ORDER BY id;

-- name: UpdateMuscleGroup :execresult
UPDATE muscle_group SET
name = ?
WHERE id = ?;

-- name: DeleteMuscleGroup :execresult
DELETE FROM muscle_group
WHERE id = ?;

-- name: DeleteAllMuscleGroups :execresult
-- no API for this query, only for testing purposes
DELETE FROM muscle_group;

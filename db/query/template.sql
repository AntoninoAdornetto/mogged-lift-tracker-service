-- name: CreateTemplate :execresult
INSERT INTO template (
	name,
	exercises,
	created_by
) VALUES (
	?,
	?,
	UUID_TO_BIN(sqlc.arg('created_by'))
);

-- name: GetTemplate :one
SELECT * FROM template
WHERE id = ? AND created_by = UUID_TO_BIN(sqlc.arg('created_by'));

-- name: ListTemplates :many
SELECT * FROM template
WHERE created_by = UUID_TO_BIN(sqlc.arg('created_by'));

-- name: UpdateTemplate :exec
UPDATE template SET
	name = COALESCE(sqlc.narg('name'), name),
	exercises = COALESCE(sqlc.arg('exercises'), exercises),
	date_last_used = COALESCE(sqlc.narg('date_last_used'), date_last_used)
WHERE id = ? AND created_by = UUID_TO_BIN(sqlc.arg('created_by'));

-- name: DeleteTemplate :exec
DELETE FROM template
WHERE id = ? AND created_by = UUID_TO_BIN(sqlc.arg('created_by'));

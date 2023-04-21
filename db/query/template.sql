-- name: CreateTemplate :execresult
INSERT INTO template (
	name,
	lifts,
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

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

-- name: UpdateTemplate :execlastid
UPDATE template SET
name = IFNULL(sqlc.arg('name'), name),
lifts = IFNULL(sqlc.arg('lifts'), lifts),
date_last_used = IFNULL(sqlc.arg('date_last_used'), date_last_used)
WHERE id = ? AND created_by = UUID_TO_BIN(sqlc.arg('created_by'));

-- name: DeleteTemplate :execresult
DELETE FROM template
WHERE id = ? AND created_by = UUID_TO_BIN(sqlc.arg('created_by'));

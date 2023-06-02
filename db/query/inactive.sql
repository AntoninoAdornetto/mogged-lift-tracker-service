-- name: InsertInactiveUser :exec
INSERT INTO `inactive_user` (id) VALUES (UUID_TO_BIN(sqlc.arg('user_id')));

-- name: GetInactiveUser :one
SELECT BIN_TO_UUID(id) AS id FROM
`inactive_user` WHERE id = UUID_TO_BIN(sqlc.arg('user_id'));

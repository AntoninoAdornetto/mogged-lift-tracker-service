-- name: CreateSession :exec
INSERT INTO `session`(
	id,
	refresh_token,
	user_agent,
	client_ip,
	expires_at,
	user_id
) VALUES (
	UUID_TO_BIN(sqlc.arg('id')),
	?,
	?,
	?,
	?,
	UUID_TO_BIN(sqlc.arg('user_id'))
);

-- name: GetSession :one
SELECT 
BIN_TO_UUID(id) AS id,
BIN_TO_UUID(user_id) AS user_id,
refresh_token,
user_agent,
client_ip,
is_banned,
expires_at,
created_at
FROM session
WHERE id = UUID_TO_BIN(sqlc.arg('id'))
LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM session
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

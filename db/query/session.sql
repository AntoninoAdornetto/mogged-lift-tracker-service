-- name: CreateSession :exec
INSERT INTO `session`(
	refresh_token,
	user_agent,
	client_ip,
	expires_at,
	user_id
) VALUES (
	?,
	?,
	?,
	?,
	UUID_TO_BIN(sqlc.arg('user_id'))
);

-- name: GetSession :one
SELECT * FROM session
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'))
LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM session
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

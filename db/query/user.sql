-- name: CreateUser :exec
INSERT INTO user (
	id,
	first_name,
	last_name,
	email_address,
	password
) VALUES (UUID_TO_BIN(UUID()),?, ?, ?, ?);

-- name: GetUserByEmail :one
SELECT BIN_TO_UUID(id) as id, first_name, last_name, email_address, password_changed_at, password
FROM user WHERE email_address = ? LIMIT 1;

-- name: GetUserById :one
SELECT BIN_TO_UUID(id) as id, first_name, last_name, email_address, password_changed_at, password
FROM user WHERE id = UUID_TO_BIN(sqlc.arg('user_id')) LIMIT 1;

-- name: UpdateUser :exec
UPDATE user SET
	first_name = COALESCE(sqlc.narg('first_name'), first_name),
	last_name = COALESCE(sqlc.narg('last_name'), last_name),
	email_address = COALESCE(sqlc.narg('email_address'), email_address)
WHERE id = UUID_TO_BIN(sqlc.arg('user_id'));


-- name: ChangePassword :exec
UPDATE user SET
password = ?,
password_changed_at = NOW()
WHERE id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = UUID_TO_BIN(sqlc.arg('user_id'));

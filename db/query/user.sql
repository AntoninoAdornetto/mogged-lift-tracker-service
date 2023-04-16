-- name: CreateUser :execresult
INSERT INTO user (
	id,
	first_name,
	last_name,
	email_address,
	password
) VALUES (UUID_TO_BIN(UUID()),?, ?, ?, ?);

-- name: GetUser :one
SELECT BIN_TO_UUID(id) as id, first_name, last_name, email_address FROM user
WHERE email_address = ? LIMIT 1;


-- name: UpdateUser :execresult
UPDATE user SET
first_name = IFNULL(sqlc.narg('first_name'), first_name),
last_name = IFNULL(sqlc.narg('last_name'), last_name),
email_address = IFNULL(sqlc.narg('email_address'), email_address)
WHERE id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: DeleteUser :execresult
DELETE FROM user
WHERE id = UUID_TO_BIN(?);
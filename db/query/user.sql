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
WHERE id = UUID_TO_BIN(?);
-- @todo using this UUID to Bin function results in the field name being called "UUIDTOBIN"
-- look into sqlc naming parameters. sqlc.narg was resulting in the struct having an interface type for id instead of a string

-- name: DeleteUser :execresult
DELETE FROM user
WHERE id = UUID_TO_BIN(?);

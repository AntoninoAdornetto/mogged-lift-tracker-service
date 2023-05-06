-- name: CreateProfile :execlastid
INSERT INTO profile (
	country,
	measurement_system,
	body_weight,
	body_fat,
	timezone_offset,
	user_id
) VALUES (?, ?, ?, ?, ?, UUID_TO_BIN(sqlc.arg('user_id')));

-- name: GetProfile :one
SELECT * FROM profile
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id')) LIMIT 1;

-- name: UpdateProfile :execresult
UPDATE profile SET
	country = COALESCE(sqlc.narg('country'), country),
	measurement_system = COALESCE(sqlc.narg('measurement_system'), measurement_system),
	body_weight = COALESCE(sqlc.narg('body_weight'), body_weight),
	body_fat = COALESCE(sqlc.narg('body_fat'), body_fat),
	timezone_offset = COALESCE(sqlc.narg('timezone_offset'), timezone_offset)
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: DeleteProfile :exec
DELETE FROM profile
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

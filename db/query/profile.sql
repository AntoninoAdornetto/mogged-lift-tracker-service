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
country = IFNULL(sqlc.narg('country'), country),
measurement_system = IFNULL(sqlc.narg('measurement_system'), measurement_system),
body_weight = IFNULL(sqlc.narg('body_weight'), body_weight),
body_fat = IFNULL(sqlc.narg('body_fat'), body_fat),
timezone_offset = IFNULL(sqlc.narg('timezone_offset'), timezone_offset)
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

-- name: DeleteProfile :execresult
DELETE FROM profile
WHERE user_id = UUID_TO_BIN(sqlc.arg('user_id'));

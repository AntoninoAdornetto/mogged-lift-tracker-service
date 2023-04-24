-- name: CreateStockExercise :execresult
INSERT INTO stock_exercise (
	name,
	muscle_group,
	category
) VALUES (
	?,
	?,
	?
);

-- name: GetStockExercise :one
SELECT * FROM stock_exercise
WHERE id = ?;

-- name: ListStockExercies :many
SELECT * FROM stock_exercise;

-- name: UpdateStockExercise :execresult
UPDATE stock_exercise SET 
name = IFNULL(sqlc.arg('name'), name),
muscle_group = IFNULL(sqlc.arg('muscle_group'), muscle_group),
category = IFNULL(sqlc.arg('category'), category)
WHERE id = ?;

-- name: DeleteStockExercise :execresult
DELETE FROM stock_exercise WHERE
id = ?;
-- name: CreateCategory :execresult
INSERT INTO category (
	name
) VALUES (?);

-- name: GetCategory :one
SELECT * FROM category
WHERE id = ?;

-- name: ListCategories :many
SELECT * FROM category;

-- name: UpdateCategory :execresult
UPDATE category SET
name = ?
WHERE id = ?;

-- name: DeleteCategory :execresult
DELETE FROM category
WHERE id = ?;

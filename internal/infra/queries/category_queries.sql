-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateCategory :exec
UPDATE categories
SET name = $2, description = $3, updated_date = $4
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListCategories :many
SELECT * 
FROM categories 
WHERE deleted_at IS NULL 
ORDER BY name
LIMIT $1 OFFSET $2; 

-- name: GetCategoryProducts :many
SELECT p.*
FROM categories c
JOIN products p ON c.id = p.category_id
WHERE c.id = $1 AND c.deleted_at IS NULL AND p.deleted_at IS NULL 
ORDER BY p.name; 
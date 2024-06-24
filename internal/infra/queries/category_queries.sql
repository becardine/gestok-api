-- name: GetCategory :one
SELECT * FROM categories WHERE id = ? AND deleted_at IS NULL;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories
SET name = ?, description = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListCategories :many
SELECT * 
FROM categories 
WHERE deleted_at IS NULL 
ORDER BY name
LIMIT ? OFFSET ?; 

-- name: GetCategoryProducts :many
SELECT p.*
FROM categories c
JOIN products p ON c.id = p.category_id
WHERE c.id = ? AND c.deleted_at IS NULL AND p.deleted_at IS NULL 
ORDER BY p.name;

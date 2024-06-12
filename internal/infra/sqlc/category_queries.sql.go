// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: category_queries.sql

package db

import (
	"context"
	"database/sql"

	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id, name, description, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5)
`

type CreateCategoryParams struct {
	ID          common.ID
	Name        string
	Description sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id common.ID) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, description, deleted_at, created_date, updated_date FROM categories WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetCategory(ctx context.Context, id common.ID) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getCategoryProducts = `-- name: GetCategoryProducts :many
SELECT p.id, p.name, p.description, p.price, p.quantity_in_stock, p.image_url, p.category_id, p.brand_id, p.deleted_at, p.created_date, p.updated_date
FROM categories c
JOIN products p ON c.id = p.category_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY p.name
`

func (q *Queries) GetCategoryProducts(ctx context.Context, id common.ID) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getCategoryProducts, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.QuantityInStock,
			&i.ImageUrl,
			&i.CategoryID,
			&i.BrandID,
			&i.DeletedAt,
			&i.CreatedDate,
			&i.UpdatedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, description, deleted_at, created_date, updated_date FROM categories WHERE deleted_at IS NULL ORDER BY name
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.DeletedAt,
			&i.CreatedDate,
			&i.UpdatedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET name = $2, description = $3, updated_date = $4
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateCategoryParams struct {
	ID          common.ID
	Name        string
	Description sql.NullString
	UpdatedDate sql.NullTime
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.UpdatedDate,
	)
	return err
}
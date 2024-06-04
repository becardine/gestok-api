// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: brand_queries.sql

package db

import (
	"context"
	"database/sql"

	common "github.com/becardine/gestock-api/internal/entity/common"
)

const createBrand = `-- name: CreateBrand :exec
INSERT INTO brands (id, name, description, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5)
`

type CreateBrandParams struct {
	ID          common.ID
	Name        string
	Description sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

func (q *Queries) CreateBrand(ctx context.Context, arg CreateBrandParams) error {
	_, err := q.db.ExecContext(ctx, createBrand,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deleteBrand = `-- name: DeleteBrand :exec
UPDATE brands
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteBrand(ctx context.Context, id common.ID) error {
	_, err := q.db.ExecContext(ctx, deleteBrand, id)
	return err
}

const getBrand = `-- name: GetBrand :one
SELECT id, name, description, deleted_at, created_date, updated_date FROM brands WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetBrand(ctx context.Context, id common.ID) (Brand, error) {
	row := q.db.QueryRowContext(ctx, getBrand, id)
	var i Brand
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

const getBrandProducts = `-- name: GetBrandProducts :many
SELECT p.id, p.name, p.description, p.price, p.quantity_in_stock, p.image_url, p.category_id, p.brand_id, p.deleted_at, p.created_date, p.updated_date
FROM brands b
JOIN products p ON b.id = p.brand_id
WHERE b.id = $1 AND b.deleted_at IS NULL
ORDER BY p.name
`

func (q *Queries) GetBrandProducts(ctx context.Context, id common.ID) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getBrandProducts, id)
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

const listBrands = `-- name: ListBrands :many
SELECT id, name, description, deleted_at, created_date, updated_date FROM brands WHERE deleted_at IS NULL ORDER BY name
`

func (q *Queries) ListBrands(ctx context.Context) ([]Brand, error) {
	rows, err := q.db.QueryContext(ctx, listBrands)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Brand
	for rows.Next() {
		var i Brand
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

const updateBrand = `-- name: UpdateBrand :exec
UPDATE brands
SET name = $2, description = $3, updated_date = $4
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateBrandParams struct {
	ID          common.ID
	Name        string
	Description sql.NullString
	UpdatedDate sql.NullTime
}

func (q *Queries) UpdateBrand(ctx context.Context, arg UpdateBrandParams) error {
	_, err := q.db.ExecContext(ctx, updateBrand,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.UpdatedDate,
	)
	return err
}

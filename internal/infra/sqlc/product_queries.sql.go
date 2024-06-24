// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: product_queries.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addProductStock = `-- name: AddProductStock :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE stock_id = stock_id
`

type AddProductStockParams struct {
	StockID   uuid.UUID
	ProductID uuid.UUID
}

func (q *Queries) AddProductStock(ctx context.Context, arg AddProductStockParams) error {
	_, err := q.db.ExecContext(ctx, addProductStock, arg.StockID, arg.ProductID)
	return err
}

const createProduct = `-- name: CreateProduct :exec
INSERT INTO products (id, name, description, price, quantity_in_stock, image_url, category_id, brand_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateProductParams struct {
	ID              uuid.UUID
	Name            string
	Description     sql.NullString
	Price           float64
	QuantityInStock int32
	ImageUrl        sql.NullString
	CategoryID      uuid.UUID
	BrandID         uuid.UUID
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.QuantityInStock,
		arg.ImageUrl,
		arg.CategoryID,
		arg.BrandID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, price, quantity_in_stock, image_url, category_id, brand_id, deleted_at, created_at, updated_at FROM products WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.QuantityInStock,
		&i.ImageUrl,
		&i.CategoryID,
		&i.BrandID,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductStocks = `-- name: GetProductStocks :many
SELECT s.id, s.name, s.location, s.capacity, s.deleted_at, s.created_at, s.updated_at
FROM products p
JOIN product_stocks ps ON p.id = ps.product_id
JOIN stocks s ON ps.stock_id = s.id
WHERE p.id = ? AND p.deleted_at IS NULL AND s.deleted_at IS NULL
ORDER BY s.name
LIMIT ? OFFSET ?
`

type GetProductStocksParams struct {
	ID     uuid.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) GetProductStocks(ctx context.Context, arg GetProductStocksParams) ([]Stock, error) {
	rows, err := q.db.QueryContext(ctx, getProductStocks, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stock
	for rows.Next() {
		var i Stock
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.Capacity,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, price, quantity_in_stock, image_url, category_id, brand_id, deleted_at, created_at, updated_at
FROM products
WHERE deleted_at IS NULL
ORDER BY name
LIMIT ? OFFSET ?
`

type ListProductsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
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
			&i.CreatedAt,
			&i.UpdatedAt,
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

const removeProductStock = `-- name: RemoveProductStock :execresult
DELETE FROM product_stocks
WHERE stock_id = ? AND product_id = ?
`

type RemoveProductStockParams struct {
	StockID   uuid.UUID
	ProductID uuid.UUID
}

func (q *Queries) RemoveProductStock(ctx context.Context, arg RemoveProductStockParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, removeProductStock, arg.StockID, arg.ProductID)
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET name = ?, description = ?, price = ?, quantity_in_stock = ?, image_url = ?, category_id = ?, brand_id = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdateProductParams struct {
	Name            string
	Description     sql.NullString
	Price           float64
	QuantityInStock int32
	ImageUrl        sql.NullString
	CategoryID      uuid.UUID
	BrandID         uuid.UUID
	UpdatedAt       sql.NullTime
	ID              uuid.UUID
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.QuantityInStock,
		arg.ImageUrl,
		arg.CategoryID,
		arg.BrandID,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: stock_queries.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addStockProduct = `-- name: AddStockProduct :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES ($1, $2)
ON CONFLICT (stock_id, product_id) DO NOTHING
`

type AddStockProductParams struct {
	StockID   uuid.NullUUID
	ProductID uuid.NullUUID
}

func (q *Queries) AddStockProduct(ctx context.Context, arg AddStockProductParams) error {
	_, err := q.db.ExecContext(ctx, addStockProduct, arg.StockID, arg.ProductID)
	return err
}

const createStock = `-- name: CreateStock :exec
INSERT INTO stocks (id, name, location, capacity, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateStockParams struct {
	ID          uuid.UUID
	Name        string
	Location    sql.NullString
	Capacity    int32
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

func (q *Queries) CreateStock(ctx context.Context, arg CreateStockParams) error {
	_, err := q.db.ExecContext(ctx, createStock,
		arg.ID,
		arg.Name,
		arg.Location,
		arg.Capacity,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deleteStock = `-- name: DeleteStock :exec
UPDATE stocks
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteStock(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteStock, id)
	return err
}

const getStock = `-- name: GetStock :one
SELECT id, name, location, capacity, deleted_at, created_date, updated_date FROM stocks WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetStock(ctx context.Context, id uuid.UUID) (Stock, error) {
	row := q.db.QueryRowContext(ctx, getStock, id)
	var i Stock
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.Capacity,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getStockProducts = `-- name: GetStockProducts :many

SELECT p.id, p.name, p.description, p.price, p.quantity_in_stock, p.image_url, p.category_id, p.brand_id, p.deleted_at, p.created_date, p.updated_date
FROM stocks s
         JOIN product_stocks ps ON s.id = ps.stock_id
         JOIN products p ON ps.product_id = p.id
WHERE s.id = $1 AND s.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name
`

// Paginação
func (q *Queries) GetStockProducts(ctx context.Context, id uuid.UUID) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getStockProducts, id)
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

const listStocks = `-- name: ListStocks :many
SELECT id, name, location, capacity, deleted_at, created_date, updated_date
FROM stocks
WHERE deleted_at IS NULL
ORDER BY name
LIMIT $1 OFFSET $2
`

type ListStocksParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListStocks(ctx context.Context, arg ListStocksParams) ([]Stock, error) {
	rows, err := q.db.QueryContext(ctx, listStocks, arg.Limit, arg.Offset)
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

const removeStockProduct = `-- name: RemoveStockProduct :execresult
DELETE FROM product_stocks
WHERE stock_id = $1 AND product_id = $2
RETURNING id, product_id, stock_id, quantity, created_at, updated_at, deleted_at
`

type RemoveStockProductParams struct {
	StockID   uuid.NullUUID
	ProductID uuid.NullUUID
}

func (q *Queries) RemoveStockProduct(ctx context.Context, arg RemoveStockProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, removeStockProduct, arg.StockID, arg.ProductID)
}

const updateStock = `-- name: UpdateStock :exec
UPDATE stocks
SET name = $2, location = $3, capacity = $4, updated_date = $5
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateStockParams struct {
	ID          uuid.UUID
	Name        string
	Location    sql.NullString
	Capacity    int32
	UpdatedDate sql.NullTime
}

func (q *Queries) UpdateStock(ctx context.Context, arg UpdateStockParams) error {
	_, err := q.db.ExecContext(ctx, updateStock,
		arg.ID,
		arg.Name,
		arg.Location,
		arg.Capacity,
		arg.UpdatedDate,
	)
	return err
}

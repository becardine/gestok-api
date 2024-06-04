// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: product_stocks_queries.sql

package db

import (
	"context"

	common "github.com/becardine/gestock-api/internal/entity/common"
)

const createProductStock = `-- name: CreateProductStock :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES ($1, $2)
`

type CreateProductStockParams struct {
	StockID   common.ID
	ProductID common.ID
}

func (q *Queries) CreateProductStock(ctx context.Context, arg CreateProductStockParams) error {
	_, err := q.db.ExecContext(ctx, createProductStock, arg.StockID, arg.ProductID)
	return err
}

const deleteProductStock = `-- name: DeleteProductStock :exec
DELETE FROM product_stocks WHERE stock_id = $1 AND product_id = $2
`

type DeleteProductStockParams struct {
	StockID   common.ID
	ProductID common.ID
}

func (q *Queries) DeleteProductStock(ctx context.Context, arg DeleteProductStockParams) error {
	_, err := q.db.ExecContext(ctx, deleteProductStock, arg.StockID, arg.ProductID)
	return err
}

const getProductStock = `-- name: GetProductStock :one
SELECT id, stock_id, product_id, deleted_at, created_date, updated_date FROM product_stocks WHERE stock_id = $1 AND product_id = $2
`

type GetProductStockParams struct {
	StockID   common.ID
	ProductID common.ID
}

func (q *Queries) GetProductStock(ctx context.Context, arg GetProductStockParams) (ProductStock, error) {
	row := q.db.QueryRowContext(ctx, getProductStock, arg.StockID, arg.ProductID)
	var i ProductStock
	err := row.Scan(
		&i.ID,
		&i.StockID,
		&i.ProductID,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getProductStocks = `-- name: GetProductStocks :many
SELECT id, stock_id, product_id, deleted_at, created_date, updated_date FROM product_stocks WHERE stock_id = $1
`

func (q *Queries) GetProductStocks(ctx context.Context, stockID common.ID) ([]ProductStock, error) {
	rows, err := q.db.QueryContext(ctx, getProductStocks, stockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductStock
	for rows.Next() {
		var i ProductStock
		if err := rows.Scan(
			&i.ID,
			&i.StockID,
			&i.ProductID,
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

const getProductStocksByProductId = `-- name: GetProductStocksByProductId :many
SELECT id, stock_id, product_id, deleted_at, created_date, updated_date FROM product_stocks WHERE product_id = $1
`

func (q *Queries) GetProductStocksByProductId(ctx context.Context, productID common.ID) ([]ProductStock, error) {
	rows, err := q.db.QueryContext(ctx, getProductStocksByProductId, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductStock
	for rows.Next() {
		var i ProductStock
		if err := rows.Scan(
			&i.ID,
			&i.StockID,
			&i.ProductID,
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

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: order_item_queries.sql

package db

import (
	"context"
	"database/sql"
)

const createOrderItem = `-- name: CreateOrderItem :exec
INSERT INTO order_items (
    id, order_id, product_id, quantity, unit_price, created_at, updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateOrderItemParams struct {
	ID        string
	OrderID   sql.NullString
	ProductID sql.NullString
	Quantity  int32
	UnitPrice float64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) error {
	_, err := q.db.ExecContext(ctx, createOrderItem,
		arg.ID,
		arg.OrderID,
		arg.ProductID,
		arg.Quantity,
		arg.UnitPrice,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteOrderItem = `-- name: DeleteOrderItem :exec
UPDATE order_items
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) DeleteOrderItem(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteOrderItem, id)
	return err
}

const getOrderItem = `-- name: GetOrderItem :one
SELECT id, order_id, product_id, quantity, unit_price, deleted_at, created_at, updated_at FROM order_items WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetOrderItem(ctx context.Context, id string) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderItem, id)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Quantity,
		&i.UnitPrice,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOrderItems = `-- name: ListOrderItems :many
SELECT id, order_id, product_id, quantity, unit_price, deleted_at, created_at, updated_at 
FROM order_items 
WHERE order_id = ? AND deleted_at IS NULL 
ORDER BY created_at
LIMIT 2 OFFSET 3
`

func (q *Queries) ListOrderItems(ctx context.Context, orderID sql.NullString) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, listOrderItems, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Quantity,
			&i.UnitPrice,
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

const updateOrderItem = `-- name: UpdateOrderItem :exec
UPDATE order_items
SET quantity = ?, unit_price = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdateOrderItemParams struct {
	Quantity  int32
	UnitPrice float64
	UpdatedAt sql.NullTime
	ID        string
}

func (q *Queries) UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) error {
	_, err := q.db.ExecContext(ctx, updateOrderItem,
		arg.Quantity,
		arg.UnitPrice,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

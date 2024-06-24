// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: delivery_queries.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createDelivery = `-- name: CreateDelivery :exec
INSERT INTO deliveries (id, order_id, customer_id, delivery_type, delivery_at, delivery_status, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateDeliveryParams struct {
	ID             uuid.UUID
	OrderID        uuid.UUID
	CustomerID     uuid.UUID
	DeliveryType   string
	DeliveryAt     sql.NullTime
	DeliveryStatus string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

func (q *Queries) CreateDelivery(ctx context.Context, arg CreateDeliveryParams) error {
	_, err := q.db.ExecContext(ctx, createDelivery,
		arg.ID,
		arg.OrderID,
		arg.CustomerID,
		arg.DeliveryType,
		arg.DeliveryAt,
		arg.DeliveryStatus,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteDelivery = `-- name: DeleteDelivery :exec
UPDATE deliveries
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) DeleteDelivery(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDelivery, id)
	return err
}

const getDeliveriesByCustomerId = `-- name: GetDeliveriesByCustomerId :many
SELECT id, order_id, customer_id, delivery_type, delivery_at, delivery_status, deleted_at, created_at, updated_at
FROM deliveries
WHERE customer_id = ? AND deleted_at IS NULL
ORDER BY delivery_at DESC
LIMIT 2 OFFSET 3
`

func (q *Queries) GetDeliveriesByCustomerId(ctx context.Context, customerID uuid.UUID) ([]Delivery, error) {
	rows, err := q.db.QueryContext(ctx, getDeliveriesByCustomerId, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Delivery
	for rows.Next() {
		var i Delivery
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.DeliveryType,
			&i.DeliveryAt,
			&i.DeliveryStatus,
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

const getDeliveriesByOrderId = `-- name: GetDeliveriesByOrderId :many
SELECT id, order_id, customer_id, delivery_type, delivery_at, delivery_status, deleted_at, created_at, updated_at
FROM deliveries
WHERE order_id = ? AND deleted_at IS NULL
ORDER BY delivery_at DESC
LIMIT 2 OFFSET 3
`

func (q *Queries) GetDeliveriesByOrderId(ctx context.Context, orderID uuid.UUID) ([]Delivery, error) {
	rows, err := q.db.QueryContext(ctx, getDeliveriesByOrderId, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Delivery
	for rows.Next() {
		var i Delivery
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.DeliveryType,
			&i.DeliveryAt,
			&i.DeliveryStatus,
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

const getDelivery = `-- name: GetDelivery :one
SELECT id, order_id, customer_id, delivery_type, delivery_at, delivery_status, deleted_at, created_at, updated_at FROM deliveries WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetDelivery(ctx context.Context, id uuid.UUID) (Delivery, error) {
	row := q.db.QueryRowContext(ctx, getDelivery, id)
	var i Delivery
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.CustomerID,
		&i.DeliveryType,
		&i.DeliveryAt,
		&i.DeliveryStatus,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listDeliveries = `-- name: ListDeliveries :many
SELECT id, order_id, customer_id, delivery_type, delivery_at, delivery_status, deleted_at, created_at, updated_at
FROM deliveries
WHERE deleted_at IS NULL
ORDER BY delivery_at DESC
LIMIT 1 OFFSET 2
`

func (q *Queries) ListDeliveries(ctx context.Context) ([]Delivery, error) {
	rows, err := q.db.QueryContext(ctx, listDeliveries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Delivery
	for rows.Next() {
		var i Delivery
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.DeliveryType,
			&i.DeliveryAt,
			&i.DeliveryStatus,
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

const updateDelivery = `-- name: UpdateDelivery :exec
UPDATE deliveries
SET delivery_type = ?, delivery_at = ?, delivery_status = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdateDeliveryParams struct {
	DeliveryType   string
	DeliveryAt     sql.NullTime
	DeliveryStatus string
	UpdatedAt      sql.NullTime
	ID             uuid.UUID
}

func (q *Queries) UpdateDelivery(ctx context.Context, arg UpdateDeliveryParams) error {
	_, err := q.db.ExecContext(ctx, updateDelivery,
		arg.DeliveryType,
		arg.DeliveryAt,
		arg.DeliveryStatus,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

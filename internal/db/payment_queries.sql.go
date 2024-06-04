// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: payment_queries.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPayment = `-- name: CreatePayment :exec
INSERT INTO payments (id, order_id, customer_id, payment_type, payment_date, payment_value, payment_status, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type CreatePaymentParams struct {
	ID            uuid.UUID
	OrderID       uuid.NullUUID
	CustomerID    uuid.NullUUID
	PaymentType   string
	PaymentDate   sql.NullTime
	PaymentValue  string
	PaymentStatus string
	CreatedDate   sql.NullTime
	UpdatedDate   sql.NullTime
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, createPayment,
		arg.ID,
		arg.OrderID,
		arg.CustomerID,
		arg.PaymentType,
		arg.PaymentDate,
		arg.PaymentValue,
		arg.PaymentStatus,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deletePayment = `-- name: DeletePayment :exec
UPDATE payments
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeletePayment(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePayment, id)
	return err
}

const getPayment = `-- name: GetPayment :one
SELECT id, order_id, customer_id, payment_type, payment_date, payment_value, payment_status, deleted_at, created_date, updated_date FROM payments WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetPayment(ctx context.Context, id uuid.UUID) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.CustomerID,
		&i.PaymentType,
		&i.PaymentDate,
		&i.PaymentValue,
		&i.PaymentStatus,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getPaymentsByCustomerId = `-- name: GetPaymentsByCustomerId :many
SELECT id, order_id, customer_id, payment_type, payment_date, payment_value, payment_status, deleted_at, created_date, updated_date
FROM payments
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY payment_date DESC
`

func (q *Queries) GetPaymentsByCustomerId(ctx context.Context, customerID uuid.NullUUID) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByCustomerId, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.PaymentType,
			&i.PaymentDate,
			&i.PaymentValue,
			&i.PaymentStatus,
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

const getPaymentsByOrderId = `-- name: GetPaymentsByOrderId :many
SELECT id, order_id, customer_id, payment_type, payment_date, payment_value, payment_status, deleted_at, created_date, updated_date
FROM payments
WHERE order_id = $1 AND deleted_at IS NULL
ORDER BY payment_date DESC
`

func (q *Queries) GetPaymentsByOrderId(ctx context.Context, orderID uuid.NullUUID) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByOrderId, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.PaymentType,
			&i.PaymentDate,
			&i.PaymentValue,
			&i.PaymentStatus,
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

const listPayments = `-- name: ListPayments :many
SELECT id, order_id, customer_id, payment_type, payment_date, payment_value, payment_status, deleted_at, created_date, updated_date FROM payments WHERE deleted_at IS NULL ORDER BY payment_date DESC
`

func (q *Queries) ListPayments(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listPayments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.PaymentType,
			&i.PaymentDate,
			&i.PaymentValue,
			&i.PaymentStatus,
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

const updatePayment = `-- name: UpdatePayment :exec
UPDATE payments
SET order_id = $2, customer_id = $3, payment_type = $4, payment_date = $5, payment_value = $6, payment_status = $7, updated_date = $8
WHERE id = $1 AND deleted_at IS NULL
`

type UpdatePaymentParams struct {
	ID            uuid.UUID
	OrderID       uuid.NullUUID
	CustomerID    uuid.NullUUID
	PaymentType   string
	PaymentDate   sql.NullTime
	PaymentValue  string
	PaymentStatus string
	UpdatedDate   sql.NullTime
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, updatePayment,
		arg.ID,
		arg.OrderID,
		arg.CustomerID,
		arg.PaymentType,
		arg.PaymentDate,
		arg.PaymentValue,
		arg.PaymentStatus,
		arg.UpdatedDate,
	)
	return err
}

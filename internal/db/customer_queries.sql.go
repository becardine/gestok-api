// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: customer_queries.sql

package db

import (
	"context"
	"database/sql"

	common "github.com/becardine/gestock-api/internal/entity/common"
)

const createCustomer = `-- name: CreateCustomer :exec
INSERT INTO customers (id, name, email, password, address, phone, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreateCustomerParams struct {
	ID          common.ID
	Name        string
	Email       string
	Password    string
	Address     sql.NullString
	Phone       sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) error {
	_, err := q.db.ExecContext(ctx, createCustomer,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Address,
		arg.Phone,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
UPDATE customers
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id common.ID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, name, email, password, address, phone, deleted_at, created_date, updated_date FROM customers WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetCustomer(ctx context.Context, id common.ID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Address,
		&i.Phone,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getCustomerDeliveries = `-- name: GetCustomerDeliveries :many
SELECT d.id, d.order_id, d.customer_id, d.delivery_type, d.delivery_date, d.delivery_status, d.deleted_at, d.created_date, d.updated_date
FROM customers c
JOIN deliveries d ON c.id = d.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY d.delivery_date DESC
`

func (q *Queries) GetCustomerDeliveries(ctx context.Context, id common.ID) ([]Delivery, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerDeliveries, id)
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
			&i.DeliveryDate,
			&i.DeliveryStatus,
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

const getCustomerFeedbacks = `-- name: GetCustomerFeedbacks :many
SELECT f.id, f.customer_id, f.order_id, f.rating, f.comment, f.deleted_at, f.created_date, f.updated_date
FROM customers c
JOIN feedbacks f ON c.id = f.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY f.created_date DESC
`

func (q *Queries) GetCustomerFeedbacks(ctx context.Context, id common.ID) ([]Feedback, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerFeedbacks, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feedback
	for rows.Next() {
		var i Feedback
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.OrderID,
			&i.Rating,
			&i.Comment,
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

const getCustomerOrders = `-- name: GetCustomerOrders :many
SELECT o.id, o.customer_id, o.order_date, o.order_status, o.total_value, o.deleted_at, o.created_date, o.updated_date
FROM customers c
JOIN orders o ON c.id = o.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY o.order_date DESC
`

func (q *Queries) GetCustomerOrders(ctx context.Context, id common.ID) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerOrders, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.OrderDate,
			&i.OrderStatus,
			&i.TotalValue,
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

const getCustomerPayments = `-- name: GetCustomerPayments :many
SELECT p.id, p.order_id, p.customer_id, p.payment_type, p.payment_date, p.payment_value, p.payment_status, p.deleted_at, p.created_date, p.updated_date
FROM customers c
JOIN payments p ON c.id = p.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY p.payment_date DESC
`

func (q *Queries) GetCustomerPayments(ctx context.Context, id common.ID) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerPayments, id)
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

const listCustomers = `-- name: ListCustomers :many
SELECT id, name, email, password, address, phone, deleted_at, created_date, updated_date FROM customers WHERE deleted_at IS NULL ORDER BY name
`

func (q *Queries) ListCustomers(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, listCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Address,
			&i.Phone,
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

const updateCustomer = `-- name: UpdateCustomer :exec
UPDATE customers
SET name = $2, email = $3, password = $4, address = $5, phone = $6, updated_date = $7
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateCustomerParams struct {
	ID          common.ID
	Name        string
	Email       string
	Password    string
	Address     sql.NullString
	Phone       sql.NullString
	UpdatedDate sql.NullTime
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) error {
	_, err := q.db.ExecContext(ctx, updateCustomer,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Address,
		arg.Phone,
		arg.UpdatedDate,
	)
	return err
}

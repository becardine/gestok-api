// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: order_queries.sql

package db

import (
	"context"
	"database/sql"

	common "github.com/becardine/gestock-api/internal/entity/common"
	"github.com/google/uuid"
)

const addOrderProduct = `-- name: AddOrderProduct :exec
INSERT INTO order_products (order_id, product_id, quantity, unit_price)
VALUES ($1, $2, $3, $4)
`

type AddOrderProductParams struct {
	OrderID   uuid.NullUUID
	ProductID uuid.NullUUID
	Quantity  int32
	UnitPrice string
}

func (q *Queries) AddOrderProduct(ctx context.Context, arg AddOrderProductParams) error {
	_, err := q.db.ExecContext(ctx, addOrderProduct,
		arg.OrderID,
		arg.ProductID,
		arg.Quantity,
		arg.UnitPrice,
	)
	return err
}

const createOrder = `-- name: CreateOrder :exec
INSERT INTO orders (id, customer_id, order_date, order_status, total_value, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateOrderParams struct {
	ID          common.ID
	CustomerID  uuid.NullUUID
	OrderDate   sql.NullTime
	OrderStatus string
	TotalValue  string
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.ExecContext(ctx, createOrder,
		arg.ID,
		arg.CustomerID,
		arg.OrderDate,
		arg.OrderStatus,
		arg.TotalValue,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deleteOrder = `-- name: DeleteOrder :exec
UPDATE orders
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id common.ID) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT id, customer_id, order_date, order_status, total_value, deleted_at, created_date, updated_date FROM orders WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetOrder(ctx context.Context, id common.ID) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.OrderDate,
		&i.OrderStatus,
		&i.TotalValue,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getOrderByCustomerId = `-- name: GetOrderByCustomerId :many
SELECT id, customer_id, order_date, order_status, total_value, deleted_at, created_date, updated_date
FROM orders
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY order_date DESC
`

func (q *Queries) GetOrderByCustomerId(ctx context.Context, customerID uuid.NullUUID) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrderByCustomerId, customerID)
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

const getOrderByProductId = `-- name: GetOrderByProductId :many
SELECT o.id, o.customer_id, o.order_date, o.order_status, o.total_value, o.deleted_at, o.created_date, o.updated_date
FROM orders o
JOIN order_products op ON o.id = op.order_id
WHERE op.product_id = $1 AND o.deleted_at IS NULL
ORDER BY o.order_date DESC
`

func (q *Queries) GetOrderByProductId(ctx context.Context, productID uuid.NullUUID) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrderByProductId, productID)
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

const getOrderProducts = `-- name: GetOrderProducts :many
SELECT p.id, p.name, p.description, p.price, p.quantity_in_stock, p.image_url, p.category_id, p.brand_id, p.deleted_at, p.created_date, p.updated_date 
FROM orders o
JOIN order_products op ON o.id = op.order_id
JOIN products p ON op.product_id = p.id
WHERE o.id = $1 AND o.deleted_at IS NULL
`

func (q *Queries) GetOrderProducts(ctx context.Context, id common.ID) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getOrderProducts, id)
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

const listOrders = `-- name: ListOrders :many
SELECT id, customer_id, order_date, order_status, total_value, deleted_at, created_date, updated_date FROM orders WHERE deleted_at IS NULL ORDER BY order_date DESC
`

func (q *Queries) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders)
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

const removeOrderProduct = `-- name: RemoveOrderProduct :exec
DELETE FROM order_products WHERE order_id = $1 AND product_id = $2
`

type RemoveOrderProductParams struct {
	OrderID   uuid.NullUUID
	ProductID uuid.NullUUID
}

func (q *Queries) RemoveOrderProduct(ctx context.Context, arg RemoveOrderProductParams) error {
	_, err := q.db.ExecContext(ctx, removeOrderProduct, arg.OrderID, arg.ProductID)
	return err
}

const updateOrder = `-- name: UpdateOrder :exec
UPDATE orders
SET customer_id = $2, order_date = $3, order_status = $4, total_value = $5, updated_date = $6
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateOrderParams struct {
	ID          common.ID
	CustomerID  uuid.NullUUID
	OrderDate   sql.NullTime
	OrderStatus string
	TotalValue  string
	UpdatedDate sql.NullTime
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateOrder,
		arg.ID,
		arg.CustomerID,
		arg.OrderDate,
		arg.OrderStatus,
		arg.TotalValue,
		arg.UpdatedDate,
	)
	return err
}

const updateOrderStatus = `-- name: UpdateOrderStatus :exec
UPDATE orders
SET order_status = $2
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateOrderStatusParams struct {
	ID          common.ID
	OrderStatus string
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateOrderStatus, arg.ID, arg.OrderStatus)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feedback_queries.sql

package db

import (
	"context"
	"database/sql"

	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/google/uuid"
)

const createFeedback = `-- name: CreateFeedback :exec
INSERT INTO feedbacks (id, customer_id, order_id, rating, comment, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateFeedbackParams struct {
	ID          common.ID
	CustomerID  uuid.NullUUID
	OrderID     uuid.NullUUID
	Rating      sql.NullInt32
	Comment     sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

func (q *Queries) CreateFeedback(ctx context.Context, arg CreateFeedbackParams) error {
	_, err := q.db.ExecContext(ctx, createFeedback,
		arg.ID,
		arg.CustomerID,
		arg.OrderID,
		arg.Rating,
		arg.Comment,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deleteFeedback = `-- name: DeleteFeedback :exec
UPDATE feedbacks
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteFeedback(ctx context.Context, id common.ID) error {
	_, err := q.db.ExecContext(ctx, deleteFeedback, id)
	return err
}

const getFeedback = `-- name: GetFeedback :one
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_date, updated_date FROM feedbacks WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetFeedback(ctx context.Context, id common.ID) (Feedback, error) {
	row := q.db.QueryRowContext(ctx, getFeedback, id)
	var i Feedback
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.OrderID,
		&i.Rating,
		&i.Comment,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getFeedbackByCustomerId = `-- name: GetFeedbackByCustomerId :many
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_date, updated_date
FROM feedbacks
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY created_date DESC
`

func (q *Queries) GetFeedbackByCustomerId(ctx context.Context, customerID uuid.NullUUID) ([]Feedback, error) {
	rows, err := q.db.QueryContext(ctx, getFeedbackByCustomerId, customerID)
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

const getFeedbackByOrderId = `-- name: GetFeedbackByOrderId :many
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_date, updated_date
FROM feedbacks
WHERE order_id = $1 AND deleted_at IS NULL
ORDER BY created_date DESC
`

func (q *Queries) GetFeedbackByOrderId(ctx context.Context, orderID uuid.NullUUID) ([]Feedback, error) {
	rows, err := q.db.QueryContext(ctx, getFeedbackByOrderId, orderID)
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

const listFeedbacks = `-- name: ListFeedbacks :many
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_date, updated_date FROM feedbacks WHERE deleted_at IS NULL ORDER BY created_date DESC
`

func (q *Queries) ListFeedbacks(ctx context.Context) ([]Feedback, error) {
	rows, err := q.db.QueryContext(ctx, listFeedbacks)
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

const updateFeedback = `-- name: UpdateFeedback :exec
UPDATE feedbacks
SET customer_id = $2, order_id = $3, rating = $4, comment = $5, updated_date = $6
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateFeedbackParams struct {
	ID          common.ID
	CustomerID  uuid.NullUUID
	OrderID     uuid.NullUUID
	Rating      sql.NullInt32
	Comment     sql.NullString
	UpdatedDate sql.NullTime
}

func (q *Queries) UpdateFeedback(ctx context.Context, arg UpdateFeedbackParams) error {
	_, err := q.db.ExecContext(ctx, updateFeedback,
		arg.ID,
		arg.CustomerID,
		arg.OrderID,
		arg.Rating,
		arg.Comment,
		arg.UpdatedDate,
	)
	return err
}
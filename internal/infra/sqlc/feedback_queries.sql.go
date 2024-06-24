// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feedback_queries.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createFeedback = `-- name: CreateFeedback :exec
INSERT INTO feedbacks (id, customer_id, order_id, rating, comment, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateFeedbackParams struct {
	ID         uuid.UUID
	CustomerID uuid.UUID
	OrderID    uuid.UUID
	Rating     sql.NullInt32
	Comment    sql.NullString
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}

func (q *Queries) CreateFeedback(ctx context.Context, arg CreateFeedbackParams) error {
	_, err := q.db.ExecContext(ctx, createFeedback,
		arg.ID,
		arg.CustomerID,
		arg.OrderID,
		arg.Rating,
		arg.Comment,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteFeedback = `-- name: DeleteFeedback :exec
UPDATE feedbacks
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) DeleteFeedback(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFeedback, id)
	return err
}

const getFeedback = `-- name: GetFeedback :one
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_at, updated_at FROM feedbacks WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetFeedback(ctx context.Context, id uuid.UUID) (Feedback, error) {
	row := q.db.QueryRowContext(ctx, getFeedback, id)
	var i Feedback
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.OrderID,
		&i.Rating,
		&i.Comment,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeedbackByCustomerId = `-- name: GetFeedbackByCustomerId :many
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_at, updated_at
FROM feedbacks
WHERE customer_id = ? AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 2 OFFSET 3
`

func (q *Queries) GetFeedbackByCustomerId(ctx context.Context, customerID uuid.UUID) ([]Feedback, error) {
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

const getFeedbackByOrderId = `-- name: GetFeedbackByOrderId :many
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_at, updated_at
FROM feedbacks
WHERE order_id = ? AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 2 OFFSET 3
`

func (q *Queries) GetFeedbackByOrderId(ctx context.Context, orderID uuid.UUID) ([]Feedback, error) {
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

const listFeedbacks = `-- name: ListFeedbacks :many
SELECT id, customer_id, order_id, rating, comment, deleted_at, created_at, updated_at
FROM feedbacks
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 1 OFFSET 2
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

const updateFeedback = `-- name: UpdateFeedback :exec
UPDATE feedbacks
SET rating = ?, comment = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdateFeedbackParams struct {
	Rating    sql.NullInt32
	Comment   sql.NullString
	UpdatedAt sql.NullTime
	ID        uuid.UUID
}

func (q *Queries) UpdateFeedback(ctx context.Context, arg UpdateFeedbackParams) error {
	_, err := q.db.ExecContext(ctx, updateFeedback,
		arg.Rating,
		arg.Comment,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

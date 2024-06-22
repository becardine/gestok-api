-- name: GetFeedback :one
SELECT * FROM feedbacks WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateFeedback :exec
INSERT INTO feedbacks (id, customer_id, order_id, rating, comment, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateFeedback :exec
UPDATE feedbacks
SET rating = $2, comment = $3, updated_date = $4
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteFeedback :exec
UPDATE feedbacks
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListFeedbacks :many
SELECT *
FROM feedbacks
WHERE deleted_at IS NULL
ORDER BY created_date DESC
    LIMIT $1 OFFSET $2;

-- name: GetFeedbackByOrderId :many
SELECT *
FROM feedbacks
WHERE order_id = $1 AND deleted_at IS NULL
ORDER BY created_date DESC
    LIMIT $2 OFFSET $3;

-- name: GetFeedbackByCustomerId :many
SELECT *
FROM feedbacks
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY created_date DESC
    LIMIT $2 OFFSET $3;
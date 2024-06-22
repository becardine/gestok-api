-- name: GetFeedback :one
SELECT * FROM feedbacks WHERE id = ? AND deleted_at IS NULL;

-- name: CreateFeedback :exec
INSERT INTO feedbacks (id, customer_id, order_id, rating, comment, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateFeedback :exec
UPDATE feedbacks
SET rating = ?, comment = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteFeedback :exec
UPDATE feedbacks
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListFeedbacks :many
SELECT *
FROM feedbacks
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 1 OFFSET 2;

-- name: GetFeedbackByOrderId :many
SELECT *
FROM feedbacks
WHERE order_id = ? AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 2 OFFSET 3;

-- name: GetFeedbackByCustomerId :many
SELECT *
FROM feedbacks
WHERE customer_id = ? AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 2 OFFSET 3;

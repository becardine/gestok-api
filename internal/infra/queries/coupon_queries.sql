-- name: GetCoupon :one
SELECT * FROM coupons WHERE id = ? AND deleted_at IS NULL;

-- name: CreateCoupon :exec
INSERT INTO coupons (id, code, discount, expiration_at, status, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateCoupon :exec
UPDATE coupons
SET code = ?, discount = ?, expiration_at = ?, status = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteCoupon :exec
UPDATE coupons
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListCoupons :many
SELECT *
FROM coupons
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 1 OFFSET 2;

-- name: GetCouponByCode :one
SELECT *
FROM coupons
WHERE LOWER(code) = LOWER(?) AND deleted_at IS NULL;

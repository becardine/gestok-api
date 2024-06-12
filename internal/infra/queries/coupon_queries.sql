-- name: GetCoupon :one
SELECT * FROM coupons WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateCoupon :exec
INSERT INTO coupons (id, code, discount, expiration_date, status, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateCoupon :exec
UPDATE coupons
SET code = $2, discount = $3, expiration_date = $4, status = $5, updated_date = $6
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteCoupon :exec
UPDATE coupons
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListCoupons :many
SELECT * 
FROM coupons 
WHERE deleted_at IS NULL 
ORDER BY created_date DESC
LIMIT $1 OFFSET $2; 

-- name: GetCouponByCode :one
SELECT * 
FROM coupons 
WHERE LOWER(code) = LOWER($1) AND deleted_at IS NULL;
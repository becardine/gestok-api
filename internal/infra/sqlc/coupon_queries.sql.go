// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: coupon_queries.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

const createCoupon = `-- name: CreateCoupon :exec
INSERT INTO coupons (id, code, discount, expiration_date, status, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateCouponParams struct {
	ID             common.ID
	Code           string
	Discount       string
	ExpirationDate time.Time
	Status         string
	CreatedDate    sql.NullTime
	UpdatedDate    sql.NullTime
}

func (q *Queries) CreateCoupon(ctx context.Context, arg CreateCouponParams) error {
	_, err := q.db.ExecContext(ctx, createCoupon,
		arg.ID,
		arg.Code,
		arg.Discount,
		arg.ExpirationDate,
		arg.Status,
		arg.CreatedDate,
		arg.UpdatedDate,
	)
	return err
}

const deleteCoupon = `-- name: DeleteCoupon :exec
UPDATE coupons
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteCoupon(ctx context.Context, id common.ID) error {
	_, err := q.db.ExecContext(ctx, deleteCoupon, id)
	return err
}

const getCoupon = `-- name: GetCoupon :one
SELECT id, code, discount, expiration_date, status, deleted_at, created_date, updated_date FROM coupons WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetCoupon(ctx context.Context, id common.ID) (Coupon, error) {
	row := q.db.QueryRowContext(ctx, getCoupon, id)
	var i Coupon
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Discount,
		&i.ExpirationDate,
		&i.Status,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getCouponByCode = `-- name: GetCouponByCode :one
SELECT id, code, discount, expiration_date, status, deleted_at, created_date, updated_date
FROM coupons
WHERE LOWER(code) = LOWER($1) AND deleted_at IS NULL
`

func (q *Queries) GetCouponByCode(ctx context.Context, lower string) (Coupon, error) {
	row := q.db.QueryRowContext(ctx, getCouponByCode, lower)
	var i Coupon
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Discount,
		&i.ExpirationDate,
		&i.Status,
		&i.DeletedAt,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const listCoupons = `-- name: ListCoupons :many
SELECT id, code, discount, expiration_date, status, deleted_at, created_date, updated_date
FROM coupons
WHERE deleted_at IS NULL
ORDER BY created_date DESC
    LIMIT $1 OFFSET $2
`

type ListCouponsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListCoupons(ctx context.Context, arg ListCouponsParams) ([]Coupon, error) {
	rows, err := q.db.QueryContext(ctx, listCoupons, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Coupon
	for rows.Next() {
		var i Coupon
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Discount,
			&i.ExpirationDate,
			&i.Status,
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

const updateCoupon = `-- name: UpdateCoupon :exec
UPDATE coupons
SET code = $2, discount = $3, expiration_date = $4, status = $5, updated_date = $6
WHERE id = $1 AND deleted_at IS NULL
`

type UpdateCouponParams struct {
	ID             common.ID
	Code           string
	Discount       string
	ExpirationDate time.Time
	Status         string
	UpdatedDate    sql.NullTime
}

func (q *Queries) UpdateCoupon(ctx context.Context, arg UpdateCouponParams) error {
	_, err := q.db.ExecContext(ctx, updateCoupon,
		arg.ID,
		arg.Code,
		arg.Discount,
		arg.ExpirationDate,
		arg.Status,
		arg.UpdatedDate,
	)
	return err
}

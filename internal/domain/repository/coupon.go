package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CouponRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Coupon, error)
	Create(ctx context.Context, coupon *entity.Coupon) (*entity.Coupon, error)
	Update(ctx context.Context, coupon *entity.Coupon) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Coupon, error)
	GetCouponProducts(ctx context.Context, couponID uuid.UUID) ([]*entity.Product, error)
}

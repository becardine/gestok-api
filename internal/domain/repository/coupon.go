package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type CouponRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Coupon, error)
	Create(ctx context.Context, coupon *entity.Coupon) (*entity.Coupon, error)
	Update(ctx context.Context, coupon *entity.Coupon) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Coupon, error)
	GetCouponProducts(ctx context.Context, couponID common.ID) ([]*entity.Product, error)
	AddCouponProduct(ctx context.Context, productID, couponID common.ID) error
	RemoveCouponProduct(ctx context.Context, productID, couponID common.ID) error
}

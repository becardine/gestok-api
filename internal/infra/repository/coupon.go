package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
)

type CouponRepository struct {
	queries *database.Queries
}

func NewCouponRepository(db *sql.DB) domain.CouponRepositoryInterface {
	return &CouponRepository{
		queries: database.New(db),
	}
}

// Create implements repository.CouponRepositoryInterface.
func (c *CouponRepository) Create(ctx context.Context, coupon *entity.Coupon) (*entity.Coupon, error) {
	err := c.queries.CreateCoupon(ctx, database.CreateCouponParams{
		ID:             coupon.ID,
		Code:           coupon.Code,
		Discount:       fmt.Sprintf("%.2f", coupon.Discount),
		ExpirationDate: coupon.ExpirationDate,
		CreatedDate:    sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedDate:    sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return coupon, nil
}

// Delete implements repository.CouponRepositoryInterface.
func (c *CouponRepository) Delete(ctx context.Context, id common.ID) error {
	err := c.queries.DeleteCoupon(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Get implements repository.CouponRepositoryInterface.
func (c *CouponRepository) Get(ctx context.Context, id common.ID) (*entity.Coupon, error) {
	coupon, err := c.queries.GetCoupon(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Coupon{
		ID:             coupon.ID,
		Code:           coupon.Code,
		Discount:       coupon.Discount,
		ExpirationDate: coupon.ExpirationDate,
		Status:         coupon.Status,
	}, nil
}

// GetCouponProducts implements repository.CouponRepositoryInterface.
func (c *CouponRepository) GetCouponProducts(ctx context.Context, couponID common.ID) ([]*entity.Product, error) {
	panic("unimplemented")
}

// List implements repository.CouponRepositoryInterface.
func (c *CouponRepository) List(ctx context.Context, page int, pageSize int) ([]*entity.Coupon, error) {
	panic("unimplemented")
}

// Update implements repository.CouponRepositoryInterface.
func (c *CouponRepository) Update(ctx context.Context, coupon *entity.Coupon) error {
	panic("unimplemented")
}

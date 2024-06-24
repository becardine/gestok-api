package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
	"github.com/google/uuid"
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
		Discount:       coupon.Discount,
		ExpirationDate: coupon.ExpirationDate,
		CreatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return coupon, nil
}

// Delete implements repository.CouponRepositoryInterface.
func (c *CouponRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.queries.DeleteCoupon(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Get implements repository.CouponRepositoryInterface.
func (c *CouponRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Coupon, error) {
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
func (c *CouponRepository) GetCouponProducts(ctx context.Context, couponID uuid.UUID) ([]*entity.Product, error) {
	products, err := c.queries.GetCategoryProducts(ctx, couponID)
	if err != nil {
		return nil, err
	}

	result := make([]*entity.Product, len(products))
	for i, product := range products {
		result[i] = &entity.Product{
			ID:              product.ID,
			Name:            product.Name,
			Description:     product.Description.String,
			Price:           product.Price,
			QuantityInStock: int(product.QuantityInStock),
			ImageURL:        product.ImageUrl.String,
			CategoryID:      product.CategoryID,
			BrandID:         product.BrandID,
		}
	}

	return result, nil
}

// List implements repository.CouponRepositoryInterface.
func (c *CouponRepository) List(ctx context.Context, page int, pageSize int) ([]*entity.Coupon, error) {
	coupons, err := c.queries.ListCoupons(ctx, database.ListCouponsParams{
		Limit:  int32(pageSize),
		Offset: int32((page - 1) * pageSize),
	})
	if err != nil {
		return nil, err
	}

	result := make([]*entity.Coupon, len(coupons))
	for i, coupon := range coupons {
		result[i] = &entity.Coupon{
			ID:             coupon.ID,
			Code:           coupon.Code,
			Discount:       coupon.Discount,
			ExpirationDate: coupon.ExpirationDate,
			Status:         coupon.Status,
		}
	}

	return result, nil
}

// Update implements repository.CouponRepositoryInterface.
func (c *CouponRepository) Update(ctx context.Context, coupon *entity.Coupon) error {
	err := c.queries.UpdateCoupon(ctx, database.UpdateCouponParams{
		ID:             coupon.ID,
		Code:           coupon.Code,
		Discount:       coupon.Discount,
		ExpirationDate: coupon.ExpirationDate,
		Status:         coupon.Status,
		UpdatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}

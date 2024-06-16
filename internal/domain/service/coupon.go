package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
)

type CouponServiceInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Coupon, error)
	Create(ctx context.Context, input *dto.CreateCouponInput) (*entity.Coupon, error)
	Update(ctx context.Context, input *dto.UpdateCouponInput) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Coupon, error)
	GetCouponProducts(ctx context.Context, couponID common.ID) ([]*entity.Product, error)
}

type couponService struct {
	repo domain.CouponRepositoryInterface
}

func NewCouponService(repo domain.CouponRepositoryInterface) CouponServiceInterface {
	return &couponService{
		repo: repo,
	}
}

func (c *couponService) Get(ctx context.Context, id common.ID) (*entity.Coupon, error) {
	coupon, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, c.handleCouponError(err, id)
	}

	return coupon, nil
}

func (c *couponService) Create(ctx context.Context, input *dto.CreateCouponInput) (*entity.Coupon, error) {
	coupon := input.ToEntity()
	if err := coupon.Validate(); err != nil {
		return nil, fmt.Errorf("error while validating coupon: %w", err)
	}

	coupon, err := c.repo.Create(ctx, coupon)
	if err != nil {
		return nil, fmt.Errorf("error while creating coupon: %w", err)
	}

	return coupon, nil
}

func (c *couponService) Update(ctx context.Context, input *dto.UpdateCouponInput) error {
	coupon, err := c.repo.Get(ctx, input.ID)
	if err != nil {
		return c.handleCouponError(err, input.ID)
	}

	coupon.Code = input.Code
	coupon.Discount = input.Discount
	coupon.ExpirationDate = input.ExpirationDate
	coupon.Status = input.Status

	if err := coupon.Validate(); err != nil {
		return fmt.Errorf("error while validating coupon: %w", err)
	}

	if err := c.repo.Update(ctx, coupon); err != nil {
		return fmt.Errorf("error while updating coupon: %w", err)
	}

	return nil
}

func (c *couponService) Delete(ctx context.Context, id common.ID) error {
	_, err := c.Get(ctx, id)
	if err != nil {
		return err
	}

	if err := c.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error while deleting coupon in repository: %w", err)
	}

	return nil
}

func (c *couponService) List(ctx context.Context, page, pageSize int) ([]*entity.Coupon, error) {
	coupons, err := c.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("error while fetching coupons from repository: %w", err)
	}

	return coupons, nil
}

func (c *couponService) GetCouponProducts(ctx context.Context, couponID common.ID) ([]*entity.Product, error) {
	products, err := c.repo.GetCouponProducts(ctx, couponID)
	if err != nil {
		return nil, fmt.Errorf("error while fetching coupon products from repository: %w", err)
	}

	return products, nil
}

func (c *couponService) handleCouponError(err error, id common.ID) error {
	if errors.Is(err, &ErrNotFound{}) {
		return &ErrNotFound{Entity: "Coupon", ID: id}
	}

	return fmt.Errorf("error while fetching coupon from repository: %w", err)
}

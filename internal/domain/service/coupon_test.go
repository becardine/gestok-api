package service_test

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type CouponRepositoryMock struct {
	mock.Mock
}

func (m *CouponRepositoryMock) Get(ctx context.Context, id common.ID) (*entity.Coupon, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Coupon), args.Error(1)
}

func (m *CouponRepositoryMock) Create(ctx context.Context, coupon *entity.Coupon) (*entity.Coupon, error) {
	args := m.Called(ctx, coupon)
	return args.Get(0).(*entity.Coupon), args.Error(1)
}

func (m *CouponRepositoryMock) Update(ctx context.Context, coupon *entity.Coupon) error {
	args := m.Called(ctx, coupon)
	return args.Error(0)
}

func (m *CouponRepositoryMock) Delete(ctx context.Context, id common.ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *CouponRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Coupon, error) {
	args := m.Called(ctx, page, pageSize)
	return args.Get(0).([]*entity.Coupon), args.Error(1)
}

func (m *CouponRepositoryMock) GetCouponProducts(ctx context.Context, couponID common.ID) ([]*entity.Product, error) {
	args := m.Called(ctx, couponID)
	return args.Get(0).([]*entity.Product), args.Error(1)
}

func TestCouponService_Get(t *testing.T) {
	t.Run("should create a new coupon successfully", func(t *testing.T) {
		mockRepo := new(CouponRepositoryMock)
		couponService := service.NewCouponService(mockRepo)

		id := common.NewID()
		expectedCoupon := &entity.Coupon{
			ID:             id,
			Code:           "NEWCOUPON",
			Discount:       10,
			ExpirationDate: time.Now().AddDate(0, 0, 7),
			Status:         "active",
		}

		mockRepo.On("Get", mock.Anything, id).Return(expectedCoupon, nil)

		coupon, err := couponService.Get(context.Background(), id)

		assert.NoError(t, err)
		assert.Equal(t, expectedCoupon, coupon)
		mockRepo.AssertExpectations(t)
	})
}

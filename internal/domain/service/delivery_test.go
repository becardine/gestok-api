package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeliveryRepositoryMock struct {
	mock.Mock
}

func (m *DeliveryRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*entity.Delivery, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Delivery), args.Error(1)
}

func (m *DeliveryRepositoryMock) Create(ctx context.Context, delivery *entity.Delivery) (*entity.Delivery, error) {
	args := m.Called(ctx, delivery)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Delivery), args.Error(1)
}

func (m *DeliveryRepositoryMock) Update(ctx context.Context, delivery *entity.Delivery) error {
	args := m.Called(ctx, delivery)
	return args.Error(0)
}

func (m *DeliveryRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DeliveryRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Delivery, error) {
	args := m.Called(ctx, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Delivery), args.Error(1)
}

func TestDeliveryService(t *testing.T) {
	t.Run("should create a delivery successfully", func(t *testing.T) {
		mockRepo := new(DeliveryRepositoryMock)
		deliveryService := service.NewDeliveryService(mockRepo)

		input := &dto.CreateDeliveryInput{
			OrderID:        uuid.New(),
			CustomerID:     uuid.New(),
			DeliveryType:   "delivery_type",
			DeliveryDate:   time.Now(),
			DeliveryStatus: "delivery_status",
		}

		delivery := input.ToEntity()
		mockRepo.On("Create", mock.Anything, delivery).Return(delivery, nil)

		delivery, err := deliveryService.Create(context.Background(), input)

		assert.Nil(t, err)
		assert.NotNil(t, delivery)
		assert.Equal(t, input.OrderID, delivery.OrderID)
		assert.Equal(t, input.CustomerID, delivery.CustomerID)
		assert.Equal(t, input.DeliveryType, delivery.DeliveryType)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to create a delivery due to validation error", func(t *testing.T) {
		mockRepo := new(DeliveryRepositoryMock)
		deliveryService := service.NewDeliveryService(mockRepo)

		input := &dto.CreateDeliveryInput{
			OrderID:        uuid.Nil,
			CustomerID:     uuid.Nil,
			DeliveryType:   "",
			DeliveryDate:   time.Time{},
			DeliveryStatus: "",
		}

		delivery, err := deliveryService.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, delivery)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to create a delivery due to repository error", func(t *testing.T) {
		mockRepo := new(DeliveryRepositoryMock)
		deliveryService := service.NewDeliveryService(mockRepo)

		input := &dto.CreateDeliveryInput{
			OrderID:        uuid.New(),
			CustomerID:     uuid.New(),
			DeliveryType:   "delivery_type",
			DeliveryDate:   time.Now(),
			DeliveryStatus: "delivery_status",
		}

		delivery := input.ToEntity()
		mockRepo.On("Create", mock.Anything, delivery).Return(nil, assert.AnError)

		delivery, err := deliveryService.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, delivery)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should delete a delivery successfully", func(t *testing.T) {
		mockRepo := new(DeliveryRepositoryMock)
		deliveryService := service.NewDeliveryService(mockRepo)

		id := uuid.New()
		mockRepo.On("Get", mock.Anything, id).Return(&entity.Delivery{}, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(nil)

		err := deliveryService.Delete(context.Background(), id)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to delete a delivery due to repository error", func(t *testing.T) {
		mockRepo := new(DeliveryRepositoryMock)
		deliveryService := service.NewDeliveryService(mockRepo)

		id := uuid.New()
		mockRepo.On("Get", mock.Anything, id).Return(nil, assert.AnError)

		err := deliveryService.Delete(context.Background(), id)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

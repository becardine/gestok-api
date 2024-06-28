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

type OrderRepositoryMock struct {
	mock.Mock
}

func (m *OrderRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*entity.Order, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *OrderRepositoryMock) Create(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	args := m.Called(ctx, order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *OrderRepositoryMock) Update(ctx context.Context, order *entity.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *OrderRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *OrderRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Order, error) {
	args := m.Called(ctx, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func TestOrderService(t *testing.T) {
	t.Run("should create order successfully", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		input := &dto.CreateOrderInput{
			CustomerID:  uuid.New(),
			OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderStatus: "pending",
			TotalValue:  100.0,
		}

		order := input.ToEntity()
		mockRepo.On("Create", mock.Anything, mock.Anything).Return(order, nil)
		order, err := orderService.Create(context.Background(), input)

		assert.Nil(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, input.CustomerID, order.CustomerID)
		assert.Equal(t, input.OrderDate, order.OrderDate)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should get order successfully", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		order := &entity.Order{
			ID:          uuid.New(),
			CustomerID:  uuid.New(),
			OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderStatus: "pending",
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(order, nil)
		order, err := orderService.Get(context.Background(), order.ID)

		assert.Nil(t, err)
		assert.NotNil(t, order)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should delete order successfully", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		order := &entity.Order{
			ID:          uuid.New(),
			CustomerID:  uuid.New(),
			OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderStatus: "pending",
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(order, nil)
		mockRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)
		err := orderService.Delete(context.Background(), order.ID)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should list orders successfully", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		orders := []*entity.Order{
			{
				ID:          uuid.New(),
				CustomerID:  uuid.New(),
				OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				OrderStatus: "pending",
			},
			{
				ID:          uuid.New(),
				CustomerID:  uuid.New(),
				OrderDate:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
				OrderStatus: "completed",
			},
		}

		mockRepo.On("List", mock.Anything, mock.Anything, mock.Anything).Return(orders, nil)
		orders, err := orderService.List(context.Background(), 1, 10)

		assert.Nil(t, err)
		assert.NotNil(t, orders)
		assert.Equal(t, 2, len(orders))
		mockRepo.AssertExpectations(t)
	})

	t.Run("should update order successfully", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		id := uuid.New()
		input := &dto.UpdateOrderInput{
			ID:          id,
			CustomerID:  uuid.New(),
			OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderStatus: "pending",
		}

		order := input.ToEntity()

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(order, nil)
		mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
		_, err := orderService.Update(context.Background(), id, input)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to create order", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		input := &dto.CreateOrderInput{
			CustomerID:  uuid.New(),
			OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderStatus: "pending",
		}

		mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		order, err := orderService.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, order)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to get order", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		order, err := orderService.Get(context.Background(), uuid.New())

		assert.NotNil(t, err)
		assert.Nil(t, order)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to delete order", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		order := &entity.Order{
			ID:          uuid.New(),
			CustomerID:  uuid.New(),
			OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderStatus: "pending",
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(order, nil)
		mockRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)
		err := orderService.Delete(context.Background(), order.ID)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to list orders", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		mockRepo.On("List", mock.Anything, mock.Anything, mock.Anything).Return(nil, assert.AnError)
		orders, err := orderService.List(context.Background(), 1, 10)

		assert.NotNil(t, err)
		assert.Nil(t, orders)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to update order", func(t *testing.T) {
		mockRepo := new(OrderRepositoryMock)
		orderService := service.NewOrderService(mockRepo)

		id := uuid.New()
		input := &dto.UpdateOrderInput{
			ID:          id,
			CustomerID:  uuid.New(),
			OrderDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderStatus: "pending",
		}

		order := input.ToEntity()

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(order, nil)
		mockRepo.On("Update", mock.Anything, mock.Anything).Return(assert.AnError)
		_, err := orderService.Update(context.Background(), id, input)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

package service_test

import (
	"context"
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type OrderItemRepositoryMock struct {
	mock.Mock
}

func (m *OrderItemRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*entity.OrderItem, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.OrderItem), args.Error(1)
}

func (m *OrderItemRepositoryMock) Create(ctx context.Context, order *entity.OrderItem) (*entity.OrderItem, error) {
	args := m.Called(order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.OrderItem), args.Error(1)
}

func (m *OrderItemRepositoryMock) Update(ctx context.Context, order *entity.OrderItem) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *OrderItemRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *OrderItemRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.OrderItem, error) {
	args := m.Called(page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*entity.OrderItem), args.Error(1)
}

func (m *OrderItemRepositoryMock) GetOrderItemProduct(ctx context.Context, orderItemID uuid.UUID) (*entity.Product, error) {
	args := m.Called(orderItemID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *OrderItemRepositoryMock) SetOrderItemProduct(ctx context.Context, productID, orderItemID uuid.UUID) error {
	args := m.Called(productID, orderItemID)
	return args.Error(0)
}

func (m *OrderItemRepositoryMock) GetOrderItemOrder(ctx context.Context, orderItemID uuid.UUID) (*entity.Order, error) {
	args := m.Called(orderItemID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *OrderItemRepositoryMock) SetOrderItemOrder(ctx context.Context, orderID, orderItemID uuid.UUID) error {
	args := m.Called(orderID, orderItemID)
	return args.Error(0)
}

func TestOrderItemService(t *testing.T) {
	t.Run("should create order item", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		input := &dto.CreateOrderItemInput{
			ProductID: uuid.New(),
			OrderID:   uuid.New(),
			Quantity:  1,
			UnitPrice: 100,
		}

		orderItem := input.ToEntity()
		mockRepo.On("Create", mock.Anything).Return(orderItem, nil)

		orderItem, err := orderItemService.Create(context.Background(), input)

		assert.Nil(t, err)
		assert.NotNil(t, orderItem)
		assert.Equal(t, input.ProductID, orderItem.ProductID)
		assert.Equal(t, input.OrderID, orderItem.OrderID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not create order item with invalid input", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		input := &dto.CreateOrderItemInput{
			ProductID: uuid.Nil,
			OrderID:   uuid.Nil,
			Quantity:  0,
			UnitPrice: 0,
		}

		orderItem, err := orderItemService.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, orderItem)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should get order item", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItem := &entity.OrderItem{
			ID:        uuid.New(),
			OrderID:   uuid.New(),
			ProductID: uuid.New(),
			Quantity:  1,
			UnitPrice: 100,
		}

		mockRepo.On("Get", orderItem.ID).Return(orderItem, nil)

		orderItem, err := orderItemService.Get(context.Background(), orderItem.ID)

		assert.Nil(t, err)
		assert.NotNil(t, orderItem)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not get order item with invalid id", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItem, err := orderItemService.Get(context.Background(), uuid.Nil)

		assert.NotNil(t, err)
		assert.Nil(t, orderItem)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should delete order item", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItem := &entity.OrderItem{
			ID:        uuid.New(),
			OrderID:   uuid.New(),
			ProductID: uuid.New(),
			Quantity:  1,
			UnitPrice: 100,
		}

		mockRepo.On("Get", orderItem.ID).Return(orderItem, nil)
		mockRepo.On("Delete", orderItem.ID).Return(nil)

		err := orderItemService.Delete(context.Background(), orderItem.ID)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not delete order item with invalid id", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		err := orderItemService.Delete(context.Background(), uuid.Nil)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should list order items", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItems := []*entity.OrderItem{
			{
				ID:        uuid.New(),
				OrderID:   uuid.New(),
				ProductID: uuid.New(),
				Quantity:  1,
				UnitPrice: 100,
			},
			{
				ID:        uuid.New(),
				OrderID:   uuid.New(),
				ProductID: uuid.New(),
				Quantity:  1,
				UnitPrice: 100,
			},
		}

		mockRepo.On("List", 1, 10).Return(orderItems, nil)

		orderItems, err := orderItemService.List(context.Background(), 1, 10)

		assert.Nil(t, err)
		assert.NotNil(t, orderItems)
		assert.Len(t, orderItems, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not list order items with invalid page", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItems, err := orderItemService.List(context.Background(), 0, 10)

		assert.NotNil(t, err)
		assert.Nil(t, orderItems)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not list order items with invalid page size", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItems, err := orderItemService.List(context.Background(), 1, 0)

		assert.NotNil(t, err)
		assert.Nil(t, orderItems)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should get order item product", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItem := &entity.OrderItem{
			ID:        uuid.New(),
			OrderID:   uuid.New(),
			ProductID: uuid.New(),
			Quantity:  1,
			UnitPrice: 100,
		}

		product := &entity.Product{
			ID:   uuid.New(),
			Name: "Product",
		}

		mockRepo.On("GetOrderItemProduct", orderItem.ID).Return(product, nil)

		product, err := orderItemService.GetOrderItemProduct(context.Background(), orderItem.ID)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not get order item product with invalid id", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		product, err := orderItemService.GetOrderItemProduct(context.Background(), uuid.Nil)

		assert.NotNil(t, err)
		assert.Nil(t, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should set order item product", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		productID := uuid.New()
		orderItemID := uuid.New()

		mockRepo.On("SetOrderItemProduct", productID, orderItemID).Return(nil)

		err := orderItemService.SetOrderItemProduct(context.Background(), productID, orderItemID)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not set order item product with invalid id", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		err := orderItemService.SetOrderItemProduct(context.Background(), uuid.Nil, uuid.Nil)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should set order item order", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderID := uuid.New()
		orderItemID := uuid.New()

		mockRepo.On("SetOrderItemOrder", orderID, orderItemID).Return(nil)

		err := orderItemService.SetOrderItemOrder(context.Background(), orderID, orderItemID)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not set order item order with invalid id", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		err := orderItemService.SetOrderItemOrder(context.Background(), uuid.Nil, uuid.Nil)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should update order item", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItem := &entity.OrderItem{
			ID:        uuid.New(),
			OrderID:   uuid.New(),
			ProductID: uuid.New(),
			Quantity:  1,
			UnitPrice: 100,
		}

		input := &dto.UpdateOrderItemInput{
			Quantity:  2,
			UnitPrice: 200,
		}

		mockRepo.On("Get", orderItem.ID).Return(orderItem, nil)
		mockRepo.On("Update", orderItem).Return(nil)

		orderItem, err := orderItemService.Update(context.Background(), orderItem.ID, input)

		assert.Nil(t, err)
		assert.NotNil(t, orderItem)
		assert.Equal(t, input.Quantity, orderItem.Quantity)
		assert.Equal(t, input.UnitPrice, orderItem.UnitPrice)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not update order item with invalid id", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		input := &dto.UpdateOrderItemInput{
			Quantity:  2,
			UnitPrice: 200,
		}

		orderItem, err := orderItemService.Update(context.Background(), uuid.Nil, input)

		assert.NotNil(t, err)
		assert.Nil(t, orderItem)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not update order item with invalid input", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		orderItem := &entity.OrderItem{
			ID:        uuid.New(),
			OrderID:   uuid.New(),
			ProductID: uuid.New(),
			Quantity:  1,
			UnitPrice: 100,
		}

		input := &dto.UpdateOrderItemInput{
			Quantity:  0,
			UnitPrice: 0,
		}

		mockRepo.On("Get", orderItem.ID).Return(orderItem, nil)

		orderItem, err := orderItemService.Update(context.Background(), orderItem.ID, input)

		assert.NotNil(t, err)
		assert.Nil(t, orderItem)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not update order item with invalid entity", func(t *testing.T) {
		mockRepo := new(OrderItemRepositoryMock)
		orderItemService := service.NewOrderItemService(mockRepo)

		input := &dto.UpdateOrderItemInput{
			Quantity:  2,
			UnitPrice: 200,
		}

		orderItem, err := orderItemService.Update(context.Background(), uuid.New(), input)

		assert.NotNil(t, err)
		assert.Nil(t, orderItem)
		mockRepo.AssertExpectations(t)
	})
}

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

type StockRepositoryMock struct {
	mock.Mock
}

func (m *StockRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*entity.Stock, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Stock), args.Error(1)
}

func (m *StockRepositoryMock) Create(ctx context.Context, stock *entity.Stock) (*entity.Stock, error) {
	args := m.Called(ctx, stock)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Stock), args.Error(1)
}

func (m *StockRepositoryMock) Update(ctx context.Context, stock *entity.Stock) error {
	args := m.Called(ctx, stock)
	return args.Error(0)
}

func (m *StockRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *StockRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Stock, error) {
	args := m.Called(ctx, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Stock), args.Error(1)
}

func TestStockService(t *testing.T) {
	t.Run("should create stock successfully", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		input := &dto.CreateStockInput{
			Name:     "Stock Test",
			Location: "Location Test",
			Capacity: 100,
		}

		stock := input.ToEntity()

		mockRepo.On("Create", mock.Anything, mock.Anything).Return(stock, nil)
		stock, err := stockService.Create(context.Background(), input)
		assert.Nil(t, err)
		assert.NotNil(t, stock)
		assert.Equal(t, input.Name, stock.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should get stock successfully", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		stock := &entity.Stock{
			ID:       uuid.New(),
			Name:     "Stock Test",
			Location: "Location Test",
			Capacity: 100,
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(stock, nil)
		stock, err := stockService.Get(context.Background(), stock.ID)
		assert.Nil(t, err)
		assert.NotNil(t, stock)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should list stocks successfully", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		stocks := []*entity.Stock{
			{
				ID:       uuid.New(),
				Name:     "Stock Test",
				Location: "Location Test",
				Capacity: 100,
			},
		}

		mockRepo.On("List", mock.Anything, mock.Anything, mock.Anything).Return(stocks, nil)
		stocks, err := stockService.List(context.Background(), 1, 10)
		assert.Nil(t, err)
		assert.NotNil(t, stocks)
		assert.Equal(t, 1, len(stocks))
		mockRepo.AssertExpectations(t)
	})

	t.Run("should update stock successfully", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		stock := &entity.Stock{
			ID:       uuid.New(),
			Name:     "Stock Test",
			Location: "Location Test",
			Capacity: 100,
		}

		input := &dto.UpdateStockInput{
			ID:       stock.ID,
			Name:     "Stock Test Updated",
			Location: "Location Test Updated",
			Capacity: 200,
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(stock, nil)
		mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)
		err := stockService.Update(context.Background(), input)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should delete stock successfully", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		stock := &entity.Stock{
			ID:       uuid.New(),
			Name:     "Stock Test",
			Location: "Location Test",
			Capacity: 100,
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(stock, nil)
		mockRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)
		err := stockService.Delete(context.Background(), stock.ID)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when stock not found", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		stock := &entity.Stock{
			ID: uuid.New(),
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		err := stockService.Delete(context.Background(), stock.ID)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when stock already exists", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		input := &dto.CreateStockInput{
			Name:     "Stock Test",
			Location: "Location Test",
			Capacity: 100,
		}

		mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		_, err := stockService.Create(context.Background(), input)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when stock validation fails", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		input := &dto.CreateStockInput{
			Name:     "Stock Test",
			Location: "Location Test",
			Capacity: -1,
		}

		stock, err := stockService.Create(context.Background(), input)
		assert.Nil(t, stock)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when stock not found", func(t *testing.T) {
		mockRepo := new(StockRepositoryMock)
		stockService := service.NewStockService(mockRepo)

		stock := &entity.Stock{
			ID: uuid.New(),
		}

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		stock, err := stockService.Get(context.Background(), stock.ID)
		assert.Nil(t, stock)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

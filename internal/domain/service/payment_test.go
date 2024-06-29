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

type PaymentRepositoryMock struct {
	mock.Mock
}

func (m *PaymentRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*entity.Payment, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *PaymentRepositoryMock) Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	args := m.Called(ctx, payment)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *PaymentRepositoryMock) Update(ctx context.Context, payment *entity.Payment) error {
	args := m.Called(ctx, payment)
	return args.Error(0)
}

func (m *PaymentRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *PaymentRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Payment, error) {
	args := m.Called(ctx, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Payment), args.Error(1)
}

func TestPaymentService(t *testing.T) {
	t.Run("should create payment successfully", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		input := &dto.CreatePaymentInput{
			OrderID:    uuid.New(),
			CustomerID: uuid.New(),
			Method:     "cash",
			Date:       time.Now(),
			Amount:     100,
			Status:     "paid",
		}

		payment := input.ToEntity()

		mockRepo.On("Create", mock.Anything, mock.Anything).Return(payment, nil)
		payment, err := paymentService.Create(context.Background(), input)
		assert.Nil(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, input.OrderID, payment.OrderID)
		assert.Equal(t, input.CustomerID, payment.CustomerID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should delete payment successfully", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		id := uuid.New()
		mockRepo.On("Get", mock.Anything, id).Return(&entity.Payment{}, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(nil)

		err := paymentService.Delete(context.Background(), id)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should get payment successfully", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		id := uuid.New()
		mockRepo.On("Get", mock.Anything, id).Return(&entity.Payment{}, nil)

		payment, err := paymentService.Get(context.Background(), id)
		assert.Nil(t, err)
		assert.NotNil(t, payment)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should list payments successfully", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		mockRepo.On("List", mock.Anything, 1, 10).Return([]*entity.Payment{}, nil)

		payments, err := paymentService.List(context.Background(), 1, 10)
		assert.Nil(t, err)
		assert.NotNil(t, payments)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should update payment successfully", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		id := uuid.New()
		input := &dto.UpdatePaymentInput{
			ID:         id,
			OrderID:    uuid.New(),
			CustomerID: uuid.New(),
			Method:     "credit card",
			Date:       time.Now(),
			Amount:     200,
			Status:     "paid",
		}

		mockRepo.On("Get", mock.Anything, id).Return(&entity.Payment{}, nil)
		mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := paymentService.Update(context.Background(), input)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to create payment", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		input := &dto.CreatePaymentInput{
			OrderID:    uuid.New(),
			CustomerID: uuid.New(),
			Method:     "cash",
			Date:       time.Now(),
			Amount:     100,
			Status:     "paid",
		}

		mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		payment, err := paymentService.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Nil(t, payment)
	})

	t.Run("should fail to delete payment", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		err := paymentService.Delete(context.Background(), uuid.New())
		assert.NotNil(t, err)
	})

	t.Run("should fail to get payment", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		mockRepo.On("Get", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		payment, err := paymentService.Get(context.Background(), uuid.New())
		assert.NotNil(t, err)
		assert.Nil(t, payment)
	})

	t.Run("should fail to list payments", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		mockRepo.On("List", mock.Anything, mock.Anything, mock.Anything).Return(nil, assert.AnError)
		payments, err := paymentService.List(context.Background(), 1, 10)
		assert.NotNil(t, err)
		assert.Nil(t, payments)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to update payment", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		input := &dto.UpdatePaymentInput{
			Method: "credit card",
			Date:   time.Now(),
			Amount: 200,
			Status: "paid",
		}

		mockRepo.On("Get", mock.Anything, uuid.Nil).Return(nil, assert.AnError)
		err := paymentService.Update(context.Background(), input)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should fail to update payment", func(t *testing.T) {
		mockRepo := new(PaymentRepositoryMock)
		paymentService := service.NewPaymentService(mockRepo)

		id := uuid.New()
		input := &dto.UpdatePaymentInput{
			Method: "credit card",
			Date:   time.Now(),
			Amount: 200,
			Status: "paid",
		}

		mockRepo.On("Get", mock.Anything, id).Return(&entity.Payment{}, nil)
		mockRepo.On("Update", mock.Anything, mock.Anything).Return(assert.AnError)

		err := paymentService.Update(context.Background(), input)
		assert.NotNil(t, err)
	})
}

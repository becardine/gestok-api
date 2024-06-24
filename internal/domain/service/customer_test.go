package service_test

import (
	"context"
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*entity.Customer, error) {
	ret := m.Called(ctx, id)

	var r0 *entity.Customer
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Customer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *CustomerRepositoryMock) Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {
	ret := m.Called(ctx, customer)

	var r0 *entity.Customer
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Customer) *entity.Customer); ok {
		r0 = rf(ctx, customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Customer) error); ok {
		r1 = rf(ctx, customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *CustomerRepositoryMock) Update(ctx context.Context, customer *entity.Customer) error {
	ret := m.Called(ctx, customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Customer) error); ok {
		r0 = rf(ctx, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *CustomerRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	ret := m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *CustomerRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Customer, error) {
	ret := m.Called(ctx, page, pageSize)

	var r0 []*entity.Customer
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*entity.Customer); ok {
		r0 = rf(ctx, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *CustomerRepositoryMock) GetCustomerOrders(ctx context.Context, customerID uuid.UUID) ([]*entity.Order, error) {
	ret := m.Called(ctx, customerID)

	var r0 []*entity.Order
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*entity.Order); ok {
		r0 = rf(ctx, customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func TestCustomer(t *testing.T) {
	t.Run("should create a new customer successfully", func(t *testing.T) {
		mockRepo := new(CustomerRepositoryMock)
		customerService := service.NewCustomerService(mockRepo)

		ctx := context.Background()
		customer := &entity.Customer{
			ID:    uuid.New(),
			Name:  "John Doe",
			Email: "john@doe.com",
			Phone: "1234567890",
		}

		mockRepo.On("Get", mock.Anything, customer.ID).Return(customer, nil)

		customer, err := customerService.Get(ctx, customer.ID)

		assert.Nil(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, "John Doe", customer.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return an error when customer does not exist", func(t *testing.T) {
		mockRepo := new(CustomerRepositoryMock)
		customerService := service.NewCustomerService(mockRepo)

		ctx := context.Background()
		id := uuid.New()
		mockRepo.On("Get", mock.Anything, id).Return(nil, &service.ErrNotFound{Entity: "Customer", ID: id})

		customer, err := customerService.Get(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, customer)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return an error when an unexpected error occurs", func(t *testing.T) {
		mockRepo := new(CustomerRepositoryMock)
		customerService := service.NewCustomerService(mockRepo)

		ctx := context.Background()
		id := uuid.New()
		mockRepo.On("Get", mock.Anything, id).Return(nil, assert.AnError)

		customer, err := customerService.Get(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, customer)
		mockRepo.AssertExpectations(t)
	})
}

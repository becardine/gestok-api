package mocks

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m *ProductServiceMock) CreateProduct(ctx context.Context, input *service.CreateProductInput) (*entity.Product, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *ProductServiceMock) GetProduct(ctx context.Context, id common.ID) (*entity.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *ProductServiceMock) UpdateProduct(ctx context.Context, id common.ID, input *service.UpdateProductInput) error {
	args := m.Called(ctx, id, input)
	return args.Error(0)
}

func (m *ProductServiceMock) DeleteProduct(ctx context.Context, id common.ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *ProductServiceMock) ListProducts(ctx context.Context) ([]*entity.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Product), args.Error(1)
}

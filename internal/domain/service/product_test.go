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

type ProductRepositoryMock struct {
	mock.Mock
}

// DeleteProduct implements repository.ProductRepository.
func (m *ProductRepositoryMock) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// ListProducts implements repository.ProductRepository.
func (m *ProductRepositoryMock) ListProducts(ctx context.Context) ([]*entity.Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements repository.ProductRepository.
func (m *ProductRepositoryMock) UpdateProduct(ctx context.Context, product *entity.Product) error {
	panic("unimplemented")
}

func (m *ProductRepositoryMock) CreateProduct(ctx context.Context, product *entity.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func TestCreateProduct(t *testing.T) {
	repoMock := new(ProductRepositoryMock)
	productService := service.NewProductService(repoMock)

	input := &service.CreateProductInput{
		Name:            "Product Test",
		Description:     "Description Test",
		Price:           10.50,
		QuantityInStock: 10,
		ImageURL:        "https://www.example.com/image.jpg",
		CategoryID:      uuid.New(),
		BrandID:         uuid.New(),
	}

	repoMock.On("CreateProduct", mock.Anything, mock.AnythingOfType("*entity.Product")).Return(nil)

	product, err := productService.CreateProduct(context.Background(), input)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, input.Name, product.Name)
	repoMock.AssertExpectations(t)
}

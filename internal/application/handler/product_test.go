package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/becardine/gestock-api/internal/application/handler"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para o ProductService
type ProductServiceMock struct {
	mock.Mock
}

// DeleteProduct implements service.ProductService.
func (m *ProductServiceMock) DeleteProduct(ctx context.Context, id common.ID) error {
	panic("unimplemented")
}

// GetProduct implements service.ProductService.
func (m *ProductServiceMock) GetProduct(ctx context.Context, id common.ID) (*entity.Product, error) {
	panic("unimplemented")
}

// ListProducts implements service.ProductService.
func (m *ProductServiceMock) ListProducts(ctx context.Context) ([]*entity.Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements service.ProductService.
func (m *ProductServiceMock) UpdateProduct(ctx context.Context, id common.ID, input *service.UpdateProductInput) error {
	panic("unimplemented")
}

// CreateProduct implementa service.ProductService.
func (m *ProductServiceMock) CreateProduct(ctx context.Context, input *service.CreateProductInput) (*entity.Product, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func TestProductHandler_createProduct(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		mockService := new(ProductServiceMock)

		expectedInput := &service.CreateProductInput{
			Name:            "Product Test",
			Description:     "Description Test",
			Price:           10.50,
			QuantityInStock: 10,
			ImageURL:        "https://example.com/image.jpg",
			CategoryID:      common.NewID(),
			BrandID:         common.NewID(),
		}

		expectedProduct := &entity.Product{
			ID:              common.NewID(),
			Name:            "Product Test",
			Description:     "Description Test",
			Price:           10.50,
			QuantityInStock: 10,
			ImageURL:        "https://example.com/image.jpg",
			CategoryID:      common.NewID(),
			BrandID:         common.NewID(),
		}

		mockService.On("CreateProduct", mock.Anything, expectedInput).Return(expectedProduct, nil)

		productHandler := handler.NewProductHandler(mockService)

		productJSON := `{
			"name": "Product Test",
			"description": "Description Test",
			"price": 10.50,
			"quantity_in_stock": 10,
			"image_url": "https://example.com/image.jpg",
			"category_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
			"brand_id": "e22d4e85-f8da-4449-8441-096e99769769"
		}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBufferString(productJSON))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := chi.NewRouter()
		productHandler.Routes(router)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockService.AssertExpectations(t)

		var createdProduct entity.Product
		err = json.Unmarshal(rr.Body.Bytes(), &createdProduct)
		assert.NoError(t, err)

		assert.Equal(t, "Product Test", createdProduct.Name)
		assert.Equal(t, "Description Test", createdProduct.Description)
	})

}

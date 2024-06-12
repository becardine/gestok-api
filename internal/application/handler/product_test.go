package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/wire"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProductHandler_createProduct(t *testing.T) {
	t.Run("success when creating product", func(t *testing.T) {
		// Arrange
		err := config.Init()
		require.NoError(t, err)

		router := chi.NewRouter()
		mockService := mocks.NewProductServiceMock()
		productHandler, err := wire.InitializeProductHandler()
		require.NoError(t, err)

		router.Route("/products", productHandler.Routes)

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
			Name:            expectedInput.Name,
			Description:     expectedInput.Description,
			Price:           expectedInput.Price,
			QuantityInStock: expectedInput.QuantityInStock,
			ImageURL:        expectedInput.ImageURL,
			CategoryID:      expectedInput.CategoryID,
			BrandID:         expectedInput.BrandID,
		}

		mockService.On("CreateProduct", mock.Anything, expectedInput).Return(expectedProduct, nil)

		productJSON, err := json.Marshal(expectedInput)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productJSON))

		require.NoError(t, err)

		rr := httptest.NewRecorder()

		// Act
		router.ServeHTTP(rr, req)

		// Assert
		assert.Equal(t, http.StatusCreated, rr.Code)

		var createdProduct entity.Product
		err = json.Unmarshal(rr.Body.Bytes(), &createdProduct)
		assert.NoError(t, err)

		assert.Equal(t, expectedProduct.Name, createdProduct.Name)
		assert.Equal(t, expectedProduct.Description, createdProduct.Description)
		assert.Equal(t, expectedProduct.Price, createdProduct.Price)
		assert.Equal(t, expectedProduct.CategoryID.String(), createdProduct.CategoryID.String())
		assert.Equal(t, expectedProduct.BrandID.String(), createdProduct.BrandID.String())

		mockService.AssertExpectations(t)
	})
}

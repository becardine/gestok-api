package entity_test

import (
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const productName = "Product Test"
const productDescription = "Description Test"
const productUrl = "https://www.example.com/image.jpg"

func TestNewProduct(t *testing.T) {
	categoryID := uuid.New()
	brandID := uuid.New()

	product, err := entity.NewProduct(productName, productDescription, 10.50, 10, productUrl, categoryID, brandID)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, productName, product.Name)
	assert.Equal(t, productDescription, product.Description)
	assert.Equal(t, 10.50, product.Price)
	assert.Equal(t, 10, product.QuantityInStock)
	assert.Equal(t, productUrl, product.ImageURL)
	assert.Equal(t, categoryID, product.CategoryID)
	assert.Equal(t, brandID, product.BrandID)
}

func TestProductValidate(t *testing.T) {
	categoryID := uuid.New()
	brandID := uuid.New()

	t.Run("Successful validation", func(t *testing.T) {
		product := &entity.Product{
			ID:              uuid.New(),
			Name:            productName,
			Description:     productDescription,
			Price:           10.50,
			QuantityInStock: 10,
			ImageURL:        productUrl,
			CategoryID:      categoryID,
			BrandID:         brandID,
		}
		err := product.Validate()
		assert.Nil(t, err)
	})

}

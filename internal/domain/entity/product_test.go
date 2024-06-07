package entity_test

import (
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	categoryID := common.NewID()
	brandID := common.NewID()

	product, err := entity.NewProduct("Product Test", "Description Test", 10.50, 10, "https://www.example.com/image.jpg", categoryID, brandID)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product Test", product.Name)
	assert.Equal(t, "Description Test", product.Description)
	assert.Equal(t, 10.50, product.Price)
	assert.Equal(t, 10, product.QuantityInStock)
	assert.Equal(t, "https://www.example.com/image.jpg", product.ImageURL)
	assert.Equal(t, categoryID, product.CategoryID)
	assert.Equal(t, brandID, product.BrandID)
}

func TestProduct_Validate(t *testing.T) {
	categoryID := common.NewID()
	brandID := common.NewID()

	t.Run("Successful validation", func(t *testing.T) {
		product := &entity.Product{
			ID:              common.NewID(),
			Name:            "Product Test",
			Description:     "Description Test",
			Price:           10.50,
			QuantityInStock: 10,
			ImageURL:        "https://www.example.com/image.jpg",
			CategoryID:      categoryID,
			BrandID:         brandID,
		}
		err := product.Validate()
		assert.Nil(t, err)
	})

}

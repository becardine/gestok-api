package entity_test

import (
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewBrand(t *testing.T) {
	t.Run("should create a new brand with valid data", func(t *testing.T) {
		brand, err := entity.NewBrand("Brand Test", "Description Test")
		assert.NoError(t, err)
		assert.NotNil(t, brand)
		assert.Equal(t, "Brand Test", brand.Name)
		assert.Equal(t, "Description Test", brand.Description)
	})

	t.Run("should return error if name is empty", func(t *testing.T) {
		brand, err := entity.NewBrand("", "Description Test")
		assert.Error(t, err)
		assert.Nil(t, brand)
	})

	t.Run("should return error if name exceeds max length", func(t *testing.T) {
		longName := "This is an extremely long brand name that certainly exceeds the maximum allowed length of 100 characters by a significant margin"
		brand, err := entity.NewBrand(longName, "Description Test")
		assert.Error(t, err)
		assert.Nil(t, brand)
	})

	t.Run("should return error if description exceeds max length", func(t *testing.T) {
		longDescription := "This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters."
		brand, err := entity.NewBrand("Brand Test", longDescription)
		assert.Error(t, err)
		assert.Nil(t, brand)
	})
}

func TestBrand_Validate(t *testing.T) {
	t.Run("should return no error if brand is valid", func(t *testing.T) {
		brand := &entity.Brand{
			ID:          uuid.New(),
			Name:        "Brand Test",
			Description: "Description Test",
		}
		err := brand.Validate()
		assert.NoError(t, err)
	})

	t.Run("should return error if name is empty", func(t *testing.T) {
		brand := &entity.Brand{
			ID:          uuid.New(),
			Name:        "",
			Description: "Description Test",
		}
		err := brand.Validate()
		assert.Error(t, err)
	})

	t.Run("should return error if name exceeds max length", func(t *testing.T) {
		longName := "This is a very long brand name that exceeds the maximum allowed length of 100 characters"
		brand := &entity.Brand{
			ID:          uuid.New(),
			Name:        longName,
			Description: "Description Test",
		}
		err := brand.Validate()
		assert.Error(t, err)
	})

	t.Run("should return error if description exceeds max length", func(t *testing.T) {
		longDescription := "This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters."
		brand := &entity.Brand{
			ID:          uuid.New(),
			Name:        "Brand Test",
			Description: longDescription,
		}
		err := brand.Validate()
		assert.Error(t, err)
	})
}

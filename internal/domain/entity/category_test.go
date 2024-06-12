package entity_test

import (
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	t.Run("should create a new category with valid data", func(t *testing.T) {
		category, err := entity.NewCategory("Category Test", "Description Test")
		assert.NoError(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, "Category Test", category.Name)
		assert.Equal(t, "Description Test", category.Description)
	})

	t.Run("should return error if name is empty", func(t *testing.T) {
		category, err := entity.NewCategory("", "Description Test")
		assert.Error(t, err)
		assert.Nil(t, category)
	})

	t.Run("should return error if name exceeds max length", func(t *testing.T) {
		longName := "This is an extremely long category name that certainly exceeds the maximum allowed length of 100 characters by a significant margin"
		category, err := entity.NewCategory(longName, "Description Test")
		assert.Error(t, err)
		assert.Nil(t, category)
	})

	t.Run("should return error if description exceeds max length", func(t *testing.T) {
		longDescription := "This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters."
		category, err := entity.NewCategory("Category Test", longDescription)
		assert.Error(t, err)
		assert.Nil(t, category)
	})
}

func TestCategory_Validate(t *testing.T) {
	t.Run("should return no error if category is valid", func(t *testing.T) {
		category := &entity.Category{
			ID:          common.NewID(),
			Name:        "Category Test",
			Description: "Description Test",
		}
		err := category.Validate()
		assert.NoError(t, err)
	})

	t.Run("should return error if name is empty", func(t *testing.T) {
		category := &entity.Category{
			ID:          common.NewID(),
			Name:        "",
			Description: "Description Test",
		}
		err := category.Validate()
		assert.Error(t, err)
	})

	t.Run("should return error if name exceeds max length", func(t *testing.T) {
		longName := "This is a very long category name that exceeds the maximum allowed length of 100 characters"
		category := &entity.Category{
			ID:          common.NewID(),
			Name:        longName,
			Description: "Description Test",
		}
		err := category.Validate()
		assert.Error(t, err)
	})

	t.Run("should return error if description exceeds max length", func(t *testing.T) {
		longDescription := "This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters. This is a very long description that exceeds the maximum allowed length of 255 characters."
		category := &entity.Category{
			ID:          common.NewID(),
			Name:        "Category Test",
			Description: longDescription,
		}
		err := category.Validate()
		assert.Error(t, err)
	})
}

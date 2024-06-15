package service_test

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (m *CategoryRepositoryMock) Get(ctx context.Context, id common.ID) (*entity.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *CategoryRepositoryMock) Create(ctx context.Context, category *entity.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *CategoryRepositoryMock) Update(ctx context.Context, category *entity.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *CategoryRepositoryMock) Delete(ctx context.Context, id common.ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *CategoryRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Category, error) {
	args := m.Called(ctx, page, pageSize)
	return args.Get(0).([]*entity.Category), args.Error(1)
}

func (m *CategoryRepositoryMock) GetCategoryProducts(ctx context.Context, categoryID common.ID) ([]*entity.Product, error) {
	args := m.Called(ctx, categoryID)
	return args.Get(0).([]*entity.Product), args.Error(1)
}

func TestCategoryService_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		id := common.NewID()
		expectedCategory := &entity.Category{
			ID:   id,
			Name: "Category Test",
		}
		mockRepo.On("Get", mock.Anything, id).Return(expectedCategory, nil)

		category, err := categoryService.Get(context.Background(), id)

		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, category)
		mockRepo.AssertExpectations(t)
	})

	t.Run("category not found", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		id := common.NewID()
		mockRepo.On("Get", mock.Anything, id).Return(nil, &service.ErrNotFound{Entity: "Category", ID: id})

		category, err := categoryService.Get(context.Background(), id)

		assert.Error(t, err)
		assert.Nil(t, category)
		assert.IsType(t, &service.ErrNotFound{}, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCategoryService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		input := &dto.CreateCategoryInput{
			Name:        "Category Test",
			Description: "Description Test",
		}
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entity.Category")).Return(nil)

		_, err := categoryService.Create(context.Background(), input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid category", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		input := &dto.CreateCategoryInput{
			Name: "",
		}

		_, err := categoryService.Create(context.Background(), input)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error while validating category")
		mockRepo.AssertNotCalled(t, "Create")
	})
}

func TestCategoryService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		id := common.NewID()
		input := &dto.UpdateCategoryInput{
			ID:          id,
			Name:        "Category Test",
			Description: "Description Test",
		}
		expectedCategory := &entity.Category{
			ID:          id,
			Name:        "Category Test",
			Description: "Description Test",
		}
		mockRepo.On("Get", mock.Anything, id).Return(expectedCategory, nil)
		mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*entity.Category")).Return(nil)

		err := categoryService.Update(context.Background(), input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("category not found", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		input := &dto.UpdateCategoryInput{
			ID: common.NewID(),
		}
		mockRepo.On("Get", mock.Anything, input.ID).Return(nil, &service.ErrNotFound{Entity: "Category", ID: input.ID})

		err := categoryService.Update(context.Background(), input)
		assert.Error(t, err)
		assert.IsType(t, &service.ErrNotFound{}, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid category", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		id := common.NewID()
		input := &dto.UpdateCategoryInput{
			ID:   id,
			Name: "",
		}
		expectedCategory := &entity.Category{
			ID: id,
		}
		mockRepo.On("Get", mock.Anything, id).Return(expectedCategory, nil)

		err := categoryService.Update(context.Background(), input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error while validating category")
		mockRepo.AssertNotCalled(t, "Update")
	})
}

func TestCategoryService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		id := common.NewID()
		mockRepo.On("Get", mock.Anything, id).Return(&entity.Category{ID: id}, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(nil)

		err := categoryService.Delete(context.Background(), id)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("category not found", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		id := common.NewID()
		mockRepo.On("Get", mock.Anything, id).Return(nil, &service.ErrNotFound{Entity: "Category", ID: id})

		err := categoryService.Delete(context.Background(), id)
		assert.Error(t, err)
		assert.IsType(t, &service.ErrNotFound{}, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCategoryService_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		page := 1
		pageSize := 10
		expectedCategories := []*entity.Category{
			{
				ID:   common.NewID(),
				Name: "Category Test",
			},
		}
		mockRepo.On("List", mock.Anything, page, pageSize).Return(expectedCategories, nil)

		categories, err := categoryService.List(context.Background(), page, pageSize)

		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, categories)
		mockRepo.AssertExpectations(t)
	})
}

func TestCategoryService_GetCategoryProducts(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(CategoryRepositoryMock)
		categoryService := service.NewCategoryService(mockRepo)

		categoryID := common.NewID()
		expectedProducts := []*entity.Product{
			{
				ID:   common.NewID(),
				Name: "Product Test",
			},
		}
		mockRepo.On("GetCategoryProducts", mock.Anything, categoryID).Return(expectedProducts, nil)

		products, err := categoryService.GetCategoryProducts(context.Background(), categoryID)

		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, products)
		mockRepo.AssertExpectations(t)
	})
}

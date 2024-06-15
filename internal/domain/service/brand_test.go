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

type BrandRepositoryMock struct {
	mock.Mock
}

func (m *BrandRepositoryMock) Get(ctx context.Context, id common.ID) (*entity.Brand, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Brand), args.Error(1)
}

func (m *BrandRepositoryMock) Create(ctx context.Context, brand *entity.Brand) error {
	args := m.Called(ctx, brand)
	return args.Error(0)
}

func (m *BrandRepositoryMock) Update(ctx context.Context, brand *entity.Brand) error {
	args := m.Called(ctx, brand)
	return args.Error(0)
}

func (m *BrandRepositoryMock) Delete(ctx context.Context, id common.ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *BrandRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Brand, error) {
	args := m.Called(ctx, page, pageSize)
	return args.Get(0).([]*entity.Brand), args.Error(1)
}

func TestBrandService_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(BrandRepositoryMock)
		brandService := service.NewBrandService(mockRepo)

		id := common.NewID()
		expectedBrand := &entity.Brand{
			ID:   id,
			Name: "Brand Test",
		}
		mockRepo.On("Get", mock.Anything, id).Return(expectedBrand, nil)

		brand, err := brandService.Get(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, expectedBrand, brand)

		mockRepo.AssertExpectations(t)
	})
}

func TestBrandService_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(BrandRepositoryMock)
		brandService := service.NewBrandService(mockRepo)

		expectedBrands := []*entity.Brand{
			{
				ID:   common.NewID(),
				Name: "Brand Test 1",
			},
			{
				ID:   common.NewID(),
				Name: "Brand Test 2",
			},
		}
		mockRepo.On("List", mock.Anything, 1, 10).Return(expectedBrands, nil)

		brands, err := brandService.List(context.Background(), 1, 10)
		assert.NoError(t, err)
		assert.Equal(t, expectedBrands, brands)

		mockRepo.AssertExpectations(t)
	})
}

func TestBrandService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(BrandRepositoryMock)
		brandService := service.NewBrandService(mockRepo)

		input := &dto.CreateBrandInput{
			Name:        "Brand Test",
			Description: "Brand Test Description",
		}

		expectedBrand := &entity.Brand{
			ID:          common.NewID(),
			Name:        input.Name,
			Description: input.Description,
		}
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entity.Brand")).Return(nil)

		brand, err := brandService.Create(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, brand)
		assert.Equal(t, expectedBrand.Name, brand.Name)

		mockRepo.AssertExpectations(t)
	})
}

func TestBrandService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(BrandRepositoryMock)
		brandService := service.NewBrandService(mockRepo)

		input := &dto.UpdateBrandInput{
			ID:          common.NewID(),
			Name:        "Brand Test",
			Description: "Brand Test Description",
		}

		mockRepo.On("Get", mock.Anything, input.ID).Return(&entity.Brand{}, nil)
		mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*entity.Brand")).Return(nil)

		err := brandService.Update(context.Background(), input)
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
}

func TestBrandService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(BrandRepositoryMock)
		brandService := service.NewBrandService(mockRepo)

		id := common.NewID()
		mockRepo.On("Get", mock.Anything, id).Return(&entity.Brand{}, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(nil)

		err := brandService.Delete(context.Background(), id)
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
}

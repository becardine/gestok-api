package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type CategoryServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Category, error)
	Create(ctx context.Context, category *dto.CreateCategoryInput) (*entity.Category, error)
	Update(ctx context.Context, category *dto.UpdateCategoryInput) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Category, error)
	GetCategoryProducts(ctx context.Context, categoryID uuid.UUID) ([]*entity.Product, error)
	//AddCategoryProduct(ctx context.Context, productID, categoryID uuid.UUID) error
	//RemoveCategoryProduct(ctx context.Context, productID, categoryID uuid.UUID) error
}

type categoryService struct {
	categoryRepo domain.CategoryRepositoryInterface
}

func NewCategoryService(categoryRepo domain.CategoryRepositoryInterface) CategoryServiceInterface {
	return &categoryService{categoryRepo: categoryRepo}
}

func (c *categoryService) Get(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	category, err := c.categoryRepo.Get(ctx, id)
	if err != nil {
		return nil, c.handleCategoryError(err, id)
	}

	return category, nil
}

func (c *categoryService) Create(ctx context.Context, input *dto.CreateCategoryInput) (*entity.Category, error) {
	category := input.ToEntity()
	if err := category.Validate(); err != nil {
		return nil, fmt.Errorf("error while validating category: %w", err)
	}

	category, err := c.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("error while creating category in repository: %w", err)
	}

	return category, nil
}

func (c *categoryService) Update(ctx context.Context, input *dto.UpdateCategoryInput) error {
	category, err := c.categoryRepo.Get(ctx, input.ID)
	if err != nil {
		return c.handleCategoryError(err, input.ID)
	}

	category.Name = input.Name
	category.Description = input.Description

	if err := category.Validate(); err != nil {
		return fmt.Errorf("error while validating category: %w", err)
	}

	if err := c.categoryRepo.Update(ctx, category); err != nil {
		return fmt.Errorf("error while updating category in repository: %w", err)
	}

	return nil
}

func (c *categoryService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := c.Get(ctx, id)
	if err != nil {
		return err
	}

	if err := c.categoryRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error while deleting category in repository: %w", err)
	}

	return nil
}

func (c *categoryService) List(ctx context.Context, page, pageSize int) ([]*entity.Category, error) {
	categories, err := c.categoryRepo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("error while fetching categories from repository: %w", err)
	}

	return categories, nil
}

func (c *categoryService) GetCategoryProducts(ctx context.Context, categoryID uuid.UUID) ([]*entity.Product, error) {
	products, err := c.categoryRepo.GetCategoryProducts(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("error while fetching category products from repository: %w", err)
	}

	return products, nil
}

func (c *categoryService) handleCategoryError(err error, id uuid.UUID) error {
	if errors.Is(err, &ErrNotFound{}) {
		return &ErrNotFound{Entity: "Category", ID: id}
	}

	return fmt.Errorf("error while fetching category from repository: %w", err)
}

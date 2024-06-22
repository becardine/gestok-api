package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CategoryRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Category, error)
	Create(ctx context.Context, category *entity.Category) (*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Category, error)
	GetCategoryProducts(ctx context.Context, categoryID uuid.UUID) ([]*entity.Product, error)
	//AddCategoryProduct(ctx context.Context, productID, categoryID uuid.UUID) error
	//RemoveCategoryProduct(ctx context.Context, productID, categoryID uuid.UUID) error
}

package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type BrandRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Brand, error)
	Create(ctx context.Context, brand *entity.Brand) error
	Update(ctx context.Context, brand *entity.Brand) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Brand, error)
	// GetBrandProducts(ctx context.Context, brandID uuid.UUID) ([]*entity.Product, error)
	// AddBrandProduct(ctx context.Context, productID, brandID uuid.UUID) error
	// RemoveBrandProduct(ctx context.Context, productID, brandID uuid.UUID) error
}

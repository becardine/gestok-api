package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type BrandRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Brand, error)
	Create(ctx context.Context, brand *entity.Brand) error
	Update(ctx context.Context, brand *entity.Brand) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Brand, error)
	// GetBrandProducts(ctx context.Context, brandID common.ID) ([]*entity.Product, error)
	// AddBrandProduct(ctx context.Context, productID, brandID common.ID) error
	// RemoveBrandProduct(ctx context.Context, productID, brandID common.ID) error
}

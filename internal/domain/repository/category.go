package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type CategoryRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Category, error)
	Create(ctx context.Context, category *entity.Category) error
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Category, error)
	GetCategoryProducts(ctx context.Context, categoryID common.ID) ([]*entity.Product, error)
	//AddCategoryProduct(ctx context.Context, productID, categoryID common.ID) error
	//RemoveCategoryProduct(ctx context.Context, productID, categoryID common.ID) error
}

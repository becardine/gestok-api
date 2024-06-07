package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, id common.ID) (*entity.Product, error)
	CreateProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id common.ID) error
	ListProducts(ctx context.Context) ([]*entity.Product, error)
	// GetProductStocks(ctx context.Context, productID common.ID) ([]*entity.Stock, error)
	// AddProductStock(ctx context.Context, stockID, productID common.ID) error
	// RemoveProductStock(ctx context.Context, stockID, productID common.ID) error
}

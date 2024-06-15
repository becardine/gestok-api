package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type StockRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Stock, error)
	Create(ctx context.Context, stock *entity.Stock) (*entity.Stock, error)
	Update(ctx context.Context, stock *entity.Stock) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Stock, error)
	GetStockProducts(ctx context.Context, stockID common.ID) ([]*entity.Product, error)
	AddStockProduct(ctx context.Context, productID, stockID common.ID) error
	RemoveStockProduct(ctx context.Context, productID, stockID common.ID) error
}

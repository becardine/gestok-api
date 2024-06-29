package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type StockRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Stock, error)
	Create(ctx context.Context, stock *entity.Stock) (*entity.Stock, error)
	Update(ctx context.Context, stock *entity.Stock) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Stock, error)
	// GetStockProducts(ctx context.Context, stockID uuid.UUID) ([]*entity.Product, error)
	// AddStockProduct(ctx context.Context, productID, stockID uuid.UUID) error
	// RemoveStockProduct(ctx context.Context, productID, stockID uuid.UUID) error
}

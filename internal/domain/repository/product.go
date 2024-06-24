package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	CreateProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	ListProducts(ctx context.Context, page, pageSize int) ([]*entity.Product, error)
	// GetProductStocks(ctx context.Context, productID uuid.UUID) ([]*entity.Stock, error)
	// AddProductStock(ctx context.Context, stockID, productID uuid.UUID) error
	// RemoveProductStock(ctx context.Context, stockID, productID uuid.UUID) error
}

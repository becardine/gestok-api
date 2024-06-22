package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	domain "github.com/becardine/gestock-api/internal/domain/repository"

	"github.com/becardine/gestock-api/internal/domain/entity"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
	"github.com/google/uuid"
)

type ProductRepository struct {
	queries *database.Queries
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &ProductRepository{
		queries: database.New(db),
	}
}

// DeleteProduct implements repository.ProductRepository.
func (pr *ProductRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return nil
}

// GetProduct implements repository.ProductRepository.
func (pr *ProductRepository) GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	return nil, nil
}

// ListProducts implements repository.ProductRepository.
func (pr *ProductRepository) ListProducts(ctx context.Context) ([]*entity.Product, error) {
	return nil, nil
}

// UpdateProduct implements repository.ProductRepository.
func (pr *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {
	return nil
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) error {

	err := pr.queries.CreateProduct(ctx, database.CreateProductParams{
		ID:              product.ID,
		Name:            product.Name,
		Description:     sql.NullString{String: product.Description, Valid: true},
		Price:           fmt.Sprintf("%.2f", product.Price),
		QuantityInStock: int32(product.QuantityInStock),
		ImageUrl:        sql.NullString{String: product.ImageURL, Valid: true},
		CategoryID: uuid.NullUUID{
			Valid: product.CategoryID != uuid.Nil,
			UUID:  product.CategoryID,
		},
		BrandID: uuid.NullUUID{
			Valid: product.BrandID != uuid.Nil,
			UUID:  product.BrandID,
		},
		CreatedDate: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedDate: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}

	return nil
}

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
	err := pr.queries.DeleteProduct(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	return nil
}

// GetProduct implements repository.ProductRepository.
func (pr *ProductRepository) GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	product, err := pr.queries.GetProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	return &entity.Product{
		ID:              product.ID,
		Name:            product.Name,
		Description:     product.Description.String,
		Price:           product.Price,
		QuantityInStock: int(product.QuantityInStock),
		ImageURL:        product.ImageUrl.String,
		CategoryID:      product.CategoryID,
		BrandID:         product.BrandID,
	}, nil

}

// ListProducts implements repository.ProductRepository.
func (pr *ProductRepository) ListProducts(ctx context.Context, page, pageSize int) ([]*entity.Product, error) {
	// TODO: Implement pagination
	products, err := pr.queries.ListProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %v", err)
	}

	result := make([]*entity.Product, len(products))
	for i, product := range products {
		result[i] = &entity.Product{
			ID:              product.ID,
			Name:            product.Name,
			Description:     product.Description.String,
			Price:           product.Price,
			QuantityInStock: int(product.QuantityInStock),
			ImageURL:        product.ImageUrl.String,
			CategoryID:      product.CategoryID,
			BrandID:         product.BrandID,
		}
	}

	return result, nil
}

// UpdateProduct implements repository.ProductRepository.
func (pr *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {
	err := pr.queries.UpdateProduct(ctx, database.UpdateProductParams{
		ID:              product.ID,
		Name:            product.Name,
		Description:     sql.NullString{String: product.Description, Valid: true},
		Price:           product.Price,
		QuantityInStock: int32(product.QuantityInStock),
		ImageUrl:        sql.NullString{String: product.ImageURL, Valid: true},
		CategoryID:      product.CategoryID,
		BrandID:         product.BrandID,
		UpdatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}

	return nil
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) error {
	err := pr.queries.CreateProduct(ctx, database.CreateProductParams{
		ID:              product.ID,
		Name:            product.Name,
		Description:     sql.NullString{String: product.Description, Valid: true},
		Price:           product.Price,
		QuantityInStock: int32(product.QuantityInStock),
		ImageUrl:        sql.NullString{String: product.ImageURL, Valid: true},
		CategoryID:      product.CategoryID,
		BrandID:         product.BrandID,
		CreatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}

	return nil
}

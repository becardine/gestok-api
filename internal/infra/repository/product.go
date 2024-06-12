package repository

import (
	"context"
	"database/sql"
	"fmt"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
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
func (pr *ProductRepository) DeleteProduct(ctx context.Context, id common.ID) error {
	return nil
}

// GetProduct implements repository.ProductRepository.
func (pr *ProductRepository) GetProduct(ctx context.Context, id common.ID) (*entity.Product, error) {
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
			Valid: !product.CategoryID.IsEmpty(),
			UUID:  product.CategoryID.Value(),
		},
		BrandID: uuid.NullUUID{
			Valid: !product.BrandID.IsEmpty(),
			UUID:  product.BrandID.Value(),
		},
		CreatedDate: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedDate: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}

	return nil
}

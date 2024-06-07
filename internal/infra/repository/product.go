package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
	"github.com/google/uuid"
)

type ProductRepository struct {
	queries *database.Queries
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		queries: database.New(db),
	}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) error {
	err := pr.queries.CreateProduct(ctx, database.CreateProductParams{
		ID:              common.NewID(),
		Name:            product.Name,
		Description:     sql.NullString{String: product.Description, Valid: true},
		Price:           fmt.Sprintf("%.2f", product.Price),
		QuantityInStock: int32(product.QuantityInStock),
		ImageUrl:        sql.NullString{String: product.ImageURL, Valid: true},
		CategoryID: uuid.NullUUID{
			Valid: product.CategoryID != common.NewID(),
			UUID:  product.CategoryID.Value(),
		},
		BrandID: uuid.NullUUID{
			Valid: product.BrandID != common.NewID(),
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

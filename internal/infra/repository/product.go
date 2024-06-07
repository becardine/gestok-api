package repository

import (
	"context"
	"database/sql"

	"github.com/becardine/gestock-api/internal/domain/entity"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
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
	// err := pr.queries.CreateProduct(ctx, database.CreateProductParams{
	// 	ID:              product.ID.Value(),
	// 	Name:            product.Name,
	// 	Description:     sql.NullString{String: product.Description, Valid: true},
	// 	Price:           product.Price,
	// 	QuantityInStock: int32(product.QuantityInStock),
	// 	ImageURL:        product.ImageURL,
	// 	CategoryID:      product.CategoryID.Value(),
	// 	BrandID:         product.BrandID.Value(),
	// 	CreatedDate:     product.CreatedAt,
	// 	UpdatedDate:     product.UpdatedAt,
	// })
	// if err != nil {
	// 	return fmt.Errorf("erro ao criar produto: %w", err)
	// }

	return nil
}

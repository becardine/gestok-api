package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
	"github.com/google/uuid"
)

type CategoryRepository struct {
	queries *database.Queries
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepositoryInterface {
	return &CategoryRepository{
		queries: database.New(db),
	}
}

func (c CategoryRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	category, err := c.queries.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description.String,
	}, nil
}

func (c CategoryRepository) Create(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	err := c.queries.CreateCategory(ctx, database.CreateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: sql.NullString{String: category.Description, Valid: true},
		CreatedDate: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedDate: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c CategoryRepository) Update(ctx context.Context, category *entity.Category) error {
	err := c.queries.UpdateCategory(ctx, database.UpdateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: sql.NullString{String: category.Description, Valid: true},
		UpdatedDate: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.queries.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) List(ctx context.Context, page, pageSize int) ([]*entity.Category, error) {
	categories, err := c.queries.ListCategories(ctx, database.ListCategoriesParams{
		Limit:  int32(pageSize),
		Offset: int32((page - 1) * pageSize),
	})
	if err != nil {
		return nil, err
	}

	var result []*entity.Category
	for _, category := range categories {
		result = append(result, &entity.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description.String,
		})
	}

	return result, nil
}

func (c CategoryRepository) GetCategoryProducts(ctx context.Context, categoryID uuid.UUID) ([]*entity.Product, error) {
	products, err := c.queries.GetCategoryProducts(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	var result []*entity.Product
	for _, product := range products {
		result = append(result, &entity.Product{
			ID:              product.ID,
			Name:            product.Name,
			Description:     product.Description.String,
			QuantityInStock: int(product.QuantityInStock),
		})
	}

	return result, nil
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
	"github.com/google/uuid"
)

type BrandRepository struct {
	queries *database.Queries
}

func NewBrandRepository(db *sql.DB) domain.BrandRepositoryInterface {
	return &BrandRepository{
		queries: database.New(db),
	}
}

func (b *BrandRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Brand, error) {
	brand, err := b.queries.GetBrand(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Brand{
		ID:   brand.ID,
		Name: brand.Name,
	}, nil
}

func (b *BrandRepository) Create(ctx context.Context, brand *entity.Brand) error {
	err := b.queries.CreateBrand(ctx, database.CreateBrandParams{
		ID:          brand.ID,
		Name:        brand.Name,
		CreatedDate: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedDate: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("error creating brand: %w", err)
	}

	return nil
}

func (b *BrandRepository) Update(ctx context.Context, brand *entity.Brand) error {
	err := b.queries.UpdateBrand(ctx, database.UpdateBrandParams{
		ID:          brand.ID,
		Name:        brand.Name,
		UpdatedDate: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("error updating brand: %w", err)
	}

	return nil
}

func (b *BrandRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := b.queries.DeleteBrand(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting brand: %w", err)
	}

	return nil
}

func (b *BrandRepository) List(ctx context.Context, page, pageSize int) ([]*entity.Brand, error) {
	brands, err := b.queries.ListBrands(ctx, database.ListBrandsParams{
		Limit:  int32(pageSize),
		Offset: int32((page - 1) * pageSize),
	})
	if err != nil {
		return nil, err
	}

	var result []*entity.Brand
	for _, brand := range brands {
		result = append(result, &entity.Brand{
			ID:   brand.ID,
			Name: brand.Name,
		})
	}

	return result, nil
}

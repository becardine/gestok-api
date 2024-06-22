package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type BrandService interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Brand, error)
	List(ctx context.Context, page, pageSize int) ([]*entity.Brand, error)
	Create(ctx context.Context, brand *dto.CreateBrandInput) (*entity.Brand, error)
	Update(ctx context.Context, brand *dto.UpdateBrandInput) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type brandService struct {
	repository domain.BrandRepositoryInterface
}

func NewBrandService(brandRepo domain.BrandRepositoryInterface) BrandService {
	return &brandService{
		repository: brandRepo,
	}
}

func (b *brandService) Get(ctx context.Context, id uuid.UUID) (*entity.Brand, error) {
	brand, err := b.repository.Get(ctx, id)
	if err != nil {
		return nil, b.handleBrandNotFound(err, id)
	}

	return brand, nil
}

func (b *brandService) List(ctx context.Context, page, pageSize int) ([]*entity.Brand, error) {
	brands, err := b.repository.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("error while fetching brands from repository: %w", err)
	}

	return brands, nil
}

func (b *brandService) Create(ctx context.Context, input *dto.CreateBrandInput) (*entity.Brand, error) {
	brand := input.ToEntity()

	if err := brand.Validate(); err != nil {
		return nil, fmt.Errorf("error while validating brand entity: %w", err)
	}

	err := b.repository.Create(ctx, brand)
	if err != nil {
		return nil, fmt.Errorf("error while creating brand in repository: %w", err)
	}

	return brand, nil
}

func (b *brandService) Update(ctx context.Context, input *dto.UpdateBrandInput) error {
	brand, err := b.repository.Get(ctx, input.ID)
	if err != nil {
		return b.handleBrandNotFound(err, input.ID)
	}

	brand.Name = input.Name
	brand.Description = input.Description

	err = b.repository.Update(ctx, brand)
	if err != nil {
		return fmt.Errorf("error while updating brand in repository: %w", err)
	}

	return nil
}

func (b *brandService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := b.repository.Get(ctx, id)
	if err != nil {
		return b.handleBrandNotFound(err, id)
	}

	err = b.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("error while deleting brand in repository: %w", err)
	}

	return nil
}

func (b *brandService) handleBrandNotFound(err error, id uuid.UUID) error {
	if errors.Is(err, &ErrNotFound{}) {
		return &ErrNotFound{Entity: "Brand", ID: id}
	}

	return fmt.Errorf("error while fetching brand from repository: %w", err)
}

package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
)

type BrandService interface {
	Get(ctx context.Context, id common.ID) (*entity.Brand, error)
	List(ctx context.Context) ([]*entity.Brand, error)
	Create(ctx context.Context, brand *entity.Brand) error
	Update(ctx context.Context, brand *entity.Brand) error
	Delete(ctx context.Context, id common.ID) error
}

type brandService struct {
	repository domain.BrandRepositoryInterface
}

func NewBrandService(brandRepo domain.BrandRepositoryInterface) BrandService {
	return &brandService{
		repository: brandRepo,
	}
}

func (b *brandService) Get(ctx context.Context, id common.ID) (*entity.Brand, error) {
	brand, err := b.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func (b *brandService) List(ctx context.Context) ([]*entity.Brand, error) {
	//TODO implement me
	panic("implement me")
}

func (b *brandService) Create(ctx context.Context, brand *entity.Brand) error {
	//TODO implement me
	panic("implement me")
}

func (b *brandService) Update(ctx context.Context, brand *entity.Brand) error {
	//TODO implement me
	panic("implement me")
}

func (b *brandService) Delete(ctx context.Context, id common.ID) error {
	//TODO implement me
	panic("implement me")
}

func (b *brandService) handleBrandNotFound(err error, id common.ID) error {
	if errors.Is(err, &ErrNotFound{}) {
		return &ErrNotFound{Entity: "Brand", ID: id}
	}

	return fmt.Errorf("error while fetching brand from repository: %w", err)
}

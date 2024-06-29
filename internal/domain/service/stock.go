package service

import (
	"context"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type StockServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Stock, error)
	Create(ctx context.Context, input *dto.CreateStockInput) (*entity.Stock, error)
	Update(ctx context.Context, input *dto.UpdateStockInput) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Stock, error)
}

type stockService struct {
	repo domain.StockRepositoryInterface
}

func NewStockService(repo domain.StockRepositoryInterface) StockServiceInterface {
	return &stockService{
		repo: repo,
	}
}

// Create implements StockService.
func (s *stockService) Create(ctx context.Context, input *dto.CreateStockInput) (*entity.Stock, error) {
	stock := input.ToEntity()
	if err := stock.Validate(); err != nil {
		return nil, fmt.Errorf("error while validating stock entity: %w", err)
	}

	stock, err := s.repo.Create(ctx, stock)
	if err != nil {
		return nil, fmt.Errorf("error while creating stock in repository: %w", err)
	}

	return stock, nil
}

// Delete implements StockService.
func (s *stockService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("error while getting stock in repository: %w", err)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error while deleting stock in repository: %w", err)
	}

	return nil
}

// Get implements StockService.
func (s *stockService) Get(ctx context.Context, id uuid.UUID) (*entity.Stock, error) {
	stock, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error while getting stock in repository: %w", err)
	}

	return stock, nil
}

// List implements StockService.
func (s *stockService) List(ctx context.Context, page int, pageSize int) ([]*entity.Stock, error) {
	stocks, err := s.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("error while listing stocks in repository: %w", err)
	}

	return stocks, nil
}

// Update implements StockService.
func (s *stockService) Update(ctx context.Context, input *dto.UpdateStockInput) error {
	_, err := s.repo.Get(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("error while getting stock in repository: %w", err)
	}

	stock := input.ToEntity()
	if err := stock.Validate(); err != nil {
		return fmt.Errorf("error while validating stock entity: %w", err)
	}

	if err := s.repo.Update(ctx, stock); err != nil {
		return fmt.Errorf("error while updating stock in repository: %w", err)
	}

	return nil
}

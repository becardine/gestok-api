package service

import (
	"context"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type DeliveryServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Delivery, error)
	Create(ctx context.Context, input *dto.CreateDeliveryInput) (*entity.Delivery, error)
	Update(ctx context.Context, id uuid.UUID, input *dto.UpdateDeliveryInput) (*entity.Delivery, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Delivery, error)
}

type deliveryService struct {
	repo domain.DeliveryRepositoryInterface
}

func NewDeliveryService(repo domain.DeliveryRepositoryInterface) DeliveryServiceInterface {
	return &deliveryService{
		repo: repo,
	}
}

// Create implements DeliveryServiceInterface.
func (d *deliveryService) Create(ctx context.Context, input *dto.CreateDeliveryInput) (*entity.Delivery, error) {
	delivery := input.ToEntity()
	if err := delivery.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate delivery: %w", err)
	}

	delivery, err := d.repo.Create(ctx, delivery)
	if err != nil {
		return nil, fmt.Errorf("failed to create delivery: %w", err)
	}

	return delivery, nil
}

// Delete implements DeliveryServiceInterface.
func (d *deliveryService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := d.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get delivery: %w", err)
	}

	if err := d.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete delivery: %w", err)
	}

	return nil
}

// Get implements DeliveryServiceInterface.
func (d *deliveryService) Get(ctx context.Context, id uuid.UUID) (*entity.Delivery, error) {
	delivery, err := d.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get delivery: %w", err)
	}

	return delivery, nil
}

// List implements DeliveryServiceInterface.
func (d *deliveryService) List(ctx context.Context, page int, pageSize int) ([]*entity.Delivery, error) {
	deliveries, err := d.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to list deliveries: %w", err)
	}

	return deliveries, nil
}

// Update implements DeliveryServiceInterface.
func (d *deliveryService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateDeliveryInput) (*entity.Delivery, error) {
	_, err := d.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get delivery: %w", err)
	}

	delivery := input.ToEntity()
	if err := delivery.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate delivery: %w", err)
	}

	if err := d.repo.Update(ctx, delivery); err != nil {
		return nil, fmt.Errorf("failed to update delivery: %w", err)
	}

	return delivery, nil
}

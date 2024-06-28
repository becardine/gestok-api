package service

import (
	"context"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type OrderServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Order, error)
	Create(ctx context.Context, input *dto.CreateOrderInput) (*entity.Order, error)
	Update(ctx context.Context, id uuid.UUID, input *dto.UpdateOrderInput) (*entity.Order, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Order, error)
}

type orderService struct {
	repo domain.OrderRepositoryInterface
}

func NewOrderService(repo domain.OrderRepositoryInterface) OrderServiceInterface {
	return &orderService{
		repo: repo,
	}
}

// Create implements OrderServiceInterface.
func (o *orderService) Create(ctx context.Context, input *dto.CreateOrderInput) (*entity.Order, error) {
	order := input.ToEntity()
	if err := order.Validate(); err != nil {
		return nil, fmt.Errorf("error validating order: %w", err)
	}

	order, err := o.repo.Create(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("error creating order: %w", err)
	}

	return order, nil
}

// Delete implements OrderServiceInterface.
func (o *orderService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := o.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("error getting order: %w", err)
	}

	if err := o.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting order: %w", err)
	}

	return nil
}

// Get implements OrderServiceInterface.
func (o *orderService) Get(ctx context.Context, id uuid.UUID) (*entity.Order, error) {
	order, err := o.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting order: %w", err)
	}

	return order, nil
}

// List implements OrderServiceInterface.
func (o *orderService) List(ctx context.Context, page, pageSize int) ([]*entity.Order, error) {
	orders, err := o.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("error listing orders: %w", err)
	}

	return orders, nil
}

// Update implements OrderServiceInterface.
func (o *orderService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateOrderInput) (*entity.Order, error) {
	_, err := o.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting order: %w", err)
	}

	order := input.ToEntity()
	if err := order.Validate(); err != nil {
		return nil, fmt.Errorf("error validating order: %w", err)
	}

	if err := o.repo.Update(ctx, order); err != nil {
		return nil, fmt.Errorf("error updating order: %w", err)
	}

	return order, nil
}

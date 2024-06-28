package service

import (
	"context"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type OrderItemServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.OrderItem, error)
	Create(ctx context.Context, input *dto.CreateOrderItemInput) (*entity.OrderItem, error)
	Update(ctx context.Context, id uuid.UUID, input *dto.UpdateOrderItemInput) (*entity.OrderItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.OrderItem, error)
	GetOrderItemProduct(ctx context.Context, orderItemID uuid.UUID) (*entity.Product, error)
	SetOrderItemProduct(ctx context.Context, productID, orderItemID uuid.UUID) error
	GetOrderItemOrder(ctx context.Context, orderItemID uuid.UUID) (*entity.Order, error)
	SetOrderItemOrder(ctx context.Context, orderID, orderItemID uuid.UUID) error
}

type orderItemService struct {
	repo domain.OrderItemRepositoryInterface
}

func NewOrderItemService(repo domain.OrderItemRepositoryInterface) OrderItemServiceInterface {
	return &orderItemService{
		repo: repo,
	}
}

// Create implements OrderItemServiceInterface.
func (o *orderItemService) Create(ctx context.Context, input *dto.CreateOrderItemInput) (*entity.OrderItem, error) {
	orderItem := input.ToEntity()
	if err := orderItem.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate order item: %w", err)
	}

	orderItem, err := o.repo.Create(ctx, orderItem)
	if err != nil {
		return nil, fmt.Errorf("failed to create order item: %w", err)
	}

	return orderItem, nil
}

// Delete implements OrderItemServiceInterface.
func (o *orderItemService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := o.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get order item: %w", err)
	}

	if err := o.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete order item: %w", err)
	}

	return nil
}

// Get implements OrderItemServiceInterface.
func (o *orderItemService) Get(ctx context.Context, id uuid.UUID) (*entity.OrderItem, error) {
	orderItem, err := o.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}

	return orderItem, nil
}

// GetOrderItemOrder implements OrderItemServiceInterface.
func (o *orderItemService) GetOrderItemOrder(ctx context.Context, orderItemID uuid.UUID) (*entity.Order, error) {
	orderItem, err := o.Get(ctx, orderItemID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}

	order, err := o.repo.GetOrderItemOrder(ctx, orderItem.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order item order: %w", err)
	}

	return order, nil
}

// GetOrderItemProduct implements OrderItemServiceInterface.
func (o *orderItemService) GetOrderItemProduct(ctx context.Context, orderItemID uuid.UUID) (*entity.Product, error) {
	orderItem, err := o.Get(ctx, orderItemID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}

	product, err := o.repo.GetOrderItemProduct(ctx, orderItem.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order item product: %w", err)
	}

	return product, nil
}

// List implements OrderItemServiceInterface.
func (o *orderItemService) List(ctx context.Context, page int, pageSize int) ([]*entity.OrderItem, error) {
	orderItems, err := o.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to list order items: %w", err)
	}

	return orderItems, nil
}

// SetOrderItemOrder implements OrderItemServiceInterface.
func (o *orderItemService) SetOrderItemOrder(ctx context.Context, orderID uuid.UUID, orderItemID uuid.UUID) error {
	orderItem, err := o.Get(ctx, orderItemID)
	if err != nil {
		return fmt.Errorf("failed to get order item: %w", err)
	}

	order, err := o.repo.GetOrderItemOrder(ctx, orderItem.ID)
	if err != nil {
		return fmt.Errorf("failed to get order item order: %w", err)
	}

	if order != nil {
		return fmt.Errorf("order item already has an order")
	}

	if err := o.repo.SetOrderItemOrder(ctx, orderID, orderItemID); err != nil {
		return fmt.Errorf("failed to set order item order: %w", err)
	}

	return nil
}

// SetOrderItemProduct implements OrderItemServiceInterface.
func (o *orderItemService) SetOrderItemProduct(ctx context.Context, productID uuid.UUID, orderItemID uuid.UUID) error {
	orderItem, err := o.Get(ctx, orderItemID)
	if err != nil {
		return fmt.Errorf("failed to get order item: %w", err)
	}

	product, err := o.repo.GetOrderItemProduct(ctx, orderItem.ID)
	if err != nil {
		return fmt.Errorf("failed to get order item product: %w", err)
	}

	if product != nil {
		return fmt.Errorf("order item already has a product")
	}

	if err := o.repo.SetOrderItemProduct(ctx, productID, orderItemID); err != nil {
		return fmt.Errorf("failed to set order item product: %w", err)
	}

	return nil
}

// Update implements OrderItemServiceInterface.
func (o *orderItemService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateOrderItemInput) (*entity.OrderItem, error) {
	_, err := o.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order item: %w", err)
	}

	orderItem := input.ToEntity()
	if err := orderItem.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate order item: %w", err)
	}

	if err := o.repo.Update(ctx, orderItem); err != nil {
		return nil, fmt.Errorf("failed to update order item: %w", err)
	}

	return orderItem, nil
}

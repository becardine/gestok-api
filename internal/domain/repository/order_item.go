package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type OrderItemRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.OrderItem, error)
	Create(ctx context.Context, order *entity.OrderItem) (*entity.OrderItem, error)
	Update(ctx context.Context, order *entity.OrderItem) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.OrderItem, error)
	GetOrderItemProduct(ctx context.Context, orderItemID uuid.UUID) (*entity.Product, error)
	SetOrderItemProduct(ctx context.Context, productID, orderItemID uuid.UUID) error
	GetOrderItemOrder(ctx context.Context, orderItemID uuid.UUID) (*entity.Order, error)
	SetOrderItemOrder(ctx context.Context, orderID, orderItemID uuid.UUID) error
}

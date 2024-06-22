package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type DeliveryRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Delivery, error)
	Create(ctx context.Context, delivery *entity.Delivery) (*entity.Delivery, error)
	Update(ctx context.Context, delivery *entity.Delivery) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Delivery, error)
	GetDeliveryOrders(ctx context.Context, deliveryID uuid.UUID) ([]*entity.Order, error)
	AddDeliveryOrder(ctx context.Context, orderID, deliveryID uuid.UUID) error
	RemoveDeliveryOrder(ctx context.Context, orderID, deliveryID uuid.UUID) error
}

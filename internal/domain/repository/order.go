package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type OrderRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Order, error)
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Order, error)
	GetOrderProducts(ctx context.Context, orderID uuid.UUID) ([]*entity.Product, error)
	AddOrderProduct(ctx context.Context, productID, orderID uuid.UUID) error
	RemoveOrderProduct(ctx context.Context, productID, orderID uuid.UUID) error
	GetOrderCustomer(ctx context.Context, orderID uuid.UUID) (*entity.Customer, error)
	SetOrderCustomer(ctx context.Context, customerID, orderID uuid.UUID) error
	GetOrderDelivery(ctx context.Context, orderID uuid.UUID) (*entity.Delivery, error)
	SetOrderDelivery(ctx context.Context, deliveryID, orderID uuid.UUID) error
	GetOrderFeedback(ctx context.Context, orderID uuid.UUID) (*entity.Feedback, error)
	SetOrderFeedback(ctx context.Context, feedbackID, orderID uuid.UUID) error
}

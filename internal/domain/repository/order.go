package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type OrderRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Order, error)
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Order, error)
	GetOrderProducts(ctx context.Context, orderID common.ID) ([]*entity.Product, error)
	AddOrderProduct(ctx context.Context, productID, orderID common.ID) error
	RemoveOrderProduct(ctx context.Context, productID, orderID common.ID) error
	GetOrderCustomer(ctx context.Context, orderID common.ID) (*entity.Customer, error)
	SetOrderCustomer(ctx context.Context, customerID, orderID common.ID) error
	GetOrderDelivery(ctx context.Context, orderID common.ID) (*entity.Delivery, error)
	SetOrderDelivery(ctx context.Context, deliveryID, orderID common.ID) error
	GetOrderFeedback(ctx context.Context, orderID common.ID) (*entity.Feedback, error)
	SetOrderFeedback(ctx context.Context, feedbackID, orderID common.ID) error
}

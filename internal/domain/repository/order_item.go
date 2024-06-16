package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type OrderItemRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.OrderItem, error)
	Create(ctx context.Context, order *entity.OrderItem) (*entity.OrderItem, error)
	Update(ctx context.Context, order *entity.OrderItem) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.OrderItem, error)
	GetOrderItemProduct(ctx context.Context, orderItemID common.ID) (*entity.Product, error)
	SetOrderItemProduct(ctx context.Context, productID, orderItemID common.ID) error
	GetOrderItemOrder(ctx context.Context, orderItemID common.ID) (*entity.Order, error)
	SetOrderItemOrder(ctx context.Context, orderID, orderItemID common.ID) error
}

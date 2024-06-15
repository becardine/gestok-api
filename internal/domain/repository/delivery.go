package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type DeliveryRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Delivery, error)
	Create(ctx context.Context, delivery *entity.Delivery) (*entity.Delivery, error)
	Update(ctx context.Context, delivery *entity.Delivery) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Delivery, error)
	GetDeliveryOrders(ctx context.Context, deliveryID common.ID) ([]*entity.Order, error)
	AddDeliveryOrder(ctx context.Context, orderID, deliveryID common.ID) error
	RemoveDeliveryOrder(ctx context.Context, orderID, deliveryID common.ID) error
}

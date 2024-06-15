package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type PaymentRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Payment, error)
	Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Payment, error)
	GetPaymentOrders(ctx context.Context, paymentID common.ID) ([]*entity.Order, error)
	AddPaymentOrder(ctx context.Context, orderID, paymentID common.ID) error
	RemovePaymentOrder(ctx context.Context, orderID, paymentID common.ID) error
}

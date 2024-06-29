package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type PaymentRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Payment, error)
	Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Payment, error)
	// GetPaymentOrders(ctx context.Context, paymentID uuid.UUID) ([]*entity.Order, error)
	// AddPaymentOrder(ctx context.Context, orderID, paymentID uuid.UUID) error
	// RemovePaymentOrder(ctx context.Context, orderID, paymentID uuid.UUID) error
}

package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CustomerRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Customer, error)
	Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error)
	Update(ctx context.Context, customer *entity.Customer) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Customer, error)
	GetCustomerOrders(ctx context.Context, customerID uuid.UUID) ([]*entity.Order, error)
	AddCustomerOrder(ctx context.Context, orderID, customerID uuid.UUID) error
	RemoveCustomerOrder(ctx context.Context, orderID, customerID uuid.UUID) error
}

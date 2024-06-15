package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type CustomerRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Customer, error)
	Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error)
	Update(ctx context.Context, customer *entity.Customer) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Customer, error)
	GetCustomerOrders(ctx context.Context, customerID common.ID) ([]*entity.Order, error)
	AddCustomerOrder(ctx context.Context, orderID, customerID common.ID) error
	RemoveCustomerOrder(ctx context.Context, orderID, customerID common.ID) error
}

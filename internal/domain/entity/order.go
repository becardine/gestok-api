package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
	"time"
)

type Order struct {
	ID          common.ID `json:"id"`
	CustomerID  common.ID `json:"customer_id"`
	OrderDate   time.Time `json:"order_date"`
	OrderStatus string    `json:"order_status"`
	TotalValue  float64   `json:"total_value"`
}

func NewOrder(customerID common.ID, orderDate time.Time, orderStatus string, totalValue float64) (*Order, error) {
	order := &Order{
		ID:          common.NewID(),
		CustomerID:  customerID,
		OrderDate:   orderDate,
		OrderStatus: orderStatus,
		TotalValue:  totalValue,
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if o.CustomerID.IsEmpty() {
		return errors.NewEntityValidationError("customer_id", "required", "")
	}

	if o.OrderDate.IsZero() {
		return errors.NewEntityValidationError("order_date", "required", "")
	}

	if o.OrderStatus == "" {
		return errors.NewEntityValidationError("order_status", "required", "")
	}

	if o.TotalValue <= 0 {
		return errors.NewEntityValidationError("total_value", "invalid_range", "> 0")
	}

	return nil
}

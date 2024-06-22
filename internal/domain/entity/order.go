package entity

import (
	"time"

	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `json:"id"`
	CustomerID  uuid.UUID `json:"customer_id"`
	OrderDate   time.Time `json:"order_at"`
	OrderStatus string    `json:"order_status"`
	TotalValue  float64   `json:"total_value"`
}

func NewOrder(customerID uuid.UUID, orderDate time.Time, orderStatus string, totalValue float64) (*Order, error) {
	order := &Order{
		ID:          uuid.New(),
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
	if o.CustomerID == uuid.Nil {
		return errors.NewEntityValidationError("customer_id", "required", "")
	}

	if o.OrderDate.IsZero() {
		return errors.NewEntityValidationError("order_at", "required", "")
	}

	if o.OrderStatus == "" {
		return errors.NewEntityValidationError("order_status", "required", "")
	}

	if o.TotalValue <= 0 {
		return errors.NewEntityValidationError("total_value", "invalid_range", "> 0")
	}

	return nil
}

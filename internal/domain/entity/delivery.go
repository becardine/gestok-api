package entity

import (
	"time"

	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
)

type Delivery struct {
	ID             uuid.UUID `json:"id"`
	OrderID        uuid.UUID `json:"order_id"`
	CustomerID     uuid.UUID `json:"customer_id"`
	DeliveryType   string    `json:"delivery_type"`
	DeliveryDate   time.Time `json:"delivery_at"`
	DeliveryStatus string    `json:"delivery_status"`
}

func NewDelivery(orderID, customerID uuid.UUID, deliveryType, deliveryStatus string, deliveryDate time.Time) (*Delivery, error) {
	return &Delivery{
		ID:             uuid.New(),
		OrderID:        orderID,
		CustomerID:     customerID,
		DeliveryType:   deliveryType,
		DeliveryDate:   deliveryDate,
		DeliveryStatus: deliveryStatus,
	}, nil
}

func (d *Delivery) Validate() error {
	if d.OrderID == uuid.Nil {
		return errors.NewEntityValidationError("order_id", "required", "")
	}

	if d.CustomerID == uuid.Nil {
		return errors.NewEntityValidationError("customer_id", "required", "")
	}

	if d.DeliveryType == "" {
		return errors.NewEntityValidationError("delivery_type", "required", "")
	}

	if d.DeliveryDate.IsZero() {
		return errors.NewEntityValidationError("delivery_at", "required", "")
	}

	if d.DeliveryStatus == "" {
		return errors.NewEntityValidationError("delivery_status", "required", "")
	}

	return nil
}

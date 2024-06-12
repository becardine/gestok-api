package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
	"time"
)

type Delivery struct {
	ID             common.ID `json:"id"`
	OrderID        common.ID `json:"order_id"`
	CustomerID     common.ID `json:"customer_id"`
	DeliveryType   string    `json:"delivery_type"`
	DeliveryDate   time.Time `json:"delivery_date"`
	DeliveryStatus string    `json:"delivery_status"`
}

func NewDelivery(orderID, customerID common.ID, deliveryType, deliveryStatus string, deliveryDate time.Time) (*Delivery, error) {
	return &Delivery{
		ID:             common.NewID(),
		OrderID:        orderID,
		CustomerID:     customerID,
		DeliveryType:   deliveryType,
		DeliveryDate:   deliveryDate,
		DeliveryStatus: deliveryStatus,
	}, nil
}

func (d *Delivery) Validate() error {
	if d.OrderID.IsEmpty() {
		return errors.NewEntityValidationError("order_id", "required", "")
	}

	if d.CustomerID.IsEmpty() {
		return errors.NewEntityValidationError("customer_id", "required", "")
	}

	if d.DeliveryType == "" {
		return errors.NewEntityValidationError("delivery_type", "required", "")
	}

	if d.DeliveryDate.IsZero() {
		return errors.NewEntityValidationError("delivery_date", "required", "")
	}

	if d.DeliveryStatus == "" {
		return errors.NewEntityValidationError("delivery_status", "required", "")
	}

	return nil
}

package dto

import (
	"encoding/json"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateDeliveryInput struct {
	OrderID        uuid.UUID `json:"order_id"`
	CustomerID     uuid.UUID `json:"customer_id"`
	DeliveryType   string    `json:"delivery_type"`
	DeliveryDate   time.Time `json:"delivery_date"`
	DeliveryStatus string    `json:"delivery_status"`
}

type UpdateDeliveryInput struct {
	ID             uuid.UUID `json:"id"`
	OrderID        uuid.UUID `json:"order_id"`
	CustomerID     uuid.UUID `json:"customer_id"`
	DeliveryType   string    `json:"delivery_type"`
	DeliveryDate   time.Time `json:"delivery_date"`
	DeliveryStatus string    `json:"delivery_status"`
}

func (input *CreateDeliveryInput) ToEntity() *entity.Delivery {
	return &entity.Delivery{
		OrderID:        input.OrderID,
		CustomerID:     input.CustomerID,
		DeliveryType:   input.DeliveryType,
		DeliveryDate:   input.DeliveryDate,
		DeliveryStatus: input.DeliveryStatus,
	}
}

func (input *CreateDeliveryInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateDeliveryInput) ToEntity() *entity.Delivery {
	return &entity.Delivery{
		ID:             input.ID,
		OrderID:        input.OrderID,
		CustomerID:     input.CustomerID,
		DeliveryType:   input.DeliveryType,
		DeliveryDate:   input.DeliveryDate,
		DeliveryStatus: input.DeliveryStatus,
	}
}

func (input *UpdateDeliveryInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

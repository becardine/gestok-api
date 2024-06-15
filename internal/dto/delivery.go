package dto

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"time"
)

type CreateDeliveryInput struct {
	OrderID        common.ID `json:"order_id"`
	CustomerID     common.ID `json:"customer_id"`
	DeliveryType   string    `json:"delivery_type"`
	DeliveryDate   time.Time `json:"delivery_date"`
	DeliveryStatus string    `json:"delivery_status"`
}

type UpdateDeliveryInput struct {
	ID             common.ID `json:"id"`
	OrderID        common.ID `json:"order_id"`
	CustomerID     common.ID `json:"customer_id"`
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

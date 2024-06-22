package dto

import (
	"encoding/json"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateOrderItemInput struct {
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
}

type UpdateOrderItemInput struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
}

func (input *CreateOrderItemInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *CreateOrderItemInput) ToEntity() *entity.OrderItem {
	return &entity.OrderItem{
		ID:        uuid.New(),
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		UnitPrice: input.UnitPrice,
	}
}

func (input *UpdateOrderItemInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateOrderItemInput) ToEntity() *entity.OrderItem {
	return &entity.OrderItem{
		ID:        input.ID,
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		UnitPrice: input.UnitPrice,
	}
}

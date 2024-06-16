package dto

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type CreateOrderItemInput struct {
	OrderID   common.ID `json:"order_id"`
	ProductID common.ID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
}

type UpdateOrderItemInput struct {
	ID        common.ID `json:"id"`
	OrderID   common.ID `json:"order_id"`
	ProductID common.ID `json:"product_id"`
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
		ID:        common.NewID(),
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

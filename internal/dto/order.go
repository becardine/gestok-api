package dto

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"time"
)

type CreateOrderInput struct {
	CustomerID  common.ID `json:"customer_id"`
	OrderDate   time.Time `json:"order_date"`
	OrderStatus string    `json:"order_status"`
	TotalValue  float64   `json:"total_value"`
}

type UpdateOrderInput struct {
	ID          common.ID `json:"id"`
	CustomerID  common.ID `json:"customer_id"`
	OrderDate   time.Time `json:"order_date"`
	OrderStatus string    `json:"order_status"`
	TotalValue  float64   `json:"total_value"`
}

func (input *CreateOrderInput) ToEntity() *entity.Order {
	return &entity.Order{
		CustomerID:  input.CustomerID,
		OrderDate:   input.OrderDate,
		OrderStatus: input.OrderStatus,
		TotalValue:  input.TotalValue,
	}
}

func (input *CreateOrderInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateOrderInput) ToEntity() *entity.Order {
	return &entity.Order{
		ID:          input.ID,
		CustomerID:  input.CustomerID,
		OrderDate:   input.OrderDate,
		OrderStatus: input.OrderStatus,
		TotalValue:  input.TotalValue,
	}
}

func (input *UpdateOrderInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

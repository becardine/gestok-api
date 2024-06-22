package dto

import (
	"encoding/json"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CreatePaymentInput struct {
	OrderID    uuid.UUID `json:"order_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Method     string    `json:"method"`
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
}

type UpdatePaymentInput struct {
	ID         uuid.UUID `json:"id"`
	OrderID    uuid.UUID `json:"order_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Method     string    `json:"method"`
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
}

func (input *CreatePaymentInput) ToEntity() *entity.Payment {
	return &entity.Payment{
		OrderID:    input.OrderID,
		CustomerID: input.CustomerID,
		Method:     input.Method,
		Date:       input.Date,
		Amount:     input.Amount,
		Status:     input.Status,
	}
}

func (input *CreatePaymentInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdatePaymentInput) ToEntity() *entity.Payment {
	return &entity.Payment{
		ID:         input.ID,
		OrderID:    input.OrderID,
		CustomerID: input.CustomerID,
		Method:     input.Method,
		Date:       input.Date,
		Amount:     input.Amount,
		Status:     input.Status,
	}
}

func (input *UpdatePaymentInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

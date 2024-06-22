package entity

import (
	"time"

	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
)

type Payment struct {
	ID         uuid.UUID `json:"id"`
	OrderID    uuid.UUID `json:"order_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Method     string    `json:"method"`
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
}

func NewPayment(orderID, customerID uuid.UUID, method string, date time.Time, amount float64, status string) (*Payment, error) {
	payment := &Payment{
		ID:         uuid.New(),
		OrderID:    orderID,
		CustomerID: customerID,
		Method:     method,
		Date:       date,
		Amount:     amount,
		Status:     status,
	}

	if err := payment.Validate(); err != nil {
		return nil, err
	}

	return payment, nil
}

func (p *Payment) Validate() error {
	if p.OrderID == uuid.Nil {
		return errors.NewEntityValidationError("order_id", "required", "")
	}

	if p.CustomerID == uuid.Nil {
		return errors.NewEntityValidationError("customer_id", "required", "")
	}

	if p.Method == "" {
		return errors.NewEntityValidationError("method", "required", "")
	}

	if p.Date.IsZero() {
		return errors.NewEntityValidationError("date", "required", "")
	}

	if p.Amount <= 0 {
		return errors.NewEntityValidationError("amount", "invalid_range", "> 0")
	}

	if p.Status == "" {
		return errors.NewEntityValidationError("status", "required", "")
	}

	return nil
}

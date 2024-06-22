package entity

import (
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
)

type Feedback struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	OrderID    uuid.UUID `json:"order_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
}

func NewFeedback(customerID, orderID uuid.UUID, rating int, comment string) (*Feedback, error) {
	feedback := &Feedback{
		ID:         uuid.New(),
		CustomerID: customerID,
		OrderID:    orderID,
		Rating:     rating,
		Comment:    comment,
	}

	if err := feedback.Validate(); err != nil {
		return nil, err
	}

	return feedback, nil
}

func (f *Feedback) Validate() error {
	if f.CustomerID == uuid.Nil {
		return errors.NewEntityValidationError("customer_id", "required", "")
	}

	if f.OrderID == uuid.Nil {
		return errors.NewEntityValidationError("order_id", "required", "")
	}

	if f.Rating < 1 || f.Rating > 5 {
		return errors.NewEntityValidationError("rating", "invalid_range", "1-5")
	}

	if f.Comment == "" {
		return errors.NewEntityValidationError("comment", "required", "")
	}

	return nil
}

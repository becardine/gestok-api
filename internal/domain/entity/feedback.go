package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
)

type Feedback struct {
	ID         common.ID `json:"id"`
	CustomerID common.ID `json:"customer_id"`
	OrderID    common.ID `json:"order_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
}

func NewFeedback(customerID, orderID common.ID, rating int, comment string) (*Feedback, error) {
	feedback := &Feedback{
		ID:         common.NewID(),
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
	if f.CustomerID.IsEmpty() {
		return errors.NewEntityValidationError("customer_id", "required", "")
	}

	if f.OrderID.IsEmpty() {
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

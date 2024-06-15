package dto

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type CreateFeedbackInput struct {
	CustomerID common.ID `json:"customer_id"`
	OrderID    common.ID `json:"order_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
}

type UpdateFeedbackInput struct {
	ID         common.ID `json:"id"`
	CustomerID common.ID `json:"customer_id"`
	OrderID    common.ID `json:"order_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
}

func (input *CreateFeedbackInput) ToEntity() *entity.Feedback {
	return &entity.Feedback{
		CustomerID: input.CustomerID,
		OrderID:    input.OrderID,
		Rating:     input.Rating,
		Comment:    input.Comment,
	}
}

func (input *CreateFeedbackInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateFeedbackInput) ToEntity() *entity.Feedback {
	return &entity.Feedback{
		ID:         input.ID,
		CustomerID: input.CustomerID,
		OrderID:    input.OrderID,
		Rating:     input.Rating,
		Comment:    input.Comment,
	}
}

func (input *UpdateFeedbackInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

package dto

import (
	"encoding/json"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateStockInput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Capacity int    `json:"capacity"`
}

type UpdateStockInput struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Capacity int       `json:"capacity"`
}

func (input *CreateStockInput) ToEntity() *entity.Stock {
	return &entity.Stock{
		Name:     input.Name,
		Location: input.Location,
		Capacity: input.Capacity,
	}
}

func (input *CreateStockInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateStockInput) ToEntity() *entity.Stock {
	return &entity.Stock{
		ID:       input.ID,
		Name:     input.Name,
		Location: input.Location,
		Capacity: input.Capacity,
	}
}

func (input *UpdateStockInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

package dto

import (
	"encoding/json"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateBrandInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateBrandInput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (input *CreateBrandInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *CreateBrandInput) ToEntity() *entity.Brand {
	return &entity.Brand{
		ID:          uuid.New(),
		Name:        input.Name,
		Description: input.Description,
	}
}

func (input *UpdateBrandInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateBrandInput) ToEntity() *entity.Brand {
	return &entity.Brand{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}
}

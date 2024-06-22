package dto

import (
	"encoding/json"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryInput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (i *CreateCategoryInput) ToEntity() *entity.Category {
	return &entity.Category{
		Name:        i.Name,
		Description: i.Description,
	}
}

func (i *UpdateCategoryInput) ToEntity() *entity.Category {
	return &entity.Category{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
	}
}

func (i *CreateCategoryInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, i); err != nil {
		return err
	}

	return nil
}

func (i *UpdateCategoryInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, i); err != nil {
		return err
	}

	return nil
}

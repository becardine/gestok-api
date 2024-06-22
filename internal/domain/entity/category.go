package entity

import (
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewCategory(name, description string) (*Category, error) {
	category := &Category{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}

	if err := category.Validate(); err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.NewEntityValidationError("name", "required", "")
	}

	if len(c.Name) > 100 {
		return errors.NewEntityValidationError("name", "max_length", "100")
	}

	if c.Description != "" && len(c.Description) > 255 {
		return errors.NewEntityValidationError("description", "max_length", "255")
	}

	return nil
}

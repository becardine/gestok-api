package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
)

type Category struct {
	ID          common.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewCategory(name, description string) (*Category, error) {
	category := &Category{
		ID:          common.NewID(),
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

package entity

import (
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
)

type Brand struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewBrand(name, description string) (*Brand, error) {
	brand := &Brand{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}

	if err := brand.Validate(); err != nil {
		return nil, err
	}

	return brand, nil
}

func (b *Brand) Validate() error {
	if b.Name == "" {
		return errors.NewEntityValidationError("name", "required", "")
	}

	if len(b.Name) > 100 {
		return errors.NewEntityValidationError("name", "max_length", "100")
	}

	if b.Description != "" && len(b.Description) > 255 {
		return errors.NewEntityValidationError("description", "max_length", "255")
	}

	return nil
}

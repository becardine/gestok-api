package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
)

type Stock struct {
	ID       common.ID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Capacity int       `json:"capacity"`
}

func NewStock(name, location string, capacity int) (*Stock, error) {
	stock := &Stock{
		ID:       common.NewID(),
		Name:     name,
		Location: location,
		Capacity: capacity,
	}

	if err := stock.Validate(); err != nil {
		return nil, err
	}

	return stock, nil
}

func (s *Stock) Validate() error {
	if s.Name == "" {
		return errors.NewEntityValidationError("name", "required", "")
	}

	if len(s.Name) > 100 {
		return errors.NewEntityValidationError("name", "max_length", "100")
	}

	if s.Location == "" {
		return errors.NewEntityValidationError("location", "required", "")
	}

	if len(s.Location) > 255 {
		return errors.NewEntityValidationError("location", "max_length", "255")
	}

	if s.Capacity < 1 {
		return errors.NewEntityValidationError("capacity", "invalid_range", "1-")
	}

	return nil
}

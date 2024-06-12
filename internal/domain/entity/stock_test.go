package entity_test

import (
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewStock(t *testing.T) {
	name := "Main Warehouse"
	location := "123 Main St"
	capacity := 1000

	stock, err := entity.NewStock(name, location, capacity)
	assert.Nil(t, err)
	assert.NotNil(t, stock)
	assert.Equal(t, name, stock.Name)
	assert.Equal(t, location, stock.Location)
	assert.Equal(t, capacity, stock.Capacity)
}

func TestStockValidate(t *testing.T) {
	testCases := []struct {
		name        string
		stock       *entity.Stock
		expectedErr error
	}{
		{
			name: "valid stock",
			stock: &entity.Stock{
				ID:       common.NewID(),
				Name:     "Main Warehouse",
				Location: "123 Main St",
				Capacity: 1000,
			},
			expectedErr: nil,
		},
		{
			name: "empty name",
			stock: &entity.Stock{
				ID:       common.NewID(),
				Location: "123 Main St",
				Capacity: 1000,
			},
			expectedErr: errors.NewEntityValidationError("name", "required", ""),
		},
		{
			name: "name exceeding max length",
			stock: &entity.Stock{
				ID:       common.NewID(),
				Name:     "Main Warehouse Main Warehouse Main Warehouse Main Warehouse Main Warehouse Main Warehouse Main Warehouse",
				Location: "123 Main St",
				Capacity: 1000,
			},
			expectedErr: errors.NewEntityValidationError("name", "max_length", "100"),
		},
		{
			name: "empty location",
			stock: &entity.Stock{
				ID:       common.NewID(),
				Name:     "Main Warehouse",
				Capacity: 1000,
			},
			expectedErr: errors.NewEntityValidationError("location", "required", ""),
		},
		{
			name: "location exceeding max length",
			stock: &entity.Stock{
				ID:       common.NewID(),
				Name:     "Main Warehouse",
				Location: "123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St 123 Main St Main St 123 Main St 123 Main St 123 Main StMain St 123 Main St 123 Main St 123 Main StMain St 123 Main St 123 Main St 123 Main St",
				Capacity: 1000,
			},
			expectedErr: errors.NewEntityValidationError("location", "max_length", "255"),
		},
		{
			name: "invalid capacity",
			stock: &entity.Stock{
				ID:       common.NewID(),
				Name:     "Main Warehouse",
				Location: "123 Main St",
				Capacity: 0,
			},
			expectedErr: errors.NewEntityValidationError("capacity", "invalid_range", "1-"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.stock.Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

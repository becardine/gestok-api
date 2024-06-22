package entity_test

import (
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderItem(t *testing.T) {
	t.Run("Should create a new order item", func(t *testing.T) {
		orderItem, err := entity.NewOrderItem(uuid.New(), uuid.New(), 1, 10.0)
		assert.NoError(t, err)
		assert.NotNil(t, orderItem)
		assert.Equal(t, 1, orderItem.Quantity)
		assert.Equal(t, 10.0, orderItem.UnitPrice)
	})

	t.Run("Should return an error when order id is empty", func(t *testing.T) {
		orderItem, err := entity.NewOrderItem(uuid.UUID{}, uuid.New(), 1, 10.0)
		assert.Error(t, err)
		assert.Nil(t, orderItem)
	})

	t.Run("Should return an error when product id is empty", func(t *testing.T) {
		orderItem, err := entity.NewOrderItem(uuid.New(), uuid.UUID{}, 1, 10.0)
		assert.Error(t, err)
		assert.Nil(t, orderItem)
	})

	t.Run("Should return an error when quantity is less than or equal to 0", func(t *testing.T) {
		orderItem, err := entity.NewOrderItem(uuid.New(), uuid.New(), 0, 10.0)
		assert.Error(t, err)
		assert.Nil(t, orderItem)
	})

	t.Run("Should return an error when unit price is less than or equal to 0", func(t *testing.T) {
		orderItem, err := entity.NewOrderItem(uuid.New(), uuid.New(), 1, 0)
		assert.Error(t, err)
		assert.Nil(t, orderItem)
	})

	t.Run("Should return an error when quantity and unit price are less than or equal to 0", func(t *testing.T) {
		orderItem, err := entity.NewOrderItem(uuid.New(), uuid.New(), 0, 0)
		assert.Error(t, err)
		assert.Nil(t, orderItem)
	})

	t.Run("Should return an correct total price", func(t *testing.T) {
		orderItem, _ := entity.NewOrderItem(uuid.New(), uuid.New(), 2, 10.0)
		assert.Equal(t, 20.0, orderItem.TotalPrice())
	})

	t.Run("Should return an error when total price is incorrect", func(t *testing.T) {
		orderItem, _ := entity.NewOrderItem(uuid.New(), uuid.New(), 2, 10.0)
		assert.NotEqual(t, 30.0, orderItem.TotalPrice())
	})
}

package entity_test

import (
	"testing"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewDelivery(t *testing.T) {
	orderID := uuid.New()
	customerID := uuid.New()
	deliveryType := "Correios"
	deliveryDate := time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC)
	deliveryStatus := "In transit"

	delivery, err := entity.NewDelivery(orderID, customerID, deliveryType, deliveryStatus, deliveryDate)

	assert.NoError(t, err)
	assert.NotNil(t, delivery)
	assert.Equal(t, orderID, delivery.OrderID)
	assert.Equal(t, customerID, delivery.CustomerID)
	assert.Equal(t, deliveryType, delivery.DeliveryType)
	assert.Equal(t, deliveryDate, delivery.DeliveryDate)
	assert.Equal(t, deliveryStatus, delivery.DeliveryStatus)
}

func TestDeliveryValidate(t *testing.T) {
	orderID := uuid.New()
	customerID := uuid.New()

	deliveryDate := time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name        string
		delivery    *entity.Delivery
		expectedErr error
	}{
		{
			name: "valid delivery",
			delivery: &entity.Delivery{
				ID:             uuid.New(),
				OrderID:        orderID,
				CustomerID:     customerID,
				DeliveryType:   "Correios",
				DeliveryDate:   deliveryDate,
				DeliveryStatus: "In transit",
			},
			expectedErr: nil,
		},
		{
			name: "empty order ID",
			delivery: &entity.Delivery{
				ID:             uuid.New(),
				CustomerID:     customerID,
				DeliveryType:   "Correios",
				DeliveryDate:   deliveryDate,
				DeliveryStatus: "In transit",
			},
			expectedErr: errors.NewEntityValidationError("order_id", "required", ""),
		},
		{
			name: "empty customer ID",
			delivery: &entity.Delivery{
				ID:             uuid.New(),
				OrderID:        orderID,
				DeliveryType:   "Correios",
				DeliveryDate:   deliveryDate,
				DeliveryStatus: "In transit",
			},
			expectedErr: errors.NewEntityValidationError("customer_id", "required", ""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.delivery.Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

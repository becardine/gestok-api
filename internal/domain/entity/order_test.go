package entity_test

import (
	"testing"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	customerID := common.NewID()
	orderDate := time.Now()
	orderStatus := "Pending"
	totalValue := 100.50

	order, err := entity.NewOrder(customerID, orderDate, orderStatus, totalValue)

	assert.Nil(t, err)

	assert.NotNil(t, order)
	assert.Equal(t, customerID, order.CustomerID)
	assert.Equal(t, orderDate, order.OrderDate)
	assert.Equal(t, orderStatus, order.OrderStatus)
	assert.Equal(t, totalValue, order.TotalValue)
}

func TestOrderValidate(t *testing.T) {
	customerID := common.NewID()
	orderDate := time.Now()

	testCases := []struct {
		name        string
		order       *entity.Order
		expectedErr error
	}{
		{
			name: "valid order",
			order: &entity.Order{
				ID:          common.NewID(),
				CustomerID:  customerID,
				OrderDate:   orderDate,
				OrderStatus: "Pending",
				TotalValue:  100.50,
			},
			expectedErr: nil,
		},
		{
			name: "empty customer ID",
			order: &entity.Order{
				ID:          common.NewID(),
				OrderDate:   orderDate,
				OrderStatus: "Pending",
				TotalValue:  100.50,
			},
			expectedErr: errors.NewEntityValidationError("customer_id", "required", ""),
		},
		{
			name: "empty order date",
			order: &entity.Order{
				ID:          common.NewID(),
				CustomerID:  customerID,
				OrderStatus: "Pending",
				TotalValue:  100.50,
			},
			expectedErr: errors.NewEntityValidationError("order_date", "required", ""),
		},
		{
			name: "empty order status",
			order: &entity.Order{
				ID:         common.NewID(),
				CustomerID: customerID,
				OrderDate:  orderDate,
				TotalValue: 100.50,
			},
			expectedErr: errors.NewEntityValidationError("order_status", "required", ""),
		},
		{
			name: "invalid total value",
			order: &entity.Order{
				ID:          common.NewID(),
				CustomerID:  customerID,
				OrderDate:   orderDate,
				OrderStatus: "Pending",
				TotalValue:  0,
			},
			expectedErr: errors.NewEntityValidationError("total_value", "invalid_range", "> 0"),
		},
		// ... add more test cases for other fields and validations
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.order.Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

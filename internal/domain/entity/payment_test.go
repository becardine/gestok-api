package entity_test

import (
	"testing"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPayment(t *testing.T) {
	orderID := uuid.New()
	customerID := uuid.New()
	method := "Credit Card"
	date := time.Now()
	amount := 100.50
	status := "Completed"

	payment, err := entity.NewPayment(orderID, customerID, method, date, amount, status)
	assert.Nil(t, err)

	assert.NotNil(t, payment)
	assert.Equal(t, orderID, payment.OrderID)
	assert.Equal(t, customerID, payment.CustomerID)
	assert.Equal(t, method, payment.Method)
	assert.Equal(t, date, payment.Date)
	assert.Equal(t, amount, payment.Amount)
	assert.Equal(t, status, payment.Status)
}

func TestPaymentValidate(t *testing.T) {
	orderID := uuid.New()
	customerID := uuid.New()
	date := time.Now()

	testCases := []struct {
		name        string
		payment     *entity.Payment
		expectedErr error
	}{
		{
			name: "valid payment",
			payment: &entity.Payment{
				ID:         uuid.New(),
				OrderID:    orderID,
				CustomerID: customerID,
				Method:     "Credit Card",
				Date:       date,
				Amount:     100.50,
				Status:     "Completed",
			},
			expectedErr: nil,
		},
		{
			name: "empty order ID",
			payment: &entity.Payment{
				ID:         uuid.New(),
				CustomerID: customerID,
				Method:     "Credit Card",
				Date:       date,
				Amount:     100.50,
				Status:     "Completed",
			},
			expectedErr: errors.NewEntityValidationError("order_id", "required", ""),
		},
		{
			name: "empty customer ID",
			payment: &entity.Payment{
				ID:      uuid.New(),
				OrderID: orderID,
				Method:  "Credit Card",
				Date:    date,
				Amount:  100.50,
				Status:  "Completed",
			},
			expectedErr: errors.NewEntityValidationError("customer_id", "required", ""),
		},
		{
			name: "empty method",
			payment: &entity.Payment{
				ID:         uuid.New(),
				OrderID:    orderID,
				CustomerID: customerID,
				Date:       date,
				Amount:     100.50,
				Status:     "Completed",
			},
			expectedErr: errors.NewEntityValidationError("method", "required", ""),
		},
		{
			name: "empty date",
			payment: &entity.Payment{
				ID:         uuid.New(),
				OrderID:    orderID,
				CustomerID: customerID,
				Method:     "Credit Card",
				Amount:     100.50,
				Status:     "Completed",
			},
			expectedErr: errors.NewEntityValidationError("date", "required", ""),
		},
		{
			name: "invalid amount",
			payment: &entity.Payment{
				ID:         uuid.New(),
				OrderID:    orderID,
				CustomerID: customerID,
				Method:     "Credit Card",
				Date:       date,
				Amount:     0,
				Status:     "Completed",
			},
			expectedErr: errors.NewEntityValidationError("amount", "invalid_range", "> 0"),
		},
		{
			name: "empty status",
			payment: &entity.Payment{
				ID:         uuid.New(),
				OrderID:    orderID,
				CustomerID: customerID,
				Method:     "Credit Card",
				Date:       date,
				Amount:     100.50,
			},
			expectedErr: errors.NewEntityValidationError("status", "required", ""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.payment.Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

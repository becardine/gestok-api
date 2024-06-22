package entity_test

import (
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewFeedback(t *testing.T) {
	customerID := uuid.New()
	orderID := uuid.New()
	rating := 4
	comment := "Great service!"

	feedback, err := entity.NewFeedback(customerID, orderID, rating, comment)
	assert.Nil(t, err)

	assert.NotNil(t, feedback)
	assert.Equal(t, customerID, feedback.CustomerID)
	assert.Equal(t, orderID, feedback.OrderID)
	assert.Equal(t, rating, feedback.Rating)
	assert.Equal(t, comment, feedback.Comment)
}

func TestFeedbackValidate(t *testing.T) {
	customerID := uuid.New()
	orderID := uuid.New()

	testCases := []struct {
		name        string
		feedback    *entity.Feedback
		expectedErr error
	}{
		{
			name: "valid feedback",
			feedback: &entity.Feedback{
				ID:         uuid.New(),
				CustomerID: customerID,
				OrderID:    orderID,
				Rating:     4,
				Comment:    "Great service!",
			},
			expectedErr: nil,
		},
		{
			name: "empty customer ID",
			feedback: &entity.Feedback{
				ID:      uuid.New(),
				OrderID: orderID,
				Rating:  4,
				Comment: "Great service!",
			},
			expectedErr: errors.NewEntityValidationError("customer_id", "required", ""),
		},
		{
			name: "empty order ID",
			feedback: &entity.Feedback{
				ID:         uuid.New(),
				CustomerID: customerID,
				Rating:     4,
				Comment:    "Great service!",
			},
			expectedErr: errors.NewEntityValidationError("order_id", "required", ""),
		},
		{
			name: "invalid rating (below minimum)",
			feedback: &entity.Feedback{
				ID:         uuid.New(),
				CustomerID: customerID,
				OrderID:    orderID,
				Rating:     0,
				Comment:    "Great service!",
			},
			expectedErr: errors.NewEntityValidationError("rating", "invalid_range", "1-5"),
		},
		{
			name: "invalid rating (above maximum)",
			feedback: &entity.Feedback{
				ID:         uuid.New(),
				CustomerID: customerID,
				OrderID:    orderID,
				Rating:     6,
				Comment:    "Great service!",
			},
			expectedErr: errors.NewEntityValidationError("rating", "invalid_range", "1-5"),
		},
		{
			name: "empty comment",
			feedback: &entity.Feedback{
				ID:         uuid.New(),
				CustomerID: customerID,
				OrderID:    orderID,
				Rating:     4,
			},
			expectedErr: errors.NewEntityValidationError("comment", "required", ""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.feedback.Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

package entity_test

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	customer, err := entity.NewCustomer("John Doe", "john.doe@example.com", "password123", "123 Main St", "555-123-4567")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.NotEmpty(t, customer.ID)
	assert.Equal(t, "John Doe", customer.Name)
	assert.Equal(t, "john.doe@example.com", customer.Email)
	assert.NotEqual(t, "password123", customer.Password) // password should be hashed
	assert.Equal(t, "123 Main St", customer.Address)
	assert.Equal(t, "555-123-4567", customer.Phone)
}

func TestCustomer_ValidatePassword(t *testing.T) {
	customer, err := entity.NewCustomer("John Doe", "john.doe@example.com", "password123", "123 Main St", "555-123-4567")
	assert.NoError(t, err)

	err = customer.ValidatePassword("password123")
	assert.NoError(t, err)

	err = customer.ValidatePassword("wrongpassword")
	assert.Error(t, err)
}

func TestCustomer_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		customer    *entity.Customer
		expectedErr error
	}{
		{
			name: "valid customer",
			customer: &entity.Customer{
				ID:       common.NewID(),
				Name:     "John Doe",
				Email:    "john.doe@example.com",
				Password: "password123",
				Address:  "123 Main St",
				Phone:    "555-123-4567",
			},
			expectedErr: nil,
		},
		{
			name: "empty name",
			customer: &entity.Customer{
				ID:       common.NewID(),
				Email:    "john.doe@example.com",
				Password: "password123",
				Address:  "123 Main St",
				Phone:    "555-123-4567",
			},
			expectedErr: errors.NewEntityValidationError("name", "required", ""),
		},
		{
			name: "name exceeding max length",
			customer: &entity.Customer{
				ID:       common.NewID(),
				Name:     "John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe",
				Email:    "john.doe@example.com",
				Password: "password123",
				Address:  "123 Main St",
				Phone:    "555-123-4567",
			},
			expectedErr: errors.NewEntityValidationError("name", "max_length", "100"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.customer.Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestCustomer_Validate_Email(t *testing.T) {
	testCases := []struct {
		name        string
		email       string
		expectedErr error
	}{
		{
			name:        "valid email",
			email:       "john.doe@example.com",
			expectedErr: nil,
		},
		{
			name:        "empty email",
			email:       "",
			expectedErr: errors.NewEntityValidationError("email", "required", ""),
		},
		{
			name:        "email exceeding max length",
			email:       "john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe.john.doe@example.com",
			expectedErr: errors.NewEntityValidationError("email", "max_length", "100"),
		},
		{
			name:        "invalid email format",
			email:       "john.doe",
			expectedErr: errors.NewEntityValidationError("email", "invalid_format", ""),
		},
		{
			name:        "invalid email format (missing @)",
			email:       "john.doe.example.com",
			expectedErr: errors.NewEntityValidationError("email", "invalid_format", ""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customer := &entity.Customer{
				ID:       common.NewID(),
				Name:     "John Doe",
				Email:    tc.email,
				Password: "password123",
				Address:  "123 Main St",
				Phone:    "555-123-4567",
			}
			err := customer.Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

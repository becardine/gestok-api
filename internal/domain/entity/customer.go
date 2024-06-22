package entity

import (
	"net/mail"

	"github.com/becardine/gestock-api/internal/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // - means that this field will not be serialized
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
}

func NewCustomer(name, email, password, address, phone string) (*Customer, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Customer{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Address:  address,
		Phone:    phone,
	}, nil
}

func (c *Customer) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password),
		[]byte(password))
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return errors.NewEntityValidationError("name", "required", "")
	}

	if len(c.Name) > 100 {
		return errors.NewEntityValidationError("name", "max_length", "100")
	}

	if c.Email == "" {
		return errors.NewEntityValidationError("email", "required", "")
	}

	if len(c.Email) > 100 {
		return errors.NewEntityValidationError("email", "max_length", "100")
	}

	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.NewEntityValidationError("email", "invalid_format", "")
	}

	if c.Address == "" {
		return errors.NewEntityValidationError("address", "required", "")
	}

	if len(c.Address) > 255 {
		return errors.NewEntityValidationError("address", "max_length", "255")
	}

	if c.Phone == "" {
		return errors.NewEntityValidationError("phone", "required", "")
	}

	if len(c.Phone) > 20 {
		return errors.NewEntityValidationError("phone", "max_length", "20")
	}

	return nil
}

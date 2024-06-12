package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       common.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // - means that this field will not be serialized
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       common.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

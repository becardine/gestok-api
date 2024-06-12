package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       common.ID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"uniqueIndex;not null"`
	Password string    `json:"-" gorm:"not null"`
	gorm.Model
}

func NewUser(name, email, password string) (*User, error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       common.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hasedPassword),
	}, nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createNewUser() (*User, error) {
	user, err := NewUser("John Doe", "john@email.com", "password")
	return user, err
}

func TestNewUser(t *testing.T) {
	user, err := createNewUser()
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@email.com", user.Email)

}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := createNewUser()
	assert.Nil(t, err)
	assert.Nil(t, user.ValidatePassword("password"))
	assert.NotNil(t, user.ValidatePassword("wrong_password"))
	assert.NotEqual(t, "password", user.Password)
}

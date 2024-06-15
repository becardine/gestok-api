package dto

import (
	"encoding/json"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type CreateCustomerInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UpdateCustomerInput struct {
	ID       common.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
}

func (input *CreateCustomerInput) ToEntity() *entity.Customer {
	return &entity.Customer{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Address:  input.Address,
		Phone:    input.Phone,
	}
}

func (input *CreateCustomerInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateCustomerInput) ToEntity() *entity.Customer {
	return &entity.Customer{
		ID:       input.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Address:  input.Address,
		Phone:    input.Phone,
	}
}

func (input *UpdateCustomerInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

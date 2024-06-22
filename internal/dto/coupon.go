package dto

import (
	"encoding/json"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateCouponInput struct {
	Code           string    `json:"code"`
	Discount       float64   `json:"discount"`
	ExpirationDate time.Time `json:"expiration_at"`
	Status         string    `json:"status"`
}

type UpdateCouponInput struct {
	ID             uuid.UUID `json:"id"`
	Code           string    `json:"code"`
	Discount       float64   `json:"discount"`
	ExpirationDate time.Time `json:"expiration_at"`
	Status         string    `json:"status"`
}

func (input *CreateCouponInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *CreateCouponInput) ToEntity() *entity.Coupon {
	return &entity.Coupon{
		Code:           input.Code,
		Discount:       input.Discount,
		ExpirationDate: input.ExpirationDate,
		Status:         input.Status,
	}
}

func (input *UpdateCouponInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *UpdateCouponInput) ToEntity() *entity.Coupon {
	return &entity.Coupon{
		ID:             input.ID,
		Code:           input.Code,
		Discount:       input.Discount,
		ExpirationDate: input.ExpirationDate,
		Status:         input.Status,
	}
}

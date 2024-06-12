package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
	"time"
)

type Coupon struct {
	ID             common.ID `json:"id"`
	Code           string    `json:"code"`
	Discount       float64   `json:"discount"`
	ExpirationDate time.Time `json:"expiration_date"`
	Status         string    `json:"status"`
}

func NewCoupon(code string, discount float64, expirationDate time.Time, status string) (*Coupon, error) {
	coupon := &Coupon{
		ID:             common.NewID(),
		Code:           code,
		Discount:       discount,
		ExpirationDate: expirationDate,
		Status:         status,
	}

	if err := coupon.Validate(); err != nil {
		return nil, err
	}

	return coupon, nil
}

func (c *Coupon) Validate() error {
	if c.Code == "" {
		return errors.NewEntityValidationError("code", "required", "")
	}

	if len(c.Code) > 100 {
		return errors.NewEntityValidationError("code", "max_length", "100")
	}

	if c.Discount <= 0 {
		return errors.NewEntityValidationError("discount", "greater_than", "0")
	}

	if c.ExpirationDate.IsZero() {
		return errors.NewEntityValidationError("expiration_date", "required", "")
	}

	if c.Status == "" {
		return errors.NewEntityValidationError("status", "required", "")
	}

	return nil
}

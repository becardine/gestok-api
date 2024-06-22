package entity_test

import (
	"testing"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCoupon(t *testing.T) {
	t.Run("should create a new coupon with valid data", func(t *testing.T) {
		expirationDate := time.Now().AddDate(0, 1, 0)
		coupon, err := entity.NewCoupon("SUMMER20", 10.0, expirationDate, "active")

		assert.NoError(t, err)
		assert.NotNil(t, coupon)
		assert.Equal(t, "SUMMER20", coupon.Code)
		assert.Equal(t, 10.0, coupon.Discount)
		assert.Equal(t, expirationDate, coupon.ExpirationDate)
		assert.Equal(t, "active", coupon.Status)
	})

	t.Run("should return error if code is empty", func(t *testing.T) {
		expirationDate := time.Now().AddDate(0, 1, 0)
		coupon, err := entity.NewCoupon("", 10.0, expirationDate, "active")

		assert.Error(t, err)
		assert.Nil(t, coupon)
	})

	t.Run("should return error if code exceeds max length", func(t *testing.T) {
		expirationDate := time.Now().AddDate(0, 1, 0)
		longCode := "ThisIsAVeryLongCouponCodeThatExceedsTheMaximumLengthOf100CharactersAndContinuesToBeEvenLongerSoThatItDefinitelyExceedsTheLimit"
		coupon, err := entity.NewCoupon(longCode, 10.0, expirationDate, "active")

		assert.Error(t, err)
		assert.Nil(t, coupon)
	})

	t.Run("should return error if discount is less than or equal to zero", func(t *testing.T) {
		expirationDate := time.Now().AddDate(0, 1, 0)
		coupon, err := entity.NewCoupon("SUMMER20", 0.0, expirationDate, "active")

		assert.Error(t, err)
		assert.Nil(t, coupon)
	})

	t.Run("should return error if expiration date is zero", func(t *testing.T) {
		var expirationDate time.Time
		coupon, err := entity.NewCoupon("SUMMER20", 10.0, expirationDate, "active")

		assert.Error(t, err)
		assert.Nil(t, coupon)
	})

	t.Run("should return error if status is empty", func(t *testing.T) {
		expirationDate := time.Now().AddDate(0, 1, 0)
		coupon, err := entity.NewCoupon("SUMMER20", 10.0, expirationDate, "")

		assert.Error(t, err)
		assert.Nil(t, coupon)
	})
}

func TestCoupon_Validate(t *testing.T) {
	t.Run("should return no error if coupon is valid", func(t *testing.T) {
		expirationDate := time.Now().AddDate(0, 1, 0)
		coupon := &entity.Coupon{
			ID:             uuid.New(),
			Code:           "SUMMER20",
			Discount:       10.0,
			ExpirationDate: expirationDate,
			Status:         "active",
		}

		err := coupon.Validate()
		assert.NoError(t, err)
	})
}

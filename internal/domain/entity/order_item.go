package entity

import (
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	"github.com/becardine/gestock-api/internal/errors"
)

type OrderItem struct {
	ID        common.ID `json:"id"`
	OrderID   common.ID `json:"order_id"`
	ProductID common.ID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
}

func NewOrderItem(orderID, productID common.ID, quantity int, unitPrice float64) (*OrderItem, error) {
	orderItem := &OrderItem{
		ID:        common.NewID(),
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		UnitPrice: unitPrice,
	}

	if err := orderItem.Validate(); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (o *OrderItem) TotalPrice() float64 {
	return float64(o.Quantity) * o.UnitPrice
}

func (o *OrderItem) Validate() error {
	if o.OrderID.IsEmpty() {
		return errors.NewEntityValidationError("order id", "required", "")
	}

	if o.ProductID.IsEmpty() {
		return errors.NewEntityValidationError("product id", "required", "")
	}

	if o.Quantity <= 0 {
		return errors.NewEntityValidationError("quantity", "must be greater than 0", "")
	}

	if o.UnitPrice <= 0 {
		return errors.NewEntityValidationError("unit price", "must be greater than 0", "")
	}

	return nil
}

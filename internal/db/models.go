// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Brand struct {
	ID          int32
	Name        string
	Description sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

type Category struct {
	ID          int32
	Name        string
	Description sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

type Coupon struct {
	ID             int32
	Code           string
	Discount       string
	ExpirationDate time.Time
	Status         string
	CreatedDate    sql.NullTime
	UpdatedDate    sql.NullTime
}

type Customer struct {
	ID          int32
	Name        string
	Email       string
	Password    string
	Address     sql.NullString
	Phone       sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

type Delivery struct {
	ID             int32
	OrderID        sql.NullInt32
	CustomerID     sql.NullInt32
	DeliveryType   string
	DeliveryDate   sql.NullTime
	DeliveryStatus string
	CreatedDate    sql.NullTime
	UpdatedDate    sql.NullTime
}

type Feedback struct {
	ID          int32
	CustomerID  sql.NullInt32
	OrderID     sql.NullInt32
	Rating      sql.NullInt32
	Comment     sql.NullString
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

type Order struct {
	ID          int32
	CustomerID  sql.NullInt32
	OrderDate   sql.NullTime
	OrderStatus string
	TotalValue  string
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}

type OrderProduct struct {
	ID        int32
	OrderID   sql.NullInt32
	ProductID sql.NullInt32
	Quantity  int32
	UnitPrice string
}

type Payment struct {
	ID            int32
	OrderID       sql.NullInt32
	CustomerID    sql.NullInt32
	PaymentType   string
	PaymentDate   sql.NullTime
	PaymentValue  string
	PaymentStatus string
	CreatedDate   sql.NullTime
	UpdatedDate   sql.NullTime
}

type Product struct {
	ID              int32
	Name            string
	Description     sql.NullString
	Price           string
	QuantityInStock int32
	ImageUrl        sql.NullString
	CategoryID      sql.NullInt32
	BrandID         sql.NullInt32
	CreatedDate     sql.NullTime
	UpdatedDate     sql.NullTime
}

type Stock struct {
	ID          int32
	Name        string
	Location    sql.NullString
	Capacity    int32
	CreatedDate sql.NullTime
	UpdatedDate sql.NullTime
}
package entity

import (
	"fmt"
	"regexp"

	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type Product struct {
	ID              common.ID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name            string    `json:"name" gorm:"not null"`
	Description     string    `json:"description"`
	Price           float64   `json:"price" gorm:"not null"`
	QuantityInStock int       `json:"quantity_in_stock" gorm:"not null"`
	ImageURL        string    `json:"image_url"`
	CategoryID      common.ID `json:"category_id" gorm:"not null;type:uuid"`
	BrandID         common.ID `json:"brand_id" gorm:"not null;type:uuid"`
}

func NewProduct(name string, description string, price float64, quantityInStock int, imageURL string, categoryID, brandID common.ID) (*Product, error) {
	product := &Product{
		ID:              common.NewID(),
		Name:            name,
		Description:     description,
		Price:           price,
		QuantityInStock: quantityInStock,
		ImageURL:        imageURL,
		CategoryID:      categoryID,
		BrandID:         brandID,
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}
func (p *Product) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("product name is required")
	}

	if len(p.Name) > 100 {
		return fmt.Errorf("product name must have a maximum of 100 characters")
	}

	if p.Price <= 0 {
		return fmt.Errorf("product price must be greater than zero")
	}

	if p.QuantityInStock < 0 {
		return fmt.Errorf("quantity in stock must be greater than or equal to zero")
	}

	if p.ImageURL != "" {
		if !isValidURL(p.ImageURL) {
			return fmt.Errorf("invalid image URL")
		}
	}

	if p.CategoryID.IsEmpty() {
		return fmt.Errorf("category ID is required")
	}

	if p.BrandID.IsEmpty() {
		return fmt.Errorf("brand ID is required")
	}

	return nil
}

func isValidURL(imageURL string) bool {
	regex := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	return regexp.MustCompile(regex).MatchString(imageURL)
}

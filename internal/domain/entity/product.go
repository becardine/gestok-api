package entity

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

type Product struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	QuantityInStock int       `json:"quantity_in_stock"`
	ImageURL        string    `json:"image_url"`
	CategoryID      uuid.UUID `json:"category_id"`
	BrandID         uuid.UUID `json:"brand_id"`
}

func NewProduct(name string, description string, price float64, quantityInStock int, imageURL string, categoryID, brandID uuid.UUID) (*Product, error) {
	product := &Product{
		ID:              uuid.New(),
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

	if p.CategoryID == uuid.Nil {
		return fmt.Errorf("category ID is required")
	}

	if p.BrandID == uuid.Nil {
		return fmt.Errorf("brand ID is required")
	}

	return nil
}

func isValidURL(imageURL string) bool {
	regex := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	return regexp.MustCompile(regex).MatchString(imageURL)
}

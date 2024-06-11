package service

import (
	"context"
	"encoding/json"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type ProductService interface {
	GetProduct(ctx context.Context, id common.ID) (*entity.Product, error)
	CreateProduct(ctx context.Context, input *CreateProductInput) (*entity.Product, error)
	UpdateProduct(ctx context.Context, id common.ID, input *UpdateProductInput) error
	DeleteProduct(ctx context.Context, id common.ID) error
	ListProducts(ctx context.Context) ([]*entity.Product, error)
}

type CreateProductInput struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	QuantityInStock int       `json:"quantity_in_stock"`
	ImageURL        string    `json:"image_url"`
	CategoryID      common.ID `json:"category_id"`
	BrandID         common.ID `json:"brand_id"`
}

type UpdateProductInput struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	QuantityInStock int       `json:"quantity_in_stock"`
	ImageURL        string    `json:"image_url"`
	CategoryID      common.ID `json:"category_id"`
	BrandID         common.ID `json:"brand_id"`
}

func (input *CreateProductInput) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, input); err != nil {
		return err
	}

	return nil
}

func (input *CreateProductInput) ToEntity() *entity.Product {
	return &entity.Product{
		ID:              common.NewID(),
		Name:            input.Name,
		Description:     input.Description,
		Price:           input.Price,
		QuantityInStock: input.QuantityInStock,
		ImageURL:        input.ImageURL,
		CategoryID:      input.CategoryID,
		BrandID:         input.BrandID,
	}
}

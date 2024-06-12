package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
)

type productService struct {
	productRepository domain.ProductRepository
}

func NewProductService(productRepository domain.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (ps *productService) CreateProduct(ctx context.Context, input *CreateProductInput) (*entity.Product, error) {

	product := input.ToEntity()

	if err := product.Validate(); err != nil {
		return nil, fmt.Errorf("error while validating product entity: %w", err)
	}

	err := ps.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("error while creating product in repository: %w", err)
	}

	return product, nil
}

func (ps *productService) GetProduct(ctx context.Context, id common.ID) (*entity.Product, error) {
	product, err := ps.productRepository.GetProduct(ctx, common.NewIDFromUUID(id.Value()))
	if err != nil {
		return nil, ps.handleProductNotFoundError(err, id)
	}

	return product, nil
}

func (ps *productService) UpdateProduct(ctx context.Context, id common.ID, input *UpdateProductInput) error {
	product, err := ps.productRepository.GetProduct(ctx, common.NewIDFromUUID(id.Value()))
	if err != nil {
		return ps.handleProductNotFoundError(err, id)
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.QuantityInStock = input.QuantityInStock
	product.ImageURL = input.ImageURL
	product.CategoryID = input.CategoryID
	product.BrandID = input.BrandID

	err = ps.productRepository.UpdateProduct(ctx, product)
	if err != nil {
		return fmt.Errorf("error while updating product in repository: %w", err)
	}

	return nil
}

func (ps *productService) DeleteProduct(ctx context.Context, id common.ID) error {
	err := ps.productRepository.DeleteProduct(ctx, common.NewIDFromUUID(id.Value()))
	if err != nil {
		return fmt.Errorf("error while deleting product in repository: %w", err)
	}

	return nil
}

func (ps *productService) ListProducts(ctx context.Context) ([]*entity.Product, error) {
	products, err := ps.productRepository.ListProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching products from repository: %w", err)
	}

	return products, nil
}

func (ps *productService) handleProductNotFoundError(err error, id common.ID) error {
	if errors.Is(err, &domain.ErrProductNotFound{}) {
		return &domain.ErrProductNotFound{ProductID: id.String()}
	}
	return fmt.Errorf("error while fetching product from repository: %w", err)
}

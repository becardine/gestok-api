package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type CustomerServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Customer, error)
	Create(ctx context.Context, input *dto.CreateCustomerInput) (*entity.Customer, error)
	Update(ctx context.Context, input *dto.UpdateCustomerInput) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Customer, error)
	GetCustomerOrders(ctx context.Context, customerID uuid.UUID) ([]*entity.Order, error)
}

type customerService struct {
	repo domain.CustomerRepositoryInterface
}

func NewCustomerService(repo domain.CustomerRepositoryInterface) CustomerServiceInterface {
	return &customerService{
		repo: repo,
	}
}

// Create implements CustomerServiceInterface.
func (c *customerService) Create(ctx context.Context, input *dto.CreateCustomerInput) (*entity.Customer, error) {
	customer := input.ToEntity()
	if err := customer.Validate(); err != nil {
		return nil, fmt.Errorf("error validating customer: %w", err)
	}

	customer, err := c.repo.Create(ctx, customer)
	if err != nil {
		return nil, fmt.Errorf("error creating customer: %w", err)
	}

	return customer, nil
}

// Delete implements CustomerServiceInterface.
func (c *customerService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := c.repo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("error getting customer: %w", err)
	}

	if err := c.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting customer: %w", err)
	}

	return nil
}

// Get implements CustomerServiceInterface.
func (c *customerService) Get(ctx context.Context, id uuid.UUID) (*entity.Customer, error) {
	customer, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, c.handleCustomerError(err, id)
	}

	return customer, nil
}

// GetCustomerOrders implements CustomerServiceInterface.
func (c *customerService) GetCustomerOrders(ctx context.Context, customerID uuid.UUID) ([]*entity.Order, error) {
	_, err := c.repo.Get(ctx, customerID)
	if err != nil {
		return nil, c.handleCustomerError(err, customerID)
	}

	orders, err := c.repo.GetCustomerOrders(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching customer orders: %w", err)
	}

	return orders, nil
}

// List implements CustomerServiceInterface.
func (c *customerService) List(ctx context.Context, page int, pageSize int) ([]*entity.Customer, error) {
	customers, err := c.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("error fetching customers: %w", err)
	}

	return customers, nil
}

// Update implements CustomerServiceInterface.
func (c *customerService) Update(ctx context.Context, input *dto.UpdateCustomerInput) error {
	customer, err := c.repo.Get(ctx, input.ID)
	if err != nil {
		return c.handleCustomerError(err, input.ID)
	}

	customer.Name = input.Name
	customer.Email = input.Email
	customer.Phone = input.Phone

	if err := customer.Validate(); err != nil {
		return fmt.Errorf("error validating customer: %w", err)
	}

	if err := c.repo.Update(ctx, customer); err != nil {
		return fmt.Errorf("error updating customer: %w", err)
	}

	return nil
}

func (c *customerService) handleCustomerError(err error, id uuid.UUID) error {
	if errors.Is(err, &ErrNotFound{}) {
		return &ErrNotFound{Entity: "Customer", ID: id}
	}

	return fmt.Errorf("error while fetching customer: %w", err)
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	database "github.com/becardine/gestock-api/internal/infra/sqlc"
	"github.com/google/uuid"
)

type CustomerRepository struct {
	queries *database.Queries
}

func NewCustomerRepository(db *sql.DB) domain.CustomerRepositoryInterface {
	return &CustomerRepository{
		queries: database.New(db),
	}
}

// AddCustomerOrder implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) AddCustomerOrder(ctx context.Context, orderID uuid.UUID, customerID uuid.UUID) error {
	panic("unimplemented")
}

// Create implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {
	err := c.queries.CreateCustomer(ctx, database.CreateCustomerParams{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		Phone:     sql.NullString{String: customer.Phone, Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// Delete implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := c.queries.DeleteCustomer(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting customer: %w", err)
	}

	return nil
}

// Get implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Customer, error) {
	customer, err := c.queries.GetCustomer(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting customer: %w", err)
	}

	return &entity.Customer{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
		Phone: customer.Phone.String,
	}, nil
}

// GetCustomerOrders implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) GetCustomerOrders(ctx context.Context, customerID uuid.UUID) ([]*entity.Order, error) {
	panic("unimplemented")
}

// List implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) List(ctx context.Context, page int, pageSize int) ([]*entity.Customer, error) {
	customers, err := c.queries.ListCustomers(ctx, database.ListCustomersParams{
		Limit:  int32(pageSize),
		Offset: int32((page - 1) * pageSize),
	})
	if err != nil {
		return nil, fmt.Errorf("error listing customers: %w", err)
	}

	var result []*entity.Customer
	for _, customer := range customers {
		result = append(result, &entity.Customer{
			ID:    customer.ID,
			Name:  customer.Name,
			Email: customer.Email,
			Phone: customer.Phone.String,
		})
	}

	return result, nil
}

// RemoveCustomerOrder implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) RemoveCustomerOrder(ctx context.Context, orderID uuid.UUID, customerID uuid.UUID) error {
	panic("unimplemented")
}

// Update implements repository.CustomerRepositoryInterface.
func (c *CustomerRepository) Update(ctx context.Context, customer *entity.Customer) error {
	err := c.queries.UpdateCustomer(ctx, database.UpdateCustomerParams{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		Phone:     sql.NullString{String: customer.Phone, Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("error updating customer: %w", err)
	}

	return nil
}

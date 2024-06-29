package service

import (
	"context"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type PaymentServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Payment, error)
	Create(ctx context.Context, input *dto.CreatePaymentInput) (*entity.Payment, error)
	Update(ctx context.Context, input *dto.UpdatePaymentInput) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Payment, error)
}

type paymentService struct {
	repo domain.PaymentRepositoryInterface
}

func NewPaymentService(repo domain.PaymentRepositoryInterface) PaymentServiceInterface {
	return &paymentService{
		repo: repo,
	}
}

// Create implements PaymentServiceInterface.
func (p *paymentService) Create(ctx context.Context, input *dto.CreatePaymentInput) (*entity.Payment, error) {
	payment := input.ToEntity()
	if err := payment.Validate(); err != nil {
		return nil, fmt.Errorf("payment validation failed: %w", err)
	}

	payment, err := p.repo.Create(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return payment, nil
}

// Delete implements PaymentServiceInterface.
func (p *paymentService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := p.repo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get payment: %w", err)
	}

	if err := p.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete payment: %w", err)
	}

	return nil
}

// Get implements PaymentServiceInterface.
func (p *paymentService) Get(ctx context.Context, id uuid.UUID) (*entity.Payment, error) {
	payment, err := p.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}

	return payment, nil
}

// List implements PaymentServiceInterface.
func (p *paymentService) List(ctx context.Context, page int, pageSize int) ([]*entity.Payment, error) {
	payments, err := p.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to list payments: %w", err)
	}

	return payments, nil
}

// Update implements PaymentServiceInterface.
func (p *paymentService) Update(ctx context.Context, input *dto.UpdatePaymentInput) error {
	_, err := p.repo.Get(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("failed to get payment: %w", err)
	}

	payment := input.ToEntity()
	if err := payment.Validate(); err != nil {
		return fmt.Errorf("payment validation failed: %w", err)
	}

	if err := p.repo.Update(ctx, payment); err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	return nil
}

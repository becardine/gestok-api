package service

import (
	"context"
	"fmt"

	"github.com/becardine/gestock-api/internal/domain/entity"
	domain "github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
)

type FeedbackServiceInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Feedback, error)
	Create(ctx context.Context, input *dto.CreateFeedbackInput) (*entity.Feedback, error)
	Update(ctx context.Context, id uuid.UUID, input *dto.UpdateFeedbackInput) (*entity.Feedback, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Feedback, error)
}

type feedbackService struct {
	repo domain.FeedbackRepositoryInterface
}

func NewFeedbackService(repo domain.FeedbackRepositoryInterface) FeedbackServiceInterface {
	return &feedbackService{
		repo: repo,
	}
}

// Create implements FeedbackServiceInterface.
func (f *feedbackService) Create(ctx context.Context, input *dto.CreateFeedbackInput) (*entity.Feedback, error) {
	feedback := input.ToEntity()
	if err := feedback.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate feedback: %w", err)
	}

	feedback, err := f.repo.Create(ctx, feedback)
	if err != nil {
		return nil, fmt.Errorf("failed to create feedback: %w", err)
	}

	return feedback, nil
}

// Delete implements FeedbackServiceInterface.
func (f *feedbackService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := f.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get feedback: %w", err)
	}

	if err := f.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete feedback: %w", err)
	}

	return nil
}

// Get implements FeedbackServiceInterface.
func (f *feedbackService) Get(ctx context.Context, id uuid.UUID) (*entity.Feedback, error) {
	feedback, err := f.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get feedback: %w", err)
	}

	return feedback, nil
}

// List implements FeedbackServiceInterface.
func (f *feedbackService) List(ctx context.Context, page int, pageSize int) ([]*entity.Feedback, error) {
	feedbacks, err := f.repo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to list feedbacks: %w", err)
	}

	return feedbacks, nil
}

// Update implements FeedbackServiceInterface.
func (f *feedbackService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateFeedbackInput) (*entity.Feedback, error) {
	_, err := f.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get feedback: %w", err)
	}

	feedback := input.ToEntity()
	if err := feedback.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate feedback: %w", err)
	}

	if err := f.repo.Update(ctx, feedback); err != nil {
		return nil, fmt.Errorf("failed to update feedback: %w", err)
	}

	return feedback, nil
}

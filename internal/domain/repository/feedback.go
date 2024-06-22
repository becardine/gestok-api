package repository

import (
	"context"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/google/uuid"
)

type FeedbackRepositoryInterface interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Feedback, error)
	Create(ctx context.Context, feedback *entity.Feedback) (*entity.Feedback, error)
	Update(ctx context.Context, feedback *entity.Feedback) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Feedback, error)
	GetFeedbackOrders(ctx context.Context, feedbackID uuid.UUID) ([]*entity.Order, error)
	AddFeedbackOrder(ctx context.Context, orderID, feedbackID uuid.UUID) error
	RemoveFeedbackOrder(ctx context.Context, orderID, feedbackID uuid.UUID) error
}

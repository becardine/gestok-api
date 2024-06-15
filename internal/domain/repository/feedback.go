package repository

import (
	"context"
	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/entity/common"
)

type FeedbackRepositoryInterface interface {
	Get(ctx context.Context, id common.ID) (*entity.Feedback, error)
	Create(ctx context.Context, feedback *entity.Feedback) (*entity.Feedback, error)
	Update(ctx context.Context, feedback *entity.Feedback) error
	Delete(ctx context.Context, id common.ID) error
	List(ctx context.Context, page, pageSize int) ([]*entity.Feedback, error)
	GetFeedbackOrders(ctx context.Context, feedbackID common.ID) ([]*entity.Order, error)
	AddFeedbackOrder(ctx context.Context, orderID, feedbackID common.ID) error
	RemoveFeedbackOrder(ctx context.Context, orderID, feedbackID common.ID) error
}

package service_test

import (
	"context"
	"testing"

	"github.com/becardine/gestock-api/internal/domain/entity"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FeedbackRepositoryMock struct {
	mock.Mock
}

func (m *FeedbackRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*entity.Feedback, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Feedback), args.Error(1)
}

func (m *FeedbackRepositoryMock) Create(ctx context.Context, feedback *entity.Feedback) (*entity.Feedback, error) {
	args := m.Called(ctx, feedback)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Feedback), args.Error(1)
}

func (m *FeedbackRepositoryMock) Update(ctx context.Context, feedback *entity.Feedback) error {
	args := m.Called(ctx, feedback)
	return args.Error(0)
}

func (m *FeedbackRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *FeedbackRepositoryMock) List(ctx context.Context, page, pageSize int) ([]*entity.Feedback, error) {
	args := m.Called(ctx, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Feedback), args.Error(1)
}

func TestFeedbackService(t *testing.T) {
	t.Run("should create a feedback successfully", func(t *testing.T) {
		mockRepo := new(FeedbackRepositoryMock)
		feedbackService := service.NewFeedbackService(mockRepo)

		input := &dto.CreateFeedbackInput{
			OrderID:    uuid.New(),
			CustomerID: uuid.New(),
			Comment:    "This is a comment",
			Rating:     5,
		}

		feedback := input.ToEntity()
		mockRepo.On("Create", mock.Anything, mock.Anything).Return(feedback, nil)
		feedback, err := feedbackService.Create(context.Background(), input)

		assert.Nil(t, err)
		assert.NotNil(t, feedback)
		assert.Equal(t, input.OrderID, feedback.OrderID)
		assert.Equal(t, input.Comment, feedback.Comment)
		assert.Equal(t, input.Rating, feedback.Rating)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should get a feedback successfully", func(t *testing.T) {
		mockRepo := new(FeedbackRepositoryMock)
		feedbackService := service.NewFeedbackService(mockRepo)

		id := uuid.New()
		feedback := &entity.Feedback{
			ID:      id,
			OrderID: uuid.New(),
			Comment: "This is a comment",
			Rating:  5,
		}

		mockRepo.On("Get", mock.Anything, id).Return(feedback, nil)
		feedback, err := feedbackService.Get(context.Background(), id)

		assert.NoError(t, err)
		assert.NotNil(t, feedback)
		assert.Equal(t, id, feedback.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should list feedbacks successfully", func(t *testing.T) {
		mockRepo := new(FeedbackRepositoryMock)
		feedbackService := service.NewFeedbackService(mockRepo)

		page := 1
		pageSize := 10
		feedbacks := []*entity.Feedback{
			{
				ID:      uuid.New(),
				OrderID: uuid.New(),
				Comment: "This is a comment",
				Rating:  5,
			},
		}

		mockRepo.On("List", mock.Anything, page, pageSize).Return(feedbacks, nil)
		feedbacks, err := feedbackService.List(context.Background(), page, pageSize)

		assert.NoError(t, err)
		assert.NotNil(t, feedbacks)
		assert.Len(t, feedbacks, 1)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should delete a feedback successfully", func(t *testing.T) {
		mockRepo := new(FeedbackRepositoryMock)
		feedbackService := service.NewFeedbackService(mockRepo)

		id := uuid.New()
		feedback := &entity.Feedback{
			ID:      id,
			OrderID: uuid.New(),
			Comment: "This is a comment",
			Rating:  5,
		}

		mockRepo.On("Get", mock.Anything, id).Return(feedback, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(nil)
		err := feedbackService.Delete(context.Background(), id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not delete a feedback if it does not exist", func(t *testing.T) {
		mockRepo := new(FeedbackRepositoryMock)
		feedbackService := service.NewFeedbackService(mockRepo)

		id := uuid.New()
		mockRepo.On("Get", mock.Anything, id).Return(nil, assert.AnError)
		err := feedbackService.Delete(context.Background(), id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not delete a feedback if an error occurs", func(t *testing.T) {
		mockRepo := new(FeedbackRepositoryMock)
		feedbackService := service.NewFeedbackService(mockRepo)

		id := uuid.New()
		feedback := &entity.Feedback{
			ID:      id,
			OrderID: uuid.New(),
			Comment: "This is a comment",
			Rating:  5,
		}

		mockRepo.On("Get", mock.Anything, id).Return(feedback, nil)
		mockRepo.On("Delete", mock.Anything, id).Return(assert.AnError)
		err := feedbackService.Delete(context.Background(), id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should update a feedback successfully", func(t *testing.T) {
		mockRepo := new(FeedbackRepositoryMock)
		feedbackService := service.NewFeedbackService(mockRepo)

		id := uuid.New()
		input := &dto.UpdateFeedbackInput{
			ID:         id,
			OrderID:    uuid.New(),
			CustomerID: uuid.New(),
			Comment:    "This is a comment",
			Rating:     5,
		}

		feedback := input.ToEntity()

		mockRepo.On("Get", mock.Anything, id).Return(feedback, nil)
		mockRepo.On("Update", mock.Anything, mock.Anything).Return(feedback, nil)
		feedback, err := feedbackService.Update(context.Background(), id, input)

		assert.Nil(t, err)
		assert.NotNil(t, feedback)
		assert.Equal(t, id, feedback.ID)
		assert.Equal(t, input.Comment, feedback.Comment)
		assert.Equal(t, input.Rating, feedback.Rating)
		mockRepo.AssertExpectations(t)
	})
}

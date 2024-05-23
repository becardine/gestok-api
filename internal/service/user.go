package service

import (
	"github.com/becardine/gestock-api/internal/entity"
	"github.com/becardine/gestock-api/internal/repository"
)

type UserService interface {
	CreateUser(name, email, password string) error
	FindUserByEmail(email string) (*entity.UserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(name, email, password string) error {
	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return err
	}

	return s.userRepository.Create(user)
}

func (s *userService) FindUserByEmail(email string) (*entity.UserResponse, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
	}, nil
}

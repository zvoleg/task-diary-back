package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/repositories"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) services.UserService {
	return &userService{repository: repository}
}

func (service *userService) Get(ctx context.Context, userId uuid.UUID) (*models.UserResponse, error) {
	user, err := service.repository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *userService) Create(ctx context.Context, userRequest *models.UserRequest) (*models.UserResponse, error) {
	user, err := service.repository.Create(ctx, userRequest)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update implements user.Service
func (service *userService) Update(ctx context.Context, userId uuid.UUID, userRequest *models.UserRequest) (*models.UserResponse, error) {
	user, err := service.repository.Update(ctx, userId, userRequest)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete implements user.Service
func (service *userService) Delete(ctx context.Context, userId uuid.UUID) error {
	err := service.repository.Delete(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

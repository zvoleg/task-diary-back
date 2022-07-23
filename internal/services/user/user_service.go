package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/repositories"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) services.UserService {
	return &userService{repo: repo}
}

func (serv *userService) Get(userId uuid.UUID) (*models.UserResponse, error) {
	ctx := context.Background()
	user, err := serv.repo.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (serv *userService) Create(userRequest *models.UserRequest) (*models.UserResponse, error) {
	ctx := context.Background()
	user, err := serv.repo.Create(ctx, userRequest)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (serv *userService) Update(userId uuid.UUID, userRequest *models.UserRequest) (*models.UserResponse, error) {
	ctx := context.Background()
	user, err := serv.repo.Update(ctx, userId, userRequest)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (serv *userService) Delete(userId uuid.UUID) error {
	ctx := context.Background()
	err := serv.repo.Delete(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/repositories"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (rep *userRepository) Create(ctx context.Context, userRequest *models.UserRequest) (*models.UserResponse, error) {
	userDb := new(models.UserDb)
	if err := rep.db.QueryRowxContext(
		ctx,
		createScript,
		uuid.New(),
		nil,
		userRequest.Name,
		userRequest.Description,
		userRequest.RoleId,
		time.Now().UTC(),
		false,
	).StructScan(userDb); err != nil {
		return nil, errors.Wrap(err, "usersRepo.Create.StructScan")
	}
	user := userDb.Map()
	return &user, nil
}

func (rep *userRepository) Update(ctx context.Context, identifier uuid.UUID, userRequest *models.UserRequest) (*models.UserResponse, error) {
	user, err := rep.getDb(ctx, identifier)
	if err != nil {
		return nil, errors.Wrap(err, "userRepo.Update: user not found")
	}
	updatedUser := new(models.UserResponse)
	if err = rep.db.QueryRowxContext(
		ctx,
		updateScript,
		user.UserId,
		userRequest.Name,
		userRequest.Description,
		userRequest.RoleId,
		time.Now().UTC(),
	).StructScan(updatedUser); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (rep *userRepository) Get(ctx context.Context, identifier uuid.UUID) (*models.UserResponse, error) {
	user, err := rep.getDb(ctx, identifier)
	if err != nil {
		return nil, errors.Wrap(err, "userRepo.Get.StructScan")
	}
	if user.IsDeleted {
		return nil, errors.New("userRepo.Get: User not founded")
	}

	userResponse := user.Map()
	return &userResponse, nil
}

func (rep *userRepository) GetList(ctx context.Context) ([]*models.AllUserResponse, error) {
	return nil, nil
}

func (rep *userRepository) Delete(ctx context.Context, identifier uuid.UUID) error {
	_, err := rep.getDb(ctx, identifier)
	if err != nil {
		return err
	}
	if _, err := rep.db.Queryx(deleteScript, identifier, time.Now().UTC()); err != nil {
		return err
	}
	return nil
}

func (rep *userRepository) getDb(ctx context.Context, identifier uuid.UUID) (*models.UserDb, error) {
	user := new(models.UserDb)
	row := rep.db.QueryRowxContext(
		ctx,
		getScript,
		identifier,
	)
	if err := row.StructScan(user); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, &models.ErrNotFoundInDb{Msg: err.Error()}
		default:
			return nil, err
		}
	}
	return user, nil
}

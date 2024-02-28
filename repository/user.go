package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var userResult entity.User
	query := r.db.First(&userResult, id)
	if query.Error != nil {
		return entity.User{}, query.Error
	}
	return userResult, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	query := r.db.First(&userResult, "email = ?", email)
	if query.Error != nil {
		if query.Error == logger.ErrRecordNotFound {
			return entity.User{}, nil
		}
		return entity.User{}, query.Error
	}
	return userResult, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := r.db.Create(&user)
	if query.Error != nil {
		return entity.User{}, query.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := r.db.Save(&user)
	if query.Error != nil {
		return entity.User{}, query.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := r.db.Delete(&entity.User{}, id)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

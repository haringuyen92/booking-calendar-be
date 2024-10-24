package user_repositories

import (
	user_dto "booking-calendar-server-backend/internal/modules/user/dto"
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_scopes "booking-calendar-server-backend/internal/modules/user/scopes"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetOne(filter *user_filters.UserFilter) (*user_models.User, error)
	GetMany(filter *user_filters.UserFilter) ([]*user_models.User, error)
	Create(user *user_dto.CreateUserDTO) error
	Update(filter *user_filters.UserFilter, user *user_dto.UpdateUserDTO) error
	Delete(filter *user_filters.UserFilter) error
}

func NewUserRepository(
	db *gorm.DB,
) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetOne(filter *user_filters.UserFilter) (*user_models.User, error) {
	var user *user_models.User
	err := r.db.Scopes(user_scopes.UserScopes(filter)).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetMany(filter *user_filters.UserFilter) ([]*user_models.User, error) {
	var users []*user_models.User
	err := r.db.Scopes(user_scopes.UserScopes(filter)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(user *user_dto.CreateUserDTO) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(filter *user_filters.UserFilter, user *user_dto.UpdateUserDTO) error {
	model := &user_models.User{}
	err := r.db.Model(model).Scopes(user_scopes.UserScopes(filter)).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(filter *user_filters.UserFilter) error {
	err := r.db.Scopes(user_scopes.UserScopes(filter)).Delete(&user_models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

package user_services

import (
	user_dto "booking-calendar-server-backend/internal/modules/user/dto"
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	user_mappers "booking-calendar-server-backend/internal/modules/user/mappers"
	user_repositories "booking-calendar-server-backend/internal/modules/user/repositories"
	user_responses "booking-calendar-server-backend/internal/modules/user/responses"
)

type UserService interface {
	Create(dto *user_dto.CreateUserDTO) error
	Update(filter *user_filters.UserFilter, dto *user_dto.UpdateUserDTO) error
	GetOne(filter *user_filters.UserFilter) (*user_responses.GetUserResponse, error)
	GetMany(filter *user_filters.UserFilter) ([]*user_responses.GetAllUserResponse, error)
	Delete(filter *user_filters.UserFilter) error
}

func NewUserService(
	userRepository user_repositories.UserRepository,
) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type userService struct {
	userRepository user_repositories.UserRepository
}

func (s *userService) Create(dto *user_dto.CreateUserDTO) error {
	err := s.userRepository.Create(dto)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Update(filter *user_filters.UserFilter, dto *user_dto.UpdateUserDTO) error {
	err := s.userRepository.Update(filter, dto)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetOne(filter *user_filters.UserFilter) (*user_responses.GetUserResponse, error) {
	res, err := s.userRepository.GetOne(filter)
	if err != nil {
		return nil, err
	}
	return user_mappers.GetUserResponseMapper(res), nil
}

func (s *userService) GetMany(filter *user_filters.UserFilter) ([]*user_responses.GetAllUserResponse, error) {
	res, err := s.userRepository.GetMany(filter)
	if err != nil {
		return nil, err
	}
	return user_mappers.GetAllUserResponse(res), nil
}

func (s *userService) Delete(filter *user_filters.UserFilter) error {
	err := s.userRepository.Delete(filter)
	if err != nil {
		return err
	}
	return nil
}

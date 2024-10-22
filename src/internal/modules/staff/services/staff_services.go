package staff_services

import (
	staff_dto "booking-calendar-server-backend/internal/modules/staff/dto"
	staff_filter "booking-calendar-server-backend/internal/modules/staff/filters"
	staff_mappers "booking-calendar-server-backend/internal/modules/staff/mappers"
	staff_repositories "booking-calendar-server-backend/internal/modules/staff/repositories"
	staff_responses "booking-calendar-server-backend/internal/modules/staff/responses"
)

type StaffService interface {
	Create(dto *staff_dto.CreateStaffDto) error
	Update(filter *staff_filter.StaffFilter, dto *staff_dto.UpdateStaffDto) error
	Delete(filter *staff_filter.StaffFilter) error
	GetOne(filter *staff_filter.StaffFilter) (*staff_responses.GetStaffResponse, error)
	GetMany(filter *staff_filter.StaffFilter) ([]*staff_responses.GetAllStaffResponse, error)
}

func NewStaffService(
	staffRepository staff_repositories.StaffRepository,
) StaffService {
	return &staffService{
		staffRepository: staffRepository,
	}
}

type staffService struct {
	staffRepository staff_repositories.StaffRepository
}

func (s *staffService) Create(dto *staff_dto.CreateStaffDto) error {
	return s.staffRepository.Create(dto)
}

func (s *staffService) Update(filter *staff_filter.StaffFilter, dto *staff_dto.UpdateStaffDto) error {
	return s.staffRepository.Update(filter, dto)
}

func (s *staffService) Delete(filter *staff_filter.StaffFilter) error {
	return s.staffRepository.Delete(filter)
}

func (s *staffService) GetOne(filter *staff_filter.StaffFilter) (*staff_responses.GetStaffResponse, error) {
	staff, err := s.staffRepository.GetOne(filter)
	if err != nil {
		return nil, err
	}
	return staff_mappers.GetStaffResponseMapper(staff), nil
}

func (s *staffService) GetMany(filter *staff_filter.StaffFilter) ([]*staff_responses.GetAllStaffResponse, error) {
	staffs, err := s.staffRepository.GetMany(filter)
	if err != nil {
		return nil, err
	}
	return staff_mappers.GetAllStaffResponseMapper(staffs), nil
}

package course_services

import (
	course_dto "booking-calendar-server-backend/internal/modules/course/dto"
	course_filters "booking-calendar-server-backend/internal/modules/course/filters"
	course_mappers "booking-calendar-server-backend/internal/modules/course/mappers"
	course_repositories "booking-calendar-server-backend/internal/modules/course/repositories"
	course_responses "booking-calendar-server-backend/internal/modules/course/responses"
)

type CourseService interface {
	Create(request *course_dto.CreateCourseDto) error
	Update(filter *course_filters.CourseFilter, request *course_dto.UpdateCourseDto) error
	GetOne(filter *course_filters.CourseFilter) (*course_responses.GetCourseResponse, error)
	GetMany(filter *course_filters.CourseFilter) ([]*course_responses.GetAllCourseResponse, error)
	DeleteByID(ID uint) error
	Delete(filter *course_filters.CourseFilter) error
}

func NewCourseService(
	courseRepository course_repositories.CourseRepository,
) CourseService {
	return &courseService{
		courseRepository: courseRepository,
	}
}

type courseService struct {
	courseRepository course_repositories.CourseRepository
}

func (s *courseService) Create(request *course_dto.CreateCourseDto) error {
	return s.courseRepository.Create(request)
}

func (s *courseService) Update(filter *course_filters.CourseFilter, request *course_dto.UpdateCourseDto) error {
	return s.courseRepository.Update(filter, request)
}

func (s *courseService) GetOne(filter *course_filters.CourseFilter) (*course_responses.GetCourseResponse, error) {
	res, err := s.courseRepository.GetOne(&course_filters.CourseFilter{
		ID: filter.ID,
	})
	if err != nil {
		return nil, err
	}
	return course_mappers.GetCourseResponseMapper(res), nil
}

func (s *courseService) GetMany(filter *course_filters.CourseFilter) ([]*course_responses.GetAllCourseResponse, error) {
	res, err := s.courseRepository.GetMany(
		&course_filters.CourseFilter{
			ID: filter.ID,
		},
	)
	if err != nil {
		return nil, err
	}
	return course_mappers.GetAllCourseResponseMapper(res), nil
}

func (s *courseService) DeleteByID(ID uint) error {
	return s.courseRepository.DeleteByID(ID)
}

func (s *courseService) Delete(filter *course_filters.CourseFilter) error {
	return s.courseRepository.Delete(filter)
}

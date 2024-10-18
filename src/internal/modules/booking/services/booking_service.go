package booking_services

import (
	booking_dto "booking-calendar-server-backend/internal/modules/booking/dto"
	booking_filters "booking-calendar-server-backend/internal/modules/booking/filters"
	booking_mappers "booking-calendar-server-backend/internal/modules/booking/mappers"
	booking_repositories "booking-calendar-server-backend/internal/modules/booking/repositories"
	booking_responses "booking-calendar-server-backend/internal/modules/booking/responses"
	"fmt"
	"time"
)

type BookingService interface {
	Create(dto *booking_dto.CreateBookingDto) error
	Update(filter *booking_filters.BookingFilter, dto *booking_dto.UpdateBookingDto) error
	Delete(filter *booking_filters.BookingFilter) error
	GetOne(filter *booking_filters.BookingFilter) (*booking_responses.GetBookingResponse, error)
	GetMany(filter *booking_filters.BookingFilter) ([]*booking_responses.GetAllBookingResponse, error)
}

type bookingService struct {
	bookingRepository booking_repositories.BookingRepository
}

func NewBookingService(
	bookingRepository booking_repositories.BookingRepository,
) BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

func (s *bookingService) Create(dto *booking_dto.CreateBookingDto) error {
	return s.bookingRepository.Create(dto)
}

func (s *bookingService) Update(filter *booking_filters.BookingFilter, dto *booking_dto.UpdateBookingDto) error {
	return s.bookingRepository.Update(filter, dto)
}

func (s *bookingService) Delete(filter *booking_filters.BookingFilter) error {
	return s.bookingRepository.Delete(filter)
}

func (s *bookingService) GetOne(filter *booking_filters.BookingFilter) (*booking_responses.GetBookingResponse, error) {
	now := time.Now()
	defer func() {
		fmt.Println("[bookingService] GetOne took payload: ", filter, time.Since(now))
	}()
	booking, err := s.bookingRepository.GetOne(filter)
	if err != nil {
		fmt.Println("[bookingService] GetOne err:", err)
		return nil, err
	}
	return booking_mappers.GetBookingResponseMapper(booking), nil
}

func (s *bookingService) GetMany(filter *booking_filters.BookingFilter) ([]*booking_responses.GetAllBookingResponse, error) {
	bookings, err := s.bookingRepository.GetMany(filter)
	if err != nil {
		return nil, err
	}
	return booking_mappers.GetAllBookingResponseMapper(bookings), nil
}

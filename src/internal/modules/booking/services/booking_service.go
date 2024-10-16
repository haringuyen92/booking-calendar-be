package services

import (
	"booking-calendar-server-backend/internal/modules/booking/dto"
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"booking-calendar-server-backend/internal/modules/booking/mappers"
	"booking-calendar-server-backend/internal/modules/booking/repositories"
	"booking-calendar-server-backend/internal/modules/booking/responses"
	"fmt"
	"time"
)

type BookingService interface {
	Create(dto *dto.CreateBookingDto) error
	Update(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error
	Delete(filter *filters.BookingFilter) error
	GetOne(filter *filters.BookingFilter) (*responses.GetBookingResponse, error)
	GetMany(filter *filters.BookingFilter) ([]*responses.GetAllBookingResponse, error)
}

type bookingService struct {
	bookingRepository repositories.BookingRepository
}

func NewBookingService(
	bookingRepository repositories.BookingRepository,
) BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

func (s *bookingService) Create(dto *dto.CreateBookingDto) error {
	return s.bookingRepository.Create(dto)
}

func (s *bookingService) Update(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error {
	return s.bookingRepository.Update(filter, dto)
}

func (s *bookingService) Delete(filter *filters.BookingFilter) error {
	return s.bookingRepository.Delete(filter)
}

func (s *bookingService) GetOne(filter *filters.BookingFilter) (*responses.GetBookingResponse, error) {
	now := time.Now()
	defer func() {
		fmt.Println("[bookingService] GetOne took payload: ", filter, time.Since(now))
	}()
	booking, err := s.bookingRepository.GetOne(filter)
	if err != nil {
		fmt.Println("[bookingService] GetOne err:", err)
		return nil, err
	}
	return mappers.GetBookingResponseMapper(booking), nil
}

func (s *bookingService) GetMany(filter *filters.BookingFilter) ([]*responses.GetAllBookingResponse, error) {
	bookings, err := s.bookingRepository.GetMany(filter)
	if err != nil {
		return nil, err
	}
	return mappers.GetAllBookingResponseMapper(bookings), nil
}

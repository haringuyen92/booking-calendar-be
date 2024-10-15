package services

import (
	"booking-calendar-server-backend/internal/modules/booking/dto"
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"booking-calendar-server-backend/internal/modules/booking/mappers"
	"booking-calendar-server-backend/internal/modules/booking/repositories"
	"booking-calendar-server-backend/internal/modules/booking/responses"
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

func (b *bookingService) Create(dto *dto.CreateBookingDto) error {
	return b.bookingRepository.Create(dto)
}

func (b *bookingService) Update(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error {
	return b.bookingRepository.Update(filter, dto)
}

func (b *bookingService) Delete(filter *filters.BookingFilter) error {
	return b.bookingRepository.Delete(filter)
}

func (b *bookingService) GetOne(filter *filters.BookingFilter) (*responses.GetBookingResponse, error) {
	booking, err := b.bookingRepository.GetOne(filter)
	if err != nil {
		return nil, err
	}
	return mappers.GetBookingResponseMapper(booking), nil
}

func (b *bookingService) GetMany(filter *filters.BookingFilter) ([]*responses.GetAllBookingResponse, error) {
	bookings, err := b.bookingRepository.GetMany(filter)
	if err != nil {
		return nil, err
	}
	return mappers.GetAllBookingResponseMapper(bookings), nil
}

package services

import (
	"booking-calendar-server-backend/internal/modules/booking/dto"
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"booking-calendar-server-backend/internal/modules/booking/mappers"
	"booking-calendar-server-backend/internal/modules/booking/repositories"
	"booking-calendar-server-backend/internal/modules/booking/responses"
)

type BookingService interface {
	CreateBooking(dto *dto.CreateBookingDto) error
	UpdateBooking(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error
	DeleteBooking(filter *filters.BookingFilter) error
	GetBooking(filter *filters.BookingFilter) (*responses.GetBookingResponse, error)
	GetBookingList(filter *filters.BookingFilter) ([]*responses.GetAllBookingResponse, error)
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

func (b *bookingService) CreateBooking(dto *dto.CreateBookingDto) error {
	return b.bookingRepository.Create(dto)
}

func (b *bookingService) UpdateBooking(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error {
	return b.bookingRepository.Update(filter, dto)
}

func (b *bookingService) DeleteBooking(filter *filters.BookingFilter) error {
	return b.bookingRepository.Delete(filter)
}

func (b *bookingService) GetBooking(filter *filters.BookingFilter) (*responses.GetBookingResponse, error) {
	booking, err := b.bookingRepository.GetOne(filter)
	if err != nil {
		return nil, err
	}
	return mappers.GetBookingResponseMapper(booking), nil
}

func (b *bookingService) GetBookingList(filter *filters.BookingFilter) ([]*responses.GetAllBookingResponse, error) {
	bookings, err := b.bookingRepository.GetMany(filter)
	if err != nil {
		return nil, err
	}
	return mappers.GetAllBookingResponseMapper(bookings), nil
}

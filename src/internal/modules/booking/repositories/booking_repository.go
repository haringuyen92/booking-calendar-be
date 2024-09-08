package repositories

import (
	"booking-calendar-server-backend/internal/modules/booking/dto"
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"booking-calendar-server-backend/internal/modules/booking/models"
)

type BookingRepository interface {
	GetBookingByID(ID uint) (*models.Booking, error)
	GetBookingByStoreID(storeID uint) ([]*models.Booking, error)
	GetBookingByUserID(userID uint) ([]*models.Booking, error)
	GetOne(filter *filters.BookingFilter) (*models.Booking, error)
	GetMany(filter *filters.BookingFilter) ([]*models.Booking, error)
	Create(dto *dto.CreateBookingDto) error
	Update(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error
	DeleteBookingByID(ID uint) error
	Delete(filter *filters.BookingFilter) error
}

type bookingRepository struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{}
}
func (b *bookingRepository) GetBookingByID(ID uint) (*models.Booking, error) {
	return b.GetOne(&filters.BookingFilter{
		ID: ID,
	})
}

func (b *bookingRepository) GetBookingByStoreID(storeID uint) ([]*models.Booking, error) {
	return b.GetMany(&filters.BookingFilter{
		StoreID: storeID,
	})
}

func (b *bookingRepository) GetBookingByUserID(userID uint) ([]*models.Booking, error) {
	return b.GetMany(&filters.BookingFilter{
		UserID: userID,
	})
}

func (b *bookingRepository) GetOne(filter *filters.BookingFilter) (*models.Booking, error) {
	return nil, nil
}

func (b *bookingRepository) GetMany(filter *filters.BookingFilter) ([]*models.Booking, error) {
	var bookings []*models.Booking
	return bookings, nil
}

func (b *bookingRepository) Create(dto *dto.CreateBookingDto) error {
	return nil
}

func (b *bookingRepository) Update(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error {
	return nil
}

func (b *bookingRepository) DeleteBookingByID(ID uint) error {
	return b.Delete(&filters.BookingFilter{ID: ID})
}

func (b *bookingRepository) Delete(filter *filters.BookingFilter) error {
	return nil
}

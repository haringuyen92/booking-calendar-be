package booking_repositories

import (
	booking_dto "booking-calendar-server-backend/internal/modules/booking/dto"
	booking_errors "booking-calendar-server-backend/internal/modules/booking/errors"
	booking_filters "booking-calendar-server-backend/internal/modules/booking/filters"
	booking_models "booking-calendar-server-backend/internal/modules/booking/models"
	booking_scopes "booking-calendar-server-backend/internal/modules/booking/scopes"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BookingRepository interface {
	GetBookingByID(ID string) (*booking_models.Booking, error)
	GetBookingByStoreID(storeID uint) ([]*booking_models.Booking, error)
	GetBookingByUserID(userID uint) ([]*booking_models.Booking, error)
	GetOne(filter *booking_filters.BookingFilter) (*booking_models.Booking, error)
	GetMany(filter *booking_filters.BookingFilter) ([]*booking_models.Booking, error)
	Create(dto *booking_dto.CreateBookingDto) error
	Update(filter *booking_filters.BookingFilter, dto *booking_dto.UpdateBookingDto) error
	DeleteByID(ID string) error
	Delete(filter *booking_filters.BookingFilter) error
}

type bookingRepository struct {
	bookingCollection *mongo.Collection
}

func NewBookingRepository(
	mongoDB *mongo.Database,
) BookingRepository {
	bookingCollection := mongoDB.Collection("bookings")
	return &bookingRepository{
		bookingCollection: bookingCollection,
	}
}
func (r *bookingRepository) GetBookingByID(ID string) (*booking_models.Booking, error) {
	return r.GetOne(&booking_filters.BookingFilter{
		ID: ID,
	})
}

func (r *bookingRepository) GetBookingByStoreID(storeID uint) ([]*booking_models.Booking, error) {
	return r.GetMany(&booking_filters.BookingFilter{
		StoreID: storeID,
	})
}

func (r *bookingRepository) GetBookingByUserID(userID uint) ([]*booking_models.Booking, error) {
	return r.GetMany(&booking_filters.BookingFilter{
		UserID: userID,
	})
}

func (r *bookingRepository) GetOne(filter *booking_filters.BookingFilter) (*booking_models.Booking, error) {
	var booking *booking_models.Booking
	scope := booking_scopes.BookingScope(filter)
	if len(scope) == 0 {
		return nil, booking_errors.BookingNotFoundError
	}

	err := r.bookingCollection.FindOne(context.Background(), scope).Decode(&booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *bookingRepository) GetMany(filter *booking_filters.BookingFilter) ([]*booking_models.Booking, error) {
	var bookings []*booking_models.Booking
	cursor, err := r.bookingCollection.Find(context.Background(), booking_scopes.BookingScope(filter))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var booking booking_models.Booking
		if err := cursor.Decode(&booking); err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}
	return bookings, nil
}

func (r *bookingRepository) Create(dto *booking_dto.CreateBookingDto) error {
	now := time.Now()
	dto.CreatedAt = &now
	dto.UpdatedAt = &now
	_, err := r.bookingCollection.InsertOne(context.Background(), dto)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) Update(filter *booking_filters.BookingFilter, dto *booking_dto.UpdateBookingDto) error {
	now := time.Now()
	dto.UpdatedAt = now

	_, err := r.bookingCollection.UpdateOne(context.Background(), booking_scopes.BookingScope(filter), dto)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) DeleteByID(ID string) error {
	return r.Delete(&booking_filters.BookingFilter{ID: ID})
}

func (r *bookingRepository) Delete(filter *booking_filters.BookingFilter) error {
	now := time.Now()
	return r.Update(filter, &booking_dto.UpdateBookingDto{
		DeletedAt: now,
	})
}

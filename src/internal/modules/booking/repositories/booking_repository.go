package repositories

import (
	"booking-calendar-server-backend/internal/modules/booking/dto"
	"booking-calendar-server-backend/internal/modules/booking/errors"
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"booking-calendar-server-backend/internal/modules/booking/models"
	"booking-calendar-server-backend/internal/modules/booking/scopes"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BookingRepository interface {
	GetBookingByID(ID string) (*models.Booking, error)
	GetBookingByStoreID(storeID uint) ([]*models.Booking, error)
	GetBookingByUserID(userID uint) ([]*models.Booking, error)
	GetOne(filter *filters.BookingFilter) (*models.Booking, error)
	GetMany(filter *filters.BookingFilter) ([]*models.Booking, error)
	Create(dto *dto.CreateBookingDto) error
	Update(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error
	DeleteByID(ID string) error
	Delete(filter *filters.BookingFilter) error
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
func (r *bookingRepository) GetBookingByID(ID string) (*models.Booking, error) {
	return r.GetOne(&filters.BookingFilter{
		ID: ID,
	})
}

func (r *bookingRepository) GetBookingByStoreID(storeID uint) ([]*models.Booking, error) {
	return r.GetMany(&filters.BookingFilter{
		StoreID: storeID,
	})
}

func (r *bookingRepository) GetBookingByUserID(userID uint) ([]*models.Booking, error) {
	return r.GetMany(&filters.BookingFilter{
		UserID: userID,
	})
}

func (r *bookingRepository) GetOne(filter *filters.BookingFilter) (*models.Booking, error) {
	var booking *models.Booking
	scope := scopes.BookingScope(filter)
	if len(scope) == 0 {
		return nil, errors.BookingNotFoundError
	}

	err := r.bookingCollection.FindOne(context.Background(), scope).Decode(&booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *bookingRepository) GetMany(filter *filters.BookingFilter) ([]*models.Booking, error) {
	var bookings []*models.Booking
	cursor, err := r.bookingCollection.Find(context.Background(), scopes.BookingScope(filter))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var booking models.Booking
		if err := cursor.Decode(&booking); err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}
	return bookings, nil
}

func (r *bookingRepository) Create(dto *dto.CreateBookingDto) error {
	now := time.Now()
	dto.CreatedAt = &now
	dto.UpdatedAt = &now
	_, err := r.bookingCollection.InsertOne(context.Background(), dto)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) Update(filter *filters.BookingFilter, dto *dto.UpdateBookingDto) error {
	now := time.Now()
	dto.UpdatedAt = now

	_, err := r.bookingCollection.UpdateOne(context.Background(), scopes.BookingScope(filter), dto)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) DeleteByID(ID string) error {
	return r.Delete(&filters.BookingFilter{ID: ID})
}

func (r *bookingRepository) Delete(filter *filters.BookingFilter) error {
	now := time.Now()
	return r.Update(filter, &dto.UpdateBookingDto{
		DeletedAt: now,
	})
}

package stores_repositories

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_errors "booking-calendar-server-backend/internal/modules/store/errors"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	store_scopes "booking-calendar-server-backend/internal/modules/store/scopes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SettingBookingRepository interface {
	GetSettingByStoreID(storeID uint) (*common_dto.SettingBookingDto, error)
	UpdateSettingByStoreID(storeID uint, setting *common_dto.SettingBookingDto) error
	GetOne(filter *store_filter.SettingBookingFilter) (*common_dto.SettingBookingDto, error)
	Create(dto *common_dto.SettingBookingDto) error
	Update(filter *store_filter.SettingBookingFilter, update *common_dto.SettingBookingDto) error
}

func NewSettingBookingRepository(
	mongoDB *mongo.Database,
) SettingBookingRepository {
	collection := mongoDB.Collection("setting_booking")
	return &settingBookingRepository{
		collection: collection,
	}
}

type settingBookingRepository struct {
	collection *mongo.Collection
}

func (r *settingBookingRepository) GetSettingByStoreID(storeID uint) (*common_dto.SettingBookingDto, error) {
	return r.GetOne(&store_filter.SettingBookingFilter{
		StoreID: storeID,
	})
}

func (r *settingBookingRepository) GetOne(filter *store_filter.SettingBookingFilter) (*common_dto.SettingBookingDto, error) {
	var result common_dto.SettingBookingDto
	scope := store_scopes.SettingBookingScope(filter)
	if len(scope) == 0 {
		return nil, store_errors.SettingBookingNotFoundError
	}
	err := r.collection.FindOne(
		context.Background(),
		scope,
	).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *settingBookingRepository) Create(dto *common_dto.SettingBookingDto) error {
	_, err := r.collection.InsertOne(
		context.Background(),
		dto,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *settingBookingRepository) UpdateSettingByStoreID(storeID uint, setting *common_dto.SettingBookingDto) error {
	filter := store_filter.SettingBookingFilter{
		StoreID: storeID,
	}
	return r.Update(&filter, setting)
}

func (r *settingBookingRepository) Update(filter *store_filter.SettingBookingFilter, update *common_dto.SettingBookingDto) error {
	scope := store_scopes.SettingBookingScope(filter)
	if len(scope) == 0 {
		return store_errors.SettingBookingNotFoundError
	}
	_, err := r.collection.UpdateOne(
		context.Background(),
		scope,
		bson.M{
			"$set": update,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

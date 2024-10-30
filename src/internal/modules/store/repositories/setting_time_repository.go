package stores_repositories

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_errors "booking-calendar-server-backend/internal/modules/store/errors"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	store_scopes "booking-calendar-server-backend/internal/modules/store/scopes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SettingTimeRepository interface {
	GetSettingTime(storeID uint) (*common_dto.SettingTimeDto, error)
	UpdateSettingTime(storeID uint, settingTimeDto *common_dto.SettingTimeDto) error
	GetOne(filter *store_filter.SettingTimeFilter) (*common_dto.SettingTimeDto, error)
	Update(filter *store_filter.SettingTimeFilter, update *common_dto.SettingTimeDto) error
	Create(dto *common_dto.SettingTimeDto) error
}

func NewSettingTimeRepository(
	mongoDB *mongo.Database,
) SettingTimeRepository {
	collection := mongoDB.Collection("setting_time")
	return &settingTimeRepository{
		collection: collection,
	}
}

type settingTimeRepository struct {
	collection *mongo.Collection
}

func (r *settingTimeRepository) GetSettingTime(storeID uint) (*common_dto.SettingTimeDto, error) {
	return r.GetOne(&store_filter.SettingTimeFilter{
		StoreID: storeID,
	})
}

func (r *settingTimeRepository) UpdateSettingTime(storeID uint, settingTimeDto *common_dto.SettingTimeDto) error {
	return r.Update(&store_filter.SettingTimeFilter{
		StoreID: storeID,
	}, settingTimeDto)
}

func (r *settingTimeRepository) GetOne(filter *store_filter.SettingTimeFilter) (*common_dto.SettingTimeDto, error) {
	var settingTimeDto *common_dto.SettingTimeDto
	scope := store_scopes.SettingTimeScope(filter)
	if len(scope) == 0 {
		return nil, store_errors.SettingTimeNotFoundError
	}
	err := r.collection.FindOne(
		context.Background(),
		scope,
	).Decode(&settingTimeDto)
	if err != nil {
		return nil, err
	}
	return settingTimeDto, nil
}

func (r *settingTimeRepository) Update(filter *store_filter.SettingTimeFilter, update *common_dto.SettingTimeDto) error {
	scope := store_scopes.SettingTimeScope(filter)
	if len(scope) == 0 {
		return store_errors.SettingTimeNotFoundError
	}
	fmt.Println("start Update ", update.IsOpen)
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

func (r *settingTimeRepository) Create(dto *common_dto.SettingTimeDto) error {
	_, err := r.collection.InsertOne(context.Background(), dto)
	if err != nil {
		return err
	}
	return nil
}

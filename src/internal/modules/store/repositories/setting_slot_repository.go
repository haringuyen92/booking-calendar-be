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

type SettingSlotRepository interface {
	GetSettingByStoreID(storeID uint) (*common_dto.SettingSlotDto, error)
	UpdateSettingByStoreID(storeID uint, setting *common_dto.SettingSlotDto) error
	GetOne(filter *store_filter.SettingSlotFilter) (*common_dto.SettingSlotDto, error)
	Create(dto *common_dto.SettingSlotDto) error
	Update(filter *store_filter.SettingSlotFilter, update *common_dto.SettingSlotDto) error
	Delete(filter *store_filter.SettingSlotFilter) error
}

func NewSettingSlotRepository(
	mongoDB *mongo.Database,
) SettingSlotRepository {
	collection := mongoDB.Collection("setting_slot")
	return &settingSlotRepository{
		collection: collection,
	}

}

type settingSlotRepository struct {
	collection *mongo.Collection
}

func (s *settingSlotRepository) UpdateSettingByStoreID(storeID uint, setting *common_dto.SettingSlotDto) error {
	return s.Update(
		&store_filter.SettingSlotFilter{
			StoreID: storeID,
		},
		setting,
	)
}

func (s *settingSlotRepository) GetSettingByStoreID(storeID uint) (*common_dto.SettingSlotDto, error) {
	return s.GetOne(&store_filter.SettingSlotFilter{
		StoreID: storeID,
	})
}

func (s *settingSlotRepository) GetOne(filter *store_filter.SettingSlotFilter) (*common_dto.SettingSlotDto, error) {
	var result common_dto.SettingSlotDto
	scope := store_scopes.SettingSlotScope(filter)
	if len(scope) == 0 {
		return nil, store_errors.SettingSlotNotFoundError
	}
	err := s.collection.FindOne(
		context.Background(),
		scope,
	).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *settingSlotRepository) Create(dto *common_dto.SettingSlotDto) error {
	_, err := s.collection.InsertOne(
		context.Background(),
		dto,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *settingSlotRepository) Update(filter *store_filter.SettingSlotFilter, update *common_dto.SettingSlotDto) error {
	scope := store_scopes.SettingSlotScope(filter)
	if len(scope) == 0 {
		return store_errors.SettingSlotNotFoundError
	}
	_, err := s.collection.UpdateOne(
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

func (s *settingSlotRepository) Delete(filter *store_filter.SettingSlotFilter) error {
	scope := store_scopes.SettingSlotScope(filter)
	if len(scope) == 0 {
		return store_errors.SettingSlotNotFoundError
	}
	_, err := s.collection.DeleteOne(
		context.Background(),
		scope,
	)
	if err != nil {
		return err
	}
	return nil
}

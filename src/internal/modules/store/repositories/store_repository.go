package stores_repositories

import (
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	"booking-calendar-server-backend/internal/modules/store/models"
	store_scopes "booking-calendar-server-backend/internal/modules/store/scopes"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(dto *store_dto.CreateStoreDto) error
	Update(filter *store_filter.StoreFilter, dto *store_dto.UpdateStoreDto) error
	GetStoreByID(id uint) (*models.Store, error)
	GetStoreByUserID(userID uint) (*models.Store, error)
	GetOne(filter *store_filter.StoreFilter) (*models.Store, error)
	GetMany(filter *store_filter.StoreFilter) ([]*models.Store, error)
	DeleteByID(ID uint) error
	Delete(filter *store_filter.StoreFilter) error
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(
	db *gorm.DB,
) StoreRepository {
	return &storeRepository{
		db: db,
	}
}

func (r *storeRepository) Create(dto *store_dto.CreateStoreDto) error {
	err := r.db.Create(&dto).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *storeRepository) Update(filter *store_filter.StoreFilter, dto *store_dto.UpdateStoreDto) error {
	model := &models.Store{}
	err := r.db.Model(model).Scopes(store_scopes.StoreScopes(filter)).Updates(&dto).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *storeRepository) GetStoreByID(id uint) (*models.Store, error) {
	return r.GetOne(&store_filter.StoreFilter{
		ID: id,
	})
}
func (r *storeRepository) GetStoreByUserID(userID uint) (*models.Store, error) {
	return r.GetOne(&store_filter.StoreFilter{
		UserID: userID,
	})
}
func (r *storeRepository) GetOne(filter *store_filter.StoreFilter) (*models.Store, error) {
	var store *models.Store
	err := r.db.Scopes(store_scopes.StoreScopes(filter)).First(&store).Error
	if err != nil {
		return nil, err
	}
	return store, nil
}
func (r *storeRepository) GetMany(filter *store_filter.StoreFilter) ([]*models.Store, error) {
	var stores []*models.Store
	err := r.db.Scopes(store_scopes.StoreScopes(filter)).Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return stores, nil
}
func (r *storeRepository) DeleteByID(ID uint) error {
	return r.Delete(&store_filter.StoreFilter{ID: ID})
}
func (r *storeRepository) Delete(filter *store_filter.StoreFilter) error {
	err := r.db.Scopes(store_scopes.StoreScopes(filter)).Delete(&models.Store{}).Error
	if err != nil {
		return err
	}
	return nil
}

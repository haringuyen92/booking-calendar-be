package store_services

import (
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	store_mappers "booking-calendar-server-backend/internal/modules/store/mappers"
	stores_repositories "booking-calendar-server-backend/internal/modules/store/repositories"
	store_responses "booking-calendar-server-backend/internal/modules/store/responses"
)

type StoreService interface {
	GetOne(filter *store_filter.StoreFilter) (*store_responses.GetStoreResponse, error)
	GetMany(filter *store_filter.StoreFilter) ([]*store_responses.GetAllStoreResponse, error)
	Create(dto *store_dto.CreateStoreDto) error
	Update(filter *store_filter.StoreFilter, dto *store_dto.UpdateStoreDto) error
	Delete(filter *store_filter.StoreFilter) error
}

type storeService struct {
	storeRepository stores_repositories.StoreRepository
}

func NewStoreService(
	storeRepository stores_repositories.StoreRepository,
) StoreService {
	return &storeService{
		storeRepository: storeRepository,
	}
}

func (s *storeService) GetOne(filter *store_filter.StoreFilter) (*store_responses.GetStoreResponse, error) {
	store, err := s.storeRepository.GetOne(filter)
	if err != nil {
		return nil, err
	}
	return store_mappers.GetStoreResponseMapper(store), nil
}

func (s *storeService) GetMany(filter *store_filter.StoreFilter) ([]*store_responses.GetAllStoreResponse, error) {
	stores, err := s.storeRepository.GetMany(filter)
	if err != nil {
		return nil, err
	}
	return store_mappers.GetAllStoreResponseMapper(stores), nil
}

func (s *storeService) Create(dto *store_dto.CreateStoreDto) error {
	err := s.storeRepository.Create(dto)
	if err != nil {
		return err
	}
	return nil
}
func (s *storeService) Update(filter *store_filter.StoreFilter, dto *store_dto.UpdateStoreDto) error {
	err := s.storeRepository.Update(filter, dto)
	if err != nil {
		return err
	}
	return nil
}
func (s *storeService) Delete(filter *store_filter.StoreFilter) error {
	err := s.storeRepository.Delete(filter)
	if err != nil {
		return err
	}
	return nil
}

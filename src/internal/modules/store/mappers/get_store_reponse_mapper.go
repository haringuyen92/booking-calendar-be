package store_mappers

import (
	"booking-calendar-server-backend/internal/modules/store/models"
	store_responses "booking-calendar-server-backend/internal/modules/store/responses"
)

func GetStoreResponseMapper(store *models.Store) *store_responses.GetStoreResponse {
	return &store_responses.GetStoreResponse{
		ID: store.ID,
	}
}

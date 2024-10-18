package store_mappers

import (
	"booking-calendar-server-backend/internal/modules/store/models"
	store_responses "booking-calendar-server-backend/internal/modules/store/responses"
)

func GetAllStoreResponseMapper(stores []*models.Store) []*store_responses.GetAllStoreResponse {
	var response []*store_responses.GetAllStoreResponse
	for _, store := range stores {
		response = append(response, &store_responses.GetAllStoreResponse{
			ID: store.ID,
		})
	}
	return response
}

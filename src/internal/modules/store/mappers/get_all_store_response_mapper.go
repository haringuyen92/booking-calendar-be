package store_mappers

import (
	"booking-calendar-server-backend/internal/modules/store/models"
	store_responses "booking-calendar-server-backend/internal/modules/store/responses"
)

func GetAllStoreResponseMapper(stores []*models.Store) []*store_responses.GetAllStoreResponse {
	var response []*store_responses.GetAllStoreResponse
	for _, store := range stores {
		response = append(response, &store_responses.GetAllStoreResponse{
			ID:          store.ID,
			Name:        store.Name,
			Description: store.Description,
			Address:     store.Address,
			Website:     store.Website,
			Logo:        store.Logo,
			Email:       store.Email,
			Phone:       store.Phone,
			Location:    store.Location,
			Status:      store.Status,
		})
	}
	return response
}

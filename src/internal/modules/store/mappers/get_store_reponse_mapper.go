package store_mappers

import (
	"booking-calendar-server-backend/internal/modules/store/models"
	store_responses "booking-calendar-server-backend/internal/modules/store/responses"
)

// GetStoreResponseMapper TODO: mapping setting or other service of store
func GetStoreResponseMapper(store *models.Store) *store_responses.GetStoreResponse {
	return &store_responses.GetStoreResponse{
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
	}
}

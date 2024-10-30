package store_responses

import "booking-calendar-server-backend/internal/core/enums"

type GetAllStoreResponse struct {
	ID          uint              `json:"id"`
	UserID      uint              `json:"user_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Address     string            `json:"address"`
	Website     string            `json:"website"`
	Logo        string            `json:"logo"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	Location    string            `json:"location"`
	Status      enums.StoreStatus `json:"status"`
}

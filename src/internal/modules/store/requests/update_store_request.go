package store_requests

import "booking-calendar-server-backend/internal/core/enums"

type UpdateStoreRequest struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Address     string            `json:"address,omitempty"`
	Website     string            `json:"website,omitempty"`
	Logo        string            `json:"logo,omitempty"`
	Email       string            `json:"email,omitempty"`
	Phone       string            `json:"phone,omitempty"`
	Location    string            `json:"location,omitempty"`
	Status      enums.StoreStatus `json:"status,omitempty"`
}

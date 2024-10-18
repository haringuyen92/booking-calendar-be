package store_requests

type CreateStoreRequest struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

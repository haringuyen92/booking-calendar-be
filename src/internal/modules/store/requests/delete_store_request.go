package store_requests

type DeleteStoreRequest struct {
	ID uint `json:"id" binding:"required"`
}

package requests

type GetAllBookingRequest struct {
	UserID  uint `json:"user_id"`
	StoreID uint `json:"store_id"`
}

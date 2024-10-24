package user_requests

type UpdateUserRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

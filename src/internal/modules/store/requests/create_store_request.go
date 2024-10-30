package store_requests

type CreateStoreRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	Logo        string `json:"logo"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Location    string `json:"location"`
}

package store_dto

type CreateStoreDto struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

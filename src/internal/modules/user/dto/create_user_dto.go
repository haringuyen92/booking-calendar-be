package user_dto

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
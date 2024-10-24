package user_mappers

import (
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_responses "booking-calendar-server-backend/internal/modules/user/responses"
)

func GetUserResponseMapper(user *user_models.User) *user_responses.GetUserResponse {
	return &user_responses.GetUserResponse{}
}

package user_mappers

import (
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_responses "booking-calendar-server-backend/internal/modules/user/responses"
)

func GetAllUserResponse(users []*user_models.User) []*user_responses.GetAllUserResponse {
	var userResponses []*user_responses.GetAllUserResponse
	for _, user := range users {
		userResponses = append(userResponses, &user_responses.GetAllUserResponse{
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return userResponses
}

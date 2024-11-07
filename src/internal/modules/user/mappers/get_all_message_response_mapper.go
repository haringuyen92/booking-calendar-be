package user_mappers

import (
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_responses "booking-calendar-server-backend/internal/modules/user/responses"
)

func GetAllMessageResponse(messages []*user_models.Message) []*user_responses.GetAllMessageResponse {
	var response []*user_responses.GetAllMessageResponse
	for _, message := range messages {
		response = append(response, &user_responses.GetAllMessageResponse{
			ID:             message.ID,
			ConversationID: message.ConversationID,
			SenderID:       message.SenderID,
			Content:        message.Content,
			IsRead:         message.IsRead,
			CreatedAt:      message.CreatedAt,
		})
	}
	return response
}

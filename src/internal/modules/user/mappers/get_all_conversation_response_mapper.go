package user_mappers

import (
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_responses "booking-calendar-server-backend/internal/modules/user/responses"
)

func GetAllConversationResponse(conversations []*user_models.Conversation) []*user_responses.GetAllConversationsResponse {
	var conversationResponses []*user_responses.GetAllConversationsResponse
	for _, conversation := range conversations {
		conversationResponses = append(conversationResponses, &user_responses.GetAllConversationsResponse{
			ID:           conversation.ID,
			Participants: getParticipantsResponse(conversation.Participants),
			CreatedAt:    conversation.CreatedAt,
			UpdatedAt:    conversation.UpdatedAt,
			LastMessage:  getLastMessageResponse(conversation.LastMessage),
		})
	}
	return conversationResponses
}

func getLastMessageResponse(lastMessage *user_models.ConversationLastMessage) user_responses.ConversationLastMessage {
	if lastMessage == nil {
		return user_responses.ConversationLastMessage{}
	}
	return user_responses.ConversationLastMessage{
		SenderID: lastMessage.SenderID,
		Content:  lastMessage.Content,
		SendAt:   lastMessage.SendAt,
	}
}

func getParticipantsResponse(users map[uint]user_models.ConversationUser) []user_responses.ConversationUser {
	var conversationUser []user_responses.ConversationUser
	for userID, user := range users {
		conversationUser = append(conversationUser, user_responses.ConversationUser{
			UserID: userID,
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		})
	}
	return conversationUser
}

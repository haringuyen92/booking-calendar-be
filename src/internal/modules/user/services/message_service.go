package user_services

import (
	user_dto "booking-calendar-server-backend/internal/modules/user/dto"
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	user_mappers "booking-calendar-server-backend/internal/modules/user/mappers"
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_repositories "booking-calendar-server-backend/internal/modules/user/repositories"
	user_responses "booking-calendar-server-backend/internal/modules/user/responses"
)

type MessageService interface {
	GetConversations(filter *user_filters.ConversationFilter) ([]*user_responses.GetAllConversationsResponse, error)
	CreateConversation(dto *user_dto.CreateConversationDto) error
	GetMessages(filter *user_filters.MessageFilter) ([]*user_responses.GetAllMessageResponse, error)
}

func NewMessageService(
	userRepository user_repositories.UserRepository,
	messageRepository user_repositories.MessageRepository,
	conversationRepository user_repositories.ConversationRepository,
) MessageService {
	return &messageService{
		userRepository:         userRepository,
		messageRepository:      messageRepository,
		conversationRepository: conversationRepository,
	}
}

type messageService struct {
	userRepository         user_repositories.UserRepository
	messageRepository      user_repositories.MessageRepository
	conversationRepository user_repositories.ConversationRepository
}

func (s *messageService) GetConversations(filter *user_filters.ConversationFilter) ([]*user_responses.GetAllConversationsResponse, error) {
	conversations, err := s.conversationRepository.GetByUserID(filter.UserID)
	if err != nil {
		return nil, err
	}
	return user_mappers.GetAllConversationResponse(conversations), nil
}

func (s *messageService) CreateConversation(dto *user_dto.CreateConversationDto) error {
	participants := make(map[uint]user_models.ConversationUser)
	users := make([]*user_models.User, 2)
	sender, err := s.userRepository.GetOneByID(dto.SenderID)
	if err != nil {
		return err
	}
	users[0] = sender
	participants[sender.ID] = user_models.ConversationUser{
		Name: sender.Name,
	}
	receiver, err := s.userRepository.GetOneByID(dto.ReceiverID)
	if err != nil {
		return err
	}
	users[1] = receiver
	for _, user := range users {
		participants[user.ID] = user_models.ConversationUser{
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		}
	}
	return s.conversationRepository.Create(&user_models.Conversation{
		Participants: participants,
	})
}

func (s *messageService) GetMessages(filter *user_filters.MessageFilter) ([]*user_responses.GetAllMessageResponse, error) {
	messages, err := s.messageRepository.GetByConversationID(filter.ConversationID)
	if err != nil {
		return nil, err
	}
	return user_mappers.GetAllMessageResponse(messages), nil
}

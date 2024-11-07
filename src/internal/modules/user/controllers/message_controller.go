package user_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	user_dto "booking-calendar-server-backend/internal/modules/user/dto"
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	user_requests "booking-calendar-server-backend/internal/modules/user/requests"
	user_services "booking-calendar-server-backend/internal/modules/user/services"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageService user_services.MessageService
}

func NewMessageController(
	messageService user_services.MessageService,
) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

func (co *MessageController) GetConversations(c *gin.Context, req *user_requests.GetAllConversationRequest) error {
	res, err := co.messageService.GetConversations(&user_filters.ConversationFilter{
		UserID: req.UserID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *MessageController) CreateConversation(c *gin.Context, req *user_requests.CreateConversationRequest) error {
	err := co.messageService.CreateConversation(&user_dto.CreateConversationDto{
		SenderID:   1,
		ReceiverID: 2,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *MessageController) GetMessages(c *gin.Context, req *user_requests.GetAllMessageRequest) error {
	res, err := co.messageService.GetMessages(&user_filters.MessageFilter{
		ConversationID: req.ConversationID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

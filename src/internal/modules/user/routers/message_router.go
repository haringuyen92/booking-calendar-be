package user_routers

import (
	user_controllers "booking-calendar-server-backend/internal/modules/user/controllers"
	user_requests "booking-calendar-server-backend/internal/modules/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageRouter struct {
	h *user_controllers.MessageController
}

func NewMessageRouter(
	h *user_controllers.MessageController,
) *MessageRouter {
	return &MessageRouter{h: h}
}

func (s *MessageRouter) Register(group *gin.RouterGroup) {
	group.GET("/conversations/", s.getConversations())
	group.POST("/conversations/", s.createConversation())
	group.GET("/messages", s.getMessages())
}

func (s *MessageRouter) getConversations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user_requests.GetAllConversationRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		_ = s.h.GetConversations(c, &req)
	}
}

func (s *MessageRouter) getMessages() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user_requests.GetAllMessageRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		_ = s.h.GetMessages(c, &req)
	}
}

func (s *MessageRouter) createConversation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user_requests.CreateConversationRequest
		//if err := c.ShouldBindJSON(&req); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//}
		_ = s.h.CreateConversation(c, &req)
	}
}

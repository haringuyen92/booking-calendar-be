package user_routers

import (
	user_controllers "booking-calendar-server-backend/internal/modules/user/controllers"
	user_requests "booking-calendar-server-backend/internal/modules/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRouter struct {
	h *user_controllers.UserController
}

func NewUserRouter(
	h *user_controllers.UserController,
) *UserRouter {
	return &UserRouter{h: h}
}

func (s *UserRouter) Register(group *gin.RouterGroup) {
	group.GET("", func(c *gin.Context) {
		var req user_requests.GetAllUserRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err := s.h.GetMany(c, &req)
		if err != nil {
			return
		}
	})

	group.GET("/:id", func(c *gin.Context) {
		var req user_requests.GetUserRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := s.h.GetOne(c, &req)
		if err != nil {
			return
		}
	})

	group.POST("/", func(c *gin.Context) {
		var req user_requests.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := s.h.Create(c, &req)
		if err != nil {
			return
		}
	})

	group.PUT("/:id", func(c *gin.Context) {
		var req user_requests.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := s.h.Update(c, &req)
		if err != nil {
			return
		}
	})

	group.DELETE("/:id", func(c *gin.Context) {
		var req *user_requests.UpdateUserRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := s.h.Delete(c, req)
		if err != nil {
			return
		}
	})
}

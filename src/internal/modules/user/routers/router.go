package user_routers

import (
	user_controllers "booking-calendar-server-backend/internal/modules/user/controllers"
	user_requests "booking-calendar-server-backend/internal/modules/user/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(
	r *gin.RouterGroup,
	controller *user_controllers.UserController,
) {
	r.GET("/", func(c *gin.Context) {
		var req *user_requests.GetAllUserRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err := controller.GetMany(c, req)
		if err != nil {
			return
		}
	})

	r.GET("/:id", func(c *gin.Context) {
		var req *user_requests.GetUserRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.GetOne(c, req)
		if err != nil {
			return
		}
	})

	r.POST("/", func(c *gin.Context) {
		var req *user_requests.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.Create(c, req)
		if err != nil {
			return
		}
	})

	r.PUT("/:id", func(c *gin.Context) {
		var req *user_requests.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.Update(c, req)
		if err != nil {
			return
		}
	})

	r.DELETE("/:id", func(c *gin.Context) {
		var req *user_requests.UpdateUserRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.Delete(c, req)
		if err != nil {
			return
		}
	})

}

package gateway_routers

import (
	gateway_controllers "booking-calendar-server-backend/internal/modules/gateway/controllers"
	gateway_requests "booking-calendar-server-backend/internal/modules/gateway/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterGatewayRouters(
	r *gin.RouterGroup,
	authController *gateway_controllers.AuthController,
) {
	r.POST("/login", func(c *gin.Context) {
		var req *gateway_requests.AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := authController.Login(c, req)
		if err != nil {
			return
		}

	})
}

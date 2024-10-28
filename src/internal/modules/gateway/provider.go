package gateway

import (
	gateway_controllers "booking-calendar-server-backend/internal/modules/gateway/controllers"
	gateway_routers "booking-calendar-server-backend/internal/modules/gateway/routers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(
		fx.Provide(gateway_controllers.NewAuthController),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	authController *gateway_controllers.AuthController,
) {
	authGroup := r.Group("/api/auth")
	gateway_routers.RegisterGatewayRouters(authGroup, authController)
}

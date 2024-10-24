package user

import (
	user_controllers "booking-calendar-server-backend/internal/modules/user/controllers"
	user_repositories "booking-calendar-server-backend/internal/modules/user/repositories"
	user_routers "booking-calendar-server-backend/internal/modules/user/routers"
	user_services "booking-calendar-server-backend/internal/modules/user/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(

		fx.Provide(user_repositories.NewUserRepository),
		fx.Provide(user_services.NewUserService),
		fx.Provide(user_controllers.NewUserController),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	controller *user_controllers.UserController,
) {
	userGroup := r.Group("/api/users")
	user_routers.Register(userGroup, controller)
}

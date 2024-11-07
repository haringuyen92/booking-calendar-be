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

		fx.Provide(user_repositories.NewMessageRepository),
		fx.Provide(user_repositories.NewConversationRepository),
		fx.Provide(user_services.NewMessageService),
		fx.Provide(user_controllers.NewMessageController),

		fx.Provide(user_routers.NewUserRouter),
		fx.Provide(user_routers.NewMessageRouter),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	rUser *user_routers.UserRouter,
	rMessage *user_routers.MessageRouter,
) {
	group := r.Group("")
	user_routers.RegisterRouters(
		group,
		rUser,
		rMessage,
	)
}

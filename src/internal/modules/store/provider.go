package store

import (
	store_controllers "booking-calendar-server-backend/internal/modules/store/controllers"
	stores_repositories "booking-calendar-server-backend/internal/modules/store/repositories"
	store_routers "booking-calendar-server-backend/internal/modules/store/routers"
	store_services "booking-calendar-server-backend/internal/modules/store/services"
	store_service "booking-calendar-server-backend/internal/modules/store/services/setting"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(
		fx.Provide(stores_repositories.NewStoreRepository),
		fx.Provide(store_services.NewStoreService),
		fx.Provide(store_controllers.NewStoreController),

		fx.Provide(stores_repositories.NewSettingTimeRepository),
		fx.Provide(store_service.NewSettingStoreService),
		fx.Provide(store_controllers.NewSettingController),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	storeController *store_controllers.StoreController,
	settingStoreController *store_controllers.SettingController,
) {
	storeGroup := r.Group("")
	store_routers.RegisterRouters(
		storeGroup,
		storeController,
		settingStoreController,
	)
}

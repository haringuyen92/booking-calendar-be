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
		fx.Provide(store_controllers.NewSettingController),
		fx.Provide(store_routers.NewSettingRouter),
		fx.Provide(store_routers.NewStoreRouter),

		fx.Provide(stores_repositories.NewSettingTimeRepository),
		fx.Provide(stores_repositories.NewSettingBookingRepository),
		fx.Provide(stores_repositories.NewSettingSlotRepository),
		fx.Provide(store_service.NewSettingStoreService),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	rSetting *store_routers.SettingRouter,
	rStore *store_routers.StoreRouter,
) {
	group := r.Group("")
	store_routers.RegisterRouters(
		group,
		rSetting,
		rStore,
	)
}

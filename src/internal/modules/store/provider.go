package store

import (
	store_controllers "booking-calendar-server-backend/internal/modules/store/controllers"
	stores_repositories "booking-calendar-server-backend/internal/modules/store/repositories"
	store_routers "booking-calendar-server-backend/internal/modules/store/routers"
	store_services "booking-calendar-server-backend/internal/modules/store/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(
		fx.Provide(stores_repositories.NewStoreRepository),
		fx.Provide(store_services.NewStoreService),
		fx.Provide(store_controllers.NewStoreController),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	storeController *store_controllers.StoreController,
) {
	storeGroup := r.Group("/api/stores")
	store_routers.RegisterRouters(storeGroup, storeController)
}

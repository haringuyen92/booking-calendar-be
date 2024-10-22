package staff

import (
	staff_controllers "booking-calendar-server-backend/internal/modules/staff/controllers"
	staff_repositories "booking-calendar-server-backend/internal/modules/staff/repositories"
	staff_routers "booking-calendar-server-backend/internal/modules/staff/routers"
	staff_services "booking-calendar-server-backend/internal/modules/staff/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(
		fx.Provide(staff_repositories.NewStaffRepository),
		fx.Provide(staff_services.NewStaffService),
		fx.Provide(staff_controllers.NewStaffController),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	staffController *staff_controllers.StaffController,
) {
	staffGroup := r.Group("/api/staffs")
	staff_routers.RegisterRouters(staffGroup, staffController)
}

package booking

import (
	booking_controllers "booking-calendar-server-backend/internal/modules/booking/controllers"
	booking_repositories "booking-calendar-server-backend/internal/modules/booking/repositories"
	booking_routers "booking-calendar-server-backend/internal/modules/booking/routers"
	booking_services "booking-calendar-server-backend/internal/modules/booking/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(
		fx.Provide(booking_repositories.NewBookingRepository),
		fx.Provide(booking_services.NewBookingService),
		fx.Provide(booking_controllers.NewBookingController),

		fx.Invoke(registerRoutes),
	)
}

func registerRoutes(
	r *gin.Engine,
	bookingController *booking_controllers.BookingController,
) {
	bookingGroup := r.Group("/api/bookings")
	booking_routers.RegisterBookingRouters(bookingGroup, bookingController)
}

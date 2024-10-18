package main

import (
	booking_controllers "booking-calendar-server-backend/internal/modules/booking/controllers"
	booking_repositories "booking-calendar-server-backend/internal/modules/booking/repositories"
	booking_routers "booking-calendar-server-backend/internal/modules/booking/routers"
	booking_services "booking-calendar-server-backend/internal/modules/booking/services"
	store_controllers "booking-calendar-server-backend/internal/modules/store/controllers"
	stores_repositories "booking-calendar-server-backend/internal/modules/store/repositories"
	store_routers "booking-calendar-server-backend/internal/modules/store/routers"
	store_services "booking-calendar-server-backend/internal/modules/store/services"
	"booking-calendar-server-backend/pkg/boostrap"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func main() {
	fx.New(
		boostrap.All(),
		boostrap.WithMongoDB(),
		boostrap.WithDatabase(),

		fx.Provide(NewGinEngine),
		fx.Provide(NewHTTPServer),

		fx.Provide(booking_repositories.NewBookingRepository),
		fx.Provide(booking_services.NewBookingService),
		fx.Provide(booking_controllers.NewBookingController),

		fx.Provide(stores_repositories.NewStoreRepository),
		fx.Provide(store_services.NewStoreService),
		fx.Provide(store_controllers.NewStoreController),

		fx.Invoke(RegisterRoutes),
		fx.Invoke(func(*http.Server) {}),
	).Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func NewHTTPServer(lc fx.Lifecycle, engine *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Server is now listening on", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func RegisterRoutes(
	r *gin.Engine,
	bookingController *booking_controllers.BookingController,
	storeController *store_controllers.StoreController,
) {
	// BOOKING
	bookingGroup := r.Group("/api/bookings")
	booking_routers.RegisterBookingRouters(bookingGroup, bookingController)

	storeGroup := r.Group("/api/stores")
	store_routers.RegisterRouters(storeGroup, storeController)

}

package main

import (
	"booking-calendar-server-backend/internal/modules/booking/controllers"
	"booking-calendar-server-backend/internal/modules/booking/repositories"
	"booking-calendar-server-backend/internal/modules/booking/routers"
	"booking-calendar-server-backend/internal/modules/booking/services"
	"booking-calendar-server-backend/pkg/boostrap"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func main() {
	fx.New(
		boostrap.All(),
		boostrap.WithMongoDB(),

		fx.Provide(NewGinEngine),
		fx.Provide(NewHTTPServer),

		fx.Provide(repositories.NewBookingRepository),
		fx.Provide(services.NewBookingService),
		fx.Provide(controllers.NewBookingController),

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
	bookingController *controllers.BookingController,
) {
	// BOOKING
	bookingGroup := r.Group("/api/bookings")
	routers.RegisterBookingRouters(bookingGroup, bookingController)
}

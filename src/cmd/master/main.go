package main

import (
	"booking-calendar-server-backend/internal/modules/booking"
	"booking-calendar-server-backend/internal/modules/course"
	"booking-calendar-server-backend/internal/modules/staff"
	"booking-calendar-server-backend/internal/modules/store"
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

		booking.Provider(),
		store.Provider(),
		staff.Provider(),
		course.Provider(),

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

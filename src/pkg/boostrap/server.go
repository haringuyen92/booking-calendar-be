package boostrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"go.uber.org/fx"
	"net/http"
	"time"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Any("/api/stores/*proxyPath", ReverseProxy)
	r.Any("/api/bookings/*proxyPath", ReverseProxy)
	r.Any("/api/users/*proxyPath", ReverseProxy)
	return r
}

func NewHTTPServer(app *golib.App, engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Port()),
		Handler: engine,
	}
}

func OnStartHttpServer(lc fx.Lifecycle, httpServer *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Start http server with address:", httpServer.Addr)
			go func() {
				err := httpServer.ListenAndServe()
				if err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						fmt.Println("Start http server error:", err)
						return
					}
				}
			}()
			return nil
		},
	})
}

func OnStopHttpServer(lc fx.Lifecycle, httpServer *http.Server) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stop http server with address:", httpServer.Addr)
			err := httpServer.Shutdown(ctx)
			if err != nil {
				fmt.Println("Stop http server error:", err)
				return err
			}
			return nil
		},
	})
}

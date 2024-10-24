package boostrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func NewHTTPServer(engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    ":8088",
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

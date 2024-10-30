package boostrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"go.uber.org/fx"
	"net/http"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(SetupCORS())
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

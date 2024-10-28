package main

import (
	"booking-calendar-server-backend/internal/modules/gateway"
	"booking-calendar-server-backend/pkg/boostrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		boostrap.All(),

		fx.Provide(boostrap.NewGinEngine),
		fx.Provide(boostrap.NewHTTPServer),

		gateway.Provider(),

		fx.Invoke(boostrap.OnStartHttpServer),
		fx.Invoke(boostrap.OnStopHttpServer),
		fx.Invoke(boostrap.ReverseProxy),
	).Run()
}

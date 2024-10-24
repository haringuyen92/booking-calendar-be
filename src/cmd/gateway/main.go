package main

import (
	"booking-calendar-server-backend/internal/modules/gateway"
	"booking-calendar-server-backend/pkg/boostrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(boostrap.NewGinEngine),
		fx.Provide(boostrap.NewHTTPServer),

		fx.Provide(gateway.Provider),

		fx.Invoke(boostrap.OnStartHttpServer),
		fx.Invoke(boostrap.OnStopHttpServer),
	).Run()
}

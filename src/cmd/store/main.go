package main

import (
	"booking-calendar-server-backend/internal/modules/course"
	"booking-calendar-server-backend/internal/modules/staff"
	"booking-calendar-server-backend/internal/modules/store"
	"booking-calendar-server-backend/pkg/boostrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		boostrap.All(),
		boostrap.WithMongoDB(),
		boostrap.WithDatabase(),
		boostrap.WithRedis(),

		fx.Provide(boostrap.NewGinEngine),
		fx.Provide(boostrap.NewHTTPServer),

		store.Provider(),
		staff.Provider(),
		course.Provider(),

		fx.Invoke(boostrap.OnStartHttpServer),
		fx.Invoke(boostrap.OnStopHttpServer),
	).Run()
}

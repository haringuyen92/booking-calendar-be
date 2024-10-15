package boostrap

import (
	"github.com/golibs-starter/golib"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		// Required
		golib.AppOpt(),
		golib.PropertiesOpt(),

		// When you want to use default logging strategy.
		golib.LoggingOpt(),
		//
		//// When you want to enable event publisher
		//golib.EventOpt(),
		//// When you want handle event in simple synchronous way
		//golib.SupplyEventBusOpt(pubsub.WithEventExecutor(executor.NewSyncExecutor())),
		//// Or want a custom executor, such as using worker pool
		//
		//// When you want to enable http request log
		//golib.HttpRequestLogOpt(),
		//
		//// When you want to enable actuator endpoints.
		//// By default, we provide HealthService and InfoService.
		//golib.ActuatorEndpointOpt(),
		//
		//// When you want to enable http client auto config with contextual client by default
		//golib.HttpClientOpt(),
		//
		//// Graceful shutdown.
		//// OnStop hooks will run in reverse order.
		//golib.OnStopEventOpt(),
	)
}

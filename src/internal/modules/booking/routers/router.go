package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type RegisterBookingRoutersIn struct {
}

func RegisterBookingRouters(r *gin.Engine, basePath string, controller interface{}) {
	controllerType := reflect.TypeOf(controller)
	controllerValue := reflect.ValueOf(controller)

	for i := 0; i < controllerType.NumMethod(); i++ {
		method := controllerType.Method(i)
		httpMethod := strings.ToUpper(strings.Split(method.Name, "_")[0])
		path := strings.ToLower(strings.Split(method.Name, "_")[1])

		r.Handle(httpMethod, basePath+"/"+path, func(c *gin.Context) {
			// Get the method type
			methodType := method.Type

			// Create a slice to hold the arguments
			args := make([]reflect.Value, methodType.NumIn())
			args[0] = controllerValue    // The receiver (controller instance)
			args[1] = reflect.ValueOf(c) // The gin.Context

			// If the method expects more than 2 arguments, we need to create the request struct
			if methodType.NumIn() > 2 {
				// Get the type of the request struct
				reqType := methodType.In(2)
				// Create a new instance of the request struct
				reqValue := reflect.New(reqType.Elem())

				// Bind the request data to the struct
				if err := c.ShouldBind(reqValue.Interface()); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				args[2] = reqValue
			}

			// Call the method
			results := method.Func.Call(args)

			// Check if an error was returned
			if len(results) > 0 && !results[0].IsNil() {
				err := results[0].Interface().(error)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		})
	}
}

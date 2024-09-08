package interceptors

import (
	"booking-calendar-server-backend/internal/core/base"
	"booking-calendar-server-backend/pkg/errors"
	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, e error) error {
	if _, ok := e.(*errors.Error); ok {
		e := e.(*errors.Error)
		var responseErr base.Response
		responseErr.Code = uint(e.Code)
		responseErr.Message = e.Message
		responseErr.Data = nil

		c.JSON(500, responseErr)
		return e
	}
	c.JSON(500, e)
	return e
}

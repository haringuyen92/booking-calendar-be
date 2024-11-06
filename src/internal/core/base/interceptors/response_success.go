package interceptors

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func ResponseSuccess(c *gin.Context, data interface{}) error {
	response := gin.H{
		"message": "success",
		"code":    200,
		"data":    ensureSlice(data),
	}

	c.JSON(200, response)
	return nil
}

func ensureSlice(data interface{}) interface{} {
	if data == nil {
		return []interface{}{}
	}

	switch v := reflect.ValueOf(data); v.Kind() {
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {
			return []interface{}{}
		}
	}

	return data
}

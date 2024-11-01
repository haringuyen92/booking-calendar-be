package interceptors

import "github.com/gin-gonic/gin"

func ResponseSuccess(c *gin.Context, data interface{}) error {
	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
		"code":    200,
	})
	return nil
}

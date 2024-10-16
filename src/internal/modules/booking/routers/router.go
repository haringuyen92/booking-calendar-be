package routers

import (
	"booking-calendar-server-backend/internal/modules/booking/controllers"
	"booking-calendar-server-backend/internal/modules/booking/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterBookingRouters(group *gin.RouterGroup, bookingController *controllers.BookingController) {
	group.GET("/", func(c *gin.Context) {
		var req requests.GetAllBookingRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := bookingController.GetAll(c, &req)
		if err != nil {
			return
		}
	})

	group.GET("/:id", func(c *gin.Context) {
		var req requests.GetBookingRequest
		req.ID = c.Param("id")

		err := bookingController.GetOne(c, &req)
		if err != nil {
			return
		}
	})

	group.POST("", func(c *gin.Context) {
		var req requests.CreateBookingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := bookingController.Create(c, &req)
		if err != nil {
			// Lỗi đã được xử lý trong controller
			return
		}
	})

	group.PUT("/:id", func(c *gin.Context) {
		var req requests.UpdateBookingRequest
		req.ID = c.Param("id")
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := bookingController.Update(c, &req)
		if err != nil {
			return
		}
	})

	group.DELETE("/:id", func(c *gin.Context) {
		var req requests.DeleteBookingRequest
		req.ID = c.Param("id")
		err := bookingController.Delete(c, &req)
		if err != nil {
			return
		}
	})
}

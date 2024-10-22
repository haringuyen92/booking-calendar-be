package course_routers

import (
	course_controllers "booking-calendar-server-backend/internal/modules/course/controllers"
	course_requests "booking-calendar-server-backend/internal/modules/course/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(
	group *gin.RouterGroup,
	controller *course_controllers.CourseController,
) {
	group.GET("/", func(c *gin.Context) {
		var req course_requests.GetAllCourseRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.GetAll(c, &req)
		if err != nil {
			return
		}
	})

	group.GET("/:id", func(c *gin.Context) {
		var req course_requests.GetCourseRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		req.ID = uint(id)

		err = controller.GetOne(c, &req)
		if err != nil {
			return
		}
	})

	group.POST("", func(c *gin.Context) {
		var req course_requests.CreateCourseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.Create(c, &req)
		if err != nil {
			// Lỗi đã được xử lý trong controller
			return
		}
	})

	group.PUT("/:id", func(c *gin.Context) {
		var req course_requests.UpdateCourseRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		req.ID = uint(id)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = controller.Update(c, &req)
		if err != nil {
			return
		}
	})

	group.DELETE("/:id", func(c *gin.Context) {
		var req course_requests.DeleteCourseRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		req.ID = uint(id)
		err = controller.Delete(c, &req)
		if err != nil {
			return
		}
	})
}

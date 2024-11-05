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
	group.GET("", func(c *gin.Context) {
		var req course_requests.GetAllCourseRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		req.StoreID = uint(storeId)
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_ = controller.GetAll(c, &req)
	})

	group.GET("/:courseId", func(c *gin.Context) {
		var req course_requests.GetCourseRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		id, err := strconv.Atoi(c.Param("courseId"))
		if err != nil {
			return
		}
		req.ID = uint(id)
		req.StoreID = uint(storeId)
		_ = controller.GetOne(c, &req)
	})

	group.POST("", func(c *gin.Context) {
		var req course_requests.CreateCourseRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		req.StoreID = uint(storeId)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_ = controller.Create(c, &req)
	})

	group.PUT("/:courseId", func(c *gin.Context) {
		var req course_requests.UpdateCourseRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		id, err := strconv.Atoi(c.Param("courseId"))
		if err != nil {
			return
		}
		req.ID = uint(id)
		req.StoreID = uint(storeId)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_ = controller.Update(c, &req)
	})

	group.DELETE("/:courseId", func(c *gin.Context) {
		var req course_requests.DeleteCourseRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		req.ID = uint(id)
		req.StoreID = uint(storeId)
		_ = controller.Delete(c, &req)
	})
}

package staff_routers

import (
	staff_controllers "booking-calendar-server-backend/internal/modules/staff/controllers"
	staff_requests "booking-calendar-server-backend/internal/modules/staff/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(group *gin.RouterGroup, controller *staff_controllers.StaffController) {
	group.GET("", func(c *gin.Context) {
		var req staff_requests.GetAllStaffRequest
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

	group.GET("/:staffId", func(c *gin.Context) {
		var req staff_requests.GetStaffRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		id, err := strconv.Atoi(c.Param("staffId"))
		if err != nil {
			return
		}
		req.StoreID = uint(storeId)
		req.ID = uint(id)

		_ = controller.GetOne(c, &req)
	})

	group.POST("", func(c *gin.Context) {
		var req staff_requests.CreateStaffRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.StoreID = uint(storeId)
		_ = controller.Create(c, &req)
	})

	group.PUT("/:staffId", func(c *gin.Context) {
		var req staff_requests.UpdateStaffRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		id, err := strconv.Atoi(c.Param("staffId"))
		if err != nil {
			return
		}
		req.StoreID = uint(storeId)
		req.ID = uint(id)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_ = controller.Update(c, &req)
	})

	group.DELETE("/:staffId", func(c *gin.Context) {
		var req staff_requests.DeleteStaffRequest
		storeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		id, err := strconv.Atoi(c.Param("staffId"))
		if err != nil {
			return
		}
		req.ID = uint(id)
		req.StoreID = uint(storeId)
		_ = controller.Delete(c, &req)
	})
}

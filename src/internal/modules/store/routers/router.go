package store_routers

import (
	store_controllers "booking-calendar-server-backend/internal/modules/store/controllers"
	store_requests "booking-calendar-server-backend/internal/modules/store/requests"
	store_requests2 "booking-calendar-server-backend/internal/modules/store/requests/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterRouters(
	group *gin.RouterGroup,
	controller *store_controllers.StoreController,
	settingStoreController *store_controllers.SettingController,
) {
	group.GET("/:id/setting-time", func(c *gin.Context) {
		var req store_requests2.GetSettingTimeRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}
		req.StoreID = uint(id)
		err = settingStoreController.GetSettingTime(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	})

	group.PUT("/:id/setting-time", func(c *gin.Context) {
		var req store_requests2.UpdateSettingTimeRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		req.StoreID = uint(id)
		err = settingStoreController.UpdateSettingTime(c, &req)
		fmt.Println("req.IsOpen", req.IsOpen)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	})
	group.GET("", func(c *gin.Context) {
		var req store_requests.GetAllStoreRequest
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
		var req store_requests.GetStoreRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		req.ID = uint(id)

		err = controller.GetOne(c, &req)
		if err != nil {
			return
		}
	})

	group.POST("", func(c *gin.Context) {
		var req store_requests.CreateStoreRequest
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
		var req store_requests.UpdateStoreRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}

		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.ID = uint(id)

		err = controller.Update(c, &req)
		if err != nil {
			return
		}
	})

	group.DELETE("/:id", func(c *gin.Context) {
		var req store_requests.DeleteStoreRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}

		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		req.ID = uint(id)
		err = controller.Delete(c, &req)
		if err != nil {
			return
		}
	})
}

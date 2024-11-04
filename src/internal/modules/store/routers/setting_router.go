package store_routers

import (
	store_controllers "booking-calendar-server-backend/internal/modules/store/controllers"
	store_requests2 "booking-calendar-server-backend/internal/modules/store/requests/settings"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SettingRouter struct {
	h *store_controllers.SettingController
}

func NewSettingRouter(
	h *store_controllers.SettingController,
) *SettingRouter {
	return &SettingRouter{
		h: h,
	}
}

func (s *SettingRouter) Register(group *gin.RouterGroup) {
	group.GET("/:id/setting-time", s.getSettingTime())
	group.PUT("/:id/setting-time", s.updateSettingTime())
	group.GET("/:id/setting-booking", s.getSettingBooking())
	group.PUT("/:id/setting-booking", s.updateSettingBooking())
	group.GET("/:id/setting-slot", s.getSettingSlot())
	group.PUT("/:id/setting-slot", s.updateSettingSlot())
}

func (s *SettingRouter) getSettingTime() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		_ = s.h.GetSettingTime(c, &req)
	}
}

func (s *SettingRouter) updateSettingTime() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		_ = s.h.UpdateSettingTime(c, &req)
	}
}

func (s *SettingRouter) getSettingBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req store_requests2.GetSettingBookingRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}
		req.StoreID = uint(id)
		_ = s.h.GetSettingBooking(c, &req)
	}
}

func (s *SettingRouter) updateSettingBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req store_requests2.UpdateSettingBookingRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.StoreID = uint(id)
		_ = s.h.UpdateSettingBooking(c, &req)
	}
}

func (s *SettingRouter) getSettingSlot() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req store_requests2.GetSettingSlotRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}
		req.StoreID = uint(id)
		err = s.h.GetSettingSlot(c, &req)
	}
}

func (s *SettingRouter) updateSettingSlot() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req store_requests2.UpdateSettingSlotRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}
		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.StoreID = uint(id)
		_ = s.h.UpdateSettingSlot(c, &req)
	}
}

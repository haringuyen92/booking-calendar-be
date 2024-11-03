package store_routers

import (
	store_controllers "booking-calendar-server-backend/internal/modules/store/controllers"
	store_requests "booking-calendar-server-backend/internal/modules/store/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type StoreRouter struct {
	h *store_controllers.StoreController
}

func NewStoreRouter(
	h *store_controllers.StoreController,
) *StoreRouter {
	return &StoreRouter{
		h: h,
	}
}

func (s *StoreRouter) Register(group *gin.RouterGroup) {
	group.GET("", s.getListStore())
	group.GET("/:id", s.getStoreByID())
	group.POST("", s.createStore())
	group.PUT("/:id", s.updateStoreByID())
	group.DELETE("/:id", s.deleteStoreByID())
}

func (s *StoreRouter) getListStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req store_requests.GetAllStoreRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := s.h.GetAll(c, &req)
		if err != nil {
			return
		}
	}
}

func (s *StoreRouter) getStoreByID() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		err = s.h.GetOne(c, &req)
		if err != nil {
			return
		}
	}
}

func (s *StoreRouter) createStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req store_requests.CreateStoreRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := s.h.Create(c, &req)
		if err != nil {
			// Lỗi đã được xử lý trong controller
			return
		}
	}
}

func (s *StoreRouter) updateStoreByID() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		err = s.h.Update(c, &req)
		if err != nil {
			return
		}
	}
}

func (s *StoreRouter) deleteStoreByID() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		err = s.h.Delete(c, &req)
		if err != nil {
			return
		}
	}
}

package gateway_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	gateway_requests "booking-calendar-server-backend/internal/modules/gateway/requests"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (co *AuthController) Login(ctx *gin.Context, req *gateway_requests.AuthRequest) error {
	fmt.Println("Login with jwt:", req.Code)
	return interceptors.ResponseSuccess(ctx, nil)
}

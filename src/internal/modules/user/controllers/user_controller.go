package user_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	user_dto "booking-calendar-server-backend/internal/modules/user/dto"
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	user_requests "booking-calendar-server-backend/internal/modules/user/requests"
	user_services "booking-calendar-server-backend/internal/modules/user/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user_services.UserService
}

func NewUserController(
	userService user_services.UserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (co *UserController) GetOne(c *gin.Context, req *user_requests.GetUserRequest) error {
	res, err := co.userService.GetOne(&user_filters.UserFilter{
		ID: req.ID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *UserController) GetMany(c *gin.Context, req *user_requests.GetAllUserRequest) error {
	res, err := co.userService.GetMany(&user_filters.UserFilter{})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *UserController) Create(c *gin.Context, req *user_requests.CreateUserRequest) error {
	err := co.userService.Create(&user_dto.CreateUserDTO{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *UserController) Update(c *gin.Context, req *user_requests.UpdateUserRequest) error {
	err := co.userService.Update(
		&user_filters.UserFilter{
			ID: req.ID,
		},
		&user_dto.UpdateUserDTO{
			Name: req.Name,
		},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *UserController) Delete(c *gin.Context, req *user_requests.UpdateUserRequest) error {
	err := co.userService.Delete(&user_filters.UserFilter{
		ID: req.ID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

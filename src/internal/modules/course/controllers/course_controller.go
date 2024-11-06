package course_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	course_dto "booking-calendar-server-backend/internal/modules/course/dto"
	course_filters "booking-calendar-server-backend/internal/modules/course/filters"
	course_requests "booking-calendar-server-backend/internal/modules/course/requests"
	course_services "booking-calendar-server-backend/internal/modules/course/services"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	courseService course_services.CourseService
}

func NewCourseController(
	courseService course_services.CourseService,
) *CourseController {
	return &CourseController{
		courseService: courseService,
	}
}

func (c *CourseController) Create(ctx *gin.Context, req *course_requests.CreateCourseRequest) error {
	err := c.courseService.Create(&course_dto.CreateCourseDto{
		StoreID:      req.StoreID,
		Name:         req.Name,
		Image:        req.Image,
		Description:  req.Description,
		Cost:         req.Cost,
		EstimateTime: req.EstimateTime,
		Position:     req.Position,
	})

	if err != nil {
		return interceptors.ResponseError(ctx, err)
	}

	return interceptors.ResponseSuccess(ctx, nil)
}

func (c *CourseController) Update(ctx *gin.Context, req *course_requests.UpdateCourseRequest) error {
	err := c.courseService.Update(
		&course_filters.CourseFilter{
			ID:      req.ID,
			StoreID: req.StoreID,
		}, &course_dto.UpdateCourseDto{
			Name:         req.Name,
			Image:        req.Image,
			Description:  req.Description,
			Cost:         req.Cost,
			Active:       req.Active,
			EstimateTime: req.EstimateTime,
			Position:     req.Position,
		})

	if err != nil {
		return interceptors.ResponseError(ctx, err)
	}

	return interceptors.ResponseSuccess(ctx, nil)
}

func (c *CourseController) Delete(ctx *gin.Context, req *course_requests.DeleteCourseRequest) error {
	err := c.courseService.Delete(&course_filters.CourseFilter{
		ID:      req.ID,
		StoreID: req.StoreID,
	})

	if err != nil {
		return interceptors.ResponseError(ctx, err)
	}

	return interceptors.ResponseSuccess(ctx, nil)
}

func (c *CourseController) GetOne(ctx *gin.Context, request *course_requests.GetCourseRequest) error {
	response, err := c.courseService.GetOne(&course_filters.CourseFilter{
		ID: request.ID,
	})

	if err != nil {
		return interceptors.ResponseError(ctx, err)
	}

	return interceptors.ResponseSuccess(ctx, response)
}

func (c *CourseController) GetAll(ctx *gin.Context, request *course_requests.GetAllCourseRequest) error {
	response, err := c.courseService.GetMany(&course_filters.CourseFilter{
		StoreID: request.StoreID,
	})

	if err != nil {
		return interceptors.ResponseError(ctx, err)
	}

	return interceptors.ResponseSuccess(ctx, response)
}

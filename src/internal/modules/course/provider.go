package course

import (
	course_controllers "booking-calendar-server-backend/internal/modules/course/controllers"
	course_repositories "booking-calendar-server-backend/internal/modules/course/repositories"
	course_routers "booking-calendar-server-backend/internal/modules/course/routers"
	course_services "booking-calendar-server-backend/internal/modules/course/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(
		fx.Provide(course_repositories.NewCourseRepository),
		fx.Provide(course_services.NewCourseService),
		fx.Provide(course_controllers.NewCourseController),

		fx.Invoke(registerRoutes),
	)
}
func registerRoutes(
	r *gin.Engine,
	courseController *course_controllers.CourseController,
) {
	staffGroup := r.Group("/api/courses")
	course_routers.RegisterRouters(staffGroup, courseController)
}

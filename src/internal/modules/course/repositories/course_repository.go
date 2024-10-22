package course_repositories

import (
	course_dto "booking-calendar-server-backend/internal/modules/course/dto"
	course_filters "booking-calendar-server-backend/internal/modules/course/filters"
	course_models "booking-calendar-server-backend/internal/modules/course/models"
	course_scopes "booking-calendar-server-backend/internal/modules/course/scopes"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(dto *course_dto.CreateCourseDto) error
	Update(filter *course_filters.CourseFilter, dto *course_dto.UpdateCourseDto) error
	GetOne(filter *course_filters.CourseFilter) (*course_models.Course, error)
	GetMany(filter *course_filters.CourseFilter) ([]*course_models.Course, error)
	DeleteByID(ID uint) error
	Delete(filter *course_filters.CourseFilter) error
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

type courseRepository struct {
	db *gorm.DB
}

func (r *courseRepository) Create(dto *course_dto.CreateCourseDto) error {
	course := &course_models.Course{
		Name: dto.Name,
	}

	return r.db.Create(course).Error
}

func (r *courseRepository) Update(filter *course_filters.CourseFilter, dto *course_dto.UpdateCourseDto) error {
	course := &course_models.Course{
		Name: dto.Name,
	}

	return r.db.Model(course).Where(filter).Updates(course).Error
}

func (r *courseRepository) GetOne(filter *course_filters.CourseFilter) (*course_models.Course, error) {
	var course course_models.Course

	return &course, r.db.Scopes(course_scopes.CourseScope(filter)).First(&course).Error
}

func (r *courseRepository) GetMany(filter *course_filters.CourseFilter) ([]*course_models.Course, error) {
	var courses []*course_models.Course

	return courses, r.db.Scopes(course_scopes.CourseScope(filter)).Find(&courses).Error
}

func (r *courseRepository) DeleteByID(ID uint) error {
	return r.db.Delete(&course_models.Course{}, ID).Error
}

func (r *courseRepository) Delete(filter *course_filters.CourseFilter) error {
	return r.db.Delete(&course_models.Course{}, filter).Error
}

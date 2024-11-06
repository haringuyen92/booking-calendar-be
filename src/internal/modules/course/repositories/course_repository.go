package course_repositories

import (
	"booking-calendar-server-backend/internal/core/enums"
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
		StoreID:      dto.StoreID,
		Name:         dto.Name,
		Image:        dto.Image,
		Description:  dto.Description,
		Cost:         dto.Cost,
		EstimateTime: dto.EstimateTime,
		Active:       enums.COURSE_INACTIVE,
		Position:     dto.Position,
	}

	return r.db.Create(course).Error
}

func (r *courseRepository) Update(filter *course_filters.CourseFilter, dto *course_dto.UpdateCourseDto) error {
	course := &course_models.Course{}

	return r.db.Model(course).Where(filter).Updates(&dto).Error
}

func (r *courseRepository) GetOne(filter *course_filters.CourseFilter) (*course_models.Course, error) {
	var course course_models.Course

	return &course, r.db.Scopes(course_scopes.CourseScope(filter)).First(&course).Error
}

func (r *courseRepository) GetMany(filter *course_filters.CourseFilter) ([]*course_models.Course, error) {
	var courses []*course_models.Course

	err := r.db.Scopes(course_scopes.CourseScope(filter)).Find(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *courseRepository) DeleteByID(ID uint) error {
	return r.db.Delete(&course_models.Course{}, ID).Error
}

func (r *courseRepository) Delete(filter *course_filters.CourseFilter) error {
	return r.db.Scopes(course_scopes.CourseScope(filter)).Delete(&course_models.Course{}, filter).Error
}

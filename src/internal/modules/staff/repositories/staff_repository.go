package staff_repositories

import (
	"booking-calendar-server-backend/internal/core/enums"
	staff_dto "booking-calendar-server-backend/internal/modules/staff/dto"
	staff_filter "booking-calendar-server-backend/internal/modules/staff/filters"
	"booking-calendar-server-backend/internal/modules/staff/models"
	staff_scopes "booking-calendar-server-backend/internal/modules/staff/scopes"

	"gorm.io/gorm"
)

type StaffRepository interface {
	Create(dto *staff_dto.CreateStaffDto) error
	Update(filter *staff_filter.StaffFilter, dto *staff_dto.UpdateStaffDto) error
	GetStaffByID(id uint) (*staff_models.Staff, error)
	GetStaffByUserID(userID uint) (*staff_models.Staff, error)
	GetOne(filter *staff_filter.StaffFilter) (*staff_models.Staff, error)
	GetMany(filter *staff_filter.StaffFilter) ([]*staff_models.Staff, error)
	DeleteByID(ID uint) error
	Delete(filter *staff_filter.StaffFilter) error
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(
	db *gorm.DB,
) StaffRepository {
	return &staffRepository{
		db: db,
	}
}

func (r *staffRepository) Create(dto *staff_dto.CreateStaffDto) error {
	err := r.db.Create(&staff_models.Staff{
		StoreID:        dto.StoreID,
		Name:           dto.Name,
		Email:          dto.Email,
		Phone:          dto.Phone,
		Cost:           dto.Cost,
		MaxBookingSlot: dto.MaxBookingSlot,
		Active:         enums.STAFF_INACTIVE,
		Color:          dto.Color,
		Position:       dto.Position,
		IsAllCourse:    dto.IsAllCourse,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *staffRepository) Update(filter *staff_filter.StaffFilter, dto *staff_dto.UpdateStaffDto) error {
	model := &staff_models.Staff{}
	err := r.db.Model(model).Scopes(staff_scopes.StaffScopes(filter)).Updates(&dto).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *staffRepository) GetStaffByID(id uint) (*staff_models.Staff, error) {
	return r.GetOne(&staff_filter.StaffFilter{
		ID: id,
	})
}
func (r *staffRepository) GetStaffByUserID(userID uint) (*staff_models.Staff, error) {
	return r.GetOne(&staff_filter.StaffFilter{
		UserID: userID,
	})
}
func (r *staffRepository) GetOne(filter *staff_filter.StaffFilter) (*staff_models.Staff, error) {
	var staff *staff_models.Staff
	err := r.db.Scopes(staff_scopes.StaffScopes(filter)).First(&staff).Error
	if err != nil {
		return nil, err
	}
	return staff, nil
}
func (r *staffRepository) GetMany(filter *staff_filter.StaffFilter) ([]*staff_models.Staff, error) {
	var staffs []*staff_models.Staff
	err := r.db.Scopes(staff_scopes.StaffScopes(filter)).Find(&staffs).Error
	if err != nil {
		return nil, err
	}
	return staffs, nil
}
func (r *staffRepository) DeleteByID(ID uint) error {
	return r.Delete(&staff_filter.StaffFilter{ID: ID})
}
func (r *staffRepository) Delete(filter *staff_filter.StaffFilter) error {
	err := r.db.Scopes(staff_scopes.StaffScopes(filter)).Delete(&staff_models.Staff{}).Error
	if err != nil {
		return err
	}
	return nil
}

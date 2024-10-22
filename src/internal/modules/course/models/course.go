package course_models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name string `json:"name"`
}

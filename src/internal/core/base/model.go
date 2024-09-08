package base

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint         `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt sql.NullTime `json:"-"`
}

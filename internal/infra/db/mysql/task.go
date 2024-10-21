package mysql

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Status    uint16    `gorm:"column:status;not null"`
	UserID    uuid.UUID `gorm:"column:user_id;type:char(36);not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (Task) TableName() string {
	return "tasks"
}

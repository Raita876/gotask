package mysql

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (User) TableName() string {
	return "users"
}

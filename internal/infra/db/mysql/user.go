package mysql

import (
	"time"

	"github.com/google/uuid"
	"github.com/raita876/gotask/internal/domain"
)

type DBUser struct {
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (DBUser) TableName() string {
	return "users"
}

func toDBUser(user *domain.User) *DBUser {
	return &DBUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func fromDBUser(dbUser *DBUser) *domain.User {
	return &domain.User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

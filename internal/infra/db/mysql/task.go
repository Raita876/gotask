package mysql

import (
	"time"

	"github.com/google/uuid"
	"github.com/raita876/gotask/internal/domain"
)

type DBTask struct {
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Status    uint16    `gorm:"column:status;not null"`
	UserID    uuid.UUID `gorm:"column:user_id;type:char(36);not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (DBTask) TableName() string {
	return "tasks"
}

func toDBTask(task *domain.Task) *DBTask {
	return &DBTask{
		ID:        task.ID,
		Name:      task.Name,
		Status:    task.Status,
		UserID:    task.UserID,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func fromDBTask(dbTask *DBTask) *domain.Task {
	return &domain.Task{
		ID:        dbTask.ID,
		Name:      dbTask.Name,
		Status:    dbTask.Status,
		UserID:    dbTask.UserID,
		CreatedAt: dbTask.CreatedAt,
		UpdatedAt: dbTask.UpdatedAt,
	}
}

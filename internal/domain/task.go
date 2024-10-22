package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *Task) (*Task, error)
	FindByID(id uuid.UUID) (*Task, error)
	FindByUserID(userID uuid.UUID) ([]*Task, error)
	Update(id uuid.UUID, name string, status uint16) error
	Delete(id uuid.UUID) error
}

type Task struct {
	ID        uuid.UUID
	Name      string
	Status    uint16
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (task *Task) Validate() error {
	// TODO
	return nil
}

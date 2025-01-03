package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByID(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(id uuid.UUID, name string) error
	Delete(id uuid.UUID) error
}

type User struct {
	ID             uuid.UUID
	Name           string
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (user *User) Validate() error {
	// TODO
	return nil
}

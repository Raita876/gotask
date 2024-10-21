package domain

import "github.com/google/uuid"

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByID(id uuid.UUID) (*User, error)
	Update(id uuid.UUID, password string) (*User, error)
	Delete(id uuid.UUID) error
}

type User struct {
	ID       uuid.UUID
	Email    string
	Password string
}

func (user *User) Validate() error {
	// TODO
	return nil
}

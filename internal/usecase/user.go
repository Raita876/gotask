package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/raita876/gotask/internal/domain"
)

type UserUseCase interface {
	CreateUser(*CreateUserInput) (*CreateUserOutput, error)
	FindUserByID(*FindUserByIDInput) (*FindUserByIDOutput, error)
	UpdateUser(*UpdatePasswordInput) (*UpdatePasswordOutput, error)
	DeleteUser(*DeleteUserInput) (*DeleteUserOutput, error)
}

type UserOutput struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}
type CreateUserOutput struct {
	UserOutput *UserOutput
}

type FindUserByIDInput struct {
	ID uuid.UUID
}
type FindUserByIDOutput struct {
	UserOutput *UserOutput
}

type UpdatePasswordInput struct {
	ID       uuid.UUID
	Name     string
	Password string
}
type UpdatePasswordOutput struct {
	UserOutput *UserOutput
}

type DeleteUserInput struct {
	ID uuid.UUID
}
type DeleteUserOutput struct {
	// no param
}

type userInteractor struct {
	repo domain.UserRepository
}

func NewUserInteractor(repo domain.UserRepository) UserUseCase {
	return &userInteractor{
		repo: repo,
	}
}

func (i *userInteractor) CreateUser(input *CreateUserInput) (*CreateUserOutput, error) {
	user := &domain.User{
		ID:        uuid.New(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if _, err := i.repo.Create(user); err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		UserOutput: &UserOutput{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (i *userInteractor) FindUserByID(input *FindUserByIDInput) (*FindUserByIDOutput, error) {
	user, err := i.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindUserByIDOutput{
		UserOutput: &UserOutput{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (i *userInteractor) UpdateUser(input *UpdatePasswordInput) (*UpdatePasswordOutput, error) {
	user, err := i.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Password = input.Password

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := i.repo.Update(input.ID, input.Name, input.Password); err != nil {
		return nil, err
	}

	return &UpdatePasswordOutput{
		UserOutput: &UserOutput{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (i *userInteractor) DeleteUser(input *DeleteUserInput) (*DeleteUserOutput, error) {
	_, err := i.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if err := i.repo.Delete(input.ID); err != nil {
		return nil, err
	}

	return &DeleteUserOutput{}, nil
}

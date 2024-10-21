package usecase

import (
	"github.com/raita876/gotask/internal/domain"
)

type UserUseCase interface {
	CreateUser(*CreateUserInput) (*CreateUserOutput, error)
	FindUserByID(*FindUserByIDInput) (*FindUserByIDOutput, error)
	UpdateUser(*UpdatePasswordInput) (*UpdatePasswordOutput, error)
	DeleteUser(*DeleteUserInput) (*DeleteUserOutput, error)
}

type CreateUserInput struct {
	// TODO
}
type CreateUserOutput struct {
	// TODO
}

type FindUserByIDInput struct {
	// TODO
}
type FindUserByIDOutput struct {
	// TODO
}

type UpdatePasswordInput struct {
	// TODO
}
type UpdatePasswordOutput struct {
	// TODO
}

type DeleteUserInput struct {
	// TODO
}
type DeleteUserOutput struct {
	// TODO
}

type userInteractor struct {
	repo domain.UserRepository
}

func NewUserInteractor(repo domain.UserRepository) UserUseCase {
	return &userInteractor{
		repo: repo,
	}
}

func (uc *userInteractor) CreateUser(input *CreateUserInput) (*CreateUserOutput, error) {
	// TODO
	return nil, nil
}

func (uc *userInteractor) FindUserByID(input *FindUserByIDInput) (*FindUserByIDOutput, error) {
	// TODO
	return nil, nil
}

func (uc *userInteractor) UpdateUser(input *UpdatePasswordInput) (*UpdatePasswordOutput, error) {
	// TODO
	return nil, nil
}

func (uc *userInteractor) DeleteUser(input *DeleteUserInput) (*DeleteUserOutput, error) {
	// TODO
	return nil, nil
}

// type CreateUserUseCase interface {
// 	Execute(*CreateUserInput) (*CreateUserOutput, error)
// }

// type CreateUserInput struct{}

// type CreateUserOutput struct{}

// type createUserInteractor struct {
// 	repo domain.UserRepository
// }

// func NewCreateUserInteractor(repo domain.UserRepository) CreateUserUseCase {
// 	return &createUserInteractor{
// 		repo: repo,
// 	}
// }

// func (i *createUserInteractor) Execute(input *CreateUserInput) (*CreateUserOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type FindUserByIDUseCase interface {
// 	Execute(*FindUserByIDInput) (*FindUserByIDOutput, error)
// }

// type FindUserByIDInput struct{}

// type FindUserByIDOutput struct{}

// type findUserByIDInteractor struct {
// 	repo domain.UserRepository
// }

// func NewFindUserByIDInteractor(repo domain.UserRepository) FindUserByIDUseCase {
// 	return &findUserByIDInteractor{
// 		repo: repo,
// 	}
// }

// func (i *findUserByIDInteractor) Execute(input *FindUserByIDInput) (*FindUserByIDOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type UpdatePasswordUsecase interface {
// 	Execute(*UpdatePasswordInput) (*UpdatePasswordOutput, error)
// }

// type UpdatePasswordInput struct{}

// type UpdatePasswordOutput struct{}

// type updatePasswordInteractor struct {
// 	repo domain.UserRepository
// }

// func NewUpdatePasswordInteractor(repo domain.UserRepository) UpdatePasswordUsecase {
// 	return &updatePasswordInteractor{
// 		repo: repo,
// 	}
// }

// func (i *updatePasswordInteractor) Execute(input *UpdatePasswordInput) (*UpdatePasswordOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type DeleteUserUseCase interface {
// 	Execute(*DeleteUserInput) (*DeleteUserOutput, error)
// }

// type DeleteUserInput struct{}

// type DeleteUserOutput struct{}

// type deleteUserInteractor struct {
// 	repo domain.UserRepository
// }

// func NewDeleteUserInteractor(repo domain.UserRepository) DeleteUserUseCase {
// 	return &deleteUserInteractor{
// 		repo: repo,
// 	}
// }

// func (i *deleteUserInteractor) Execute(input *DeleteUserInput) (*DeleteUserOutput, error) {
// 	// TODO
// 	return nil, nil
// }

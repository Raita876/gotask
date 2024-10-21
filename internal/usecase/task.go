package usecase

import "github.com/raita876/gotask/internal/domain"

type TaskUseCase interface {
	CreateTask(*CreateTaskInput) (*CreateTaskOutput, error)
	FindTaskByID(*FindTaskByIDInput) (*FindTaskByIDOutput, error)
	FindTasksByUserID(*FindTasksByUserIDInput) (*FindTasksByUserIDOutput, error)
	UpdateTask(*UpdateTaskInput) (*UpdateTaskOutput, error)
	DeleteTask(*DeleteTaskInput) (*DeleteTaskOutput, error)
}

type CreateTaskInput struct {
	// TODO
}
type CreateTaskOutput struct {
	// TODO
}

type FindTaskByIDInput struct {
	// TODO
}
type FindTaskByIDOutput struct {
	// TODO
}

type FindTasksByUserIDInput struct {
	// TODO
}
type FindTasksByUserIDOutput struct {
	// TODO
}

type UpdateTaskInput struct {
	// TODO
}
type UpdateTaskOutput struct {
	// TODO
}

type DeleteTaskInput struct {
	// TODO
}
type DeleteTaskOutput struct {
	// TODO
}

type taskInteractor struct {
	repo domain.TaskRepository
}

func NewTaskInteractor(repo domain.TaskRepository) TaskUseCase {
	return &taskInteractor{
		repo: repo,
	}
}

func (i *taskInteractor) CreateTask(input *CreateTaskInput) (*CreateTaskOutput, error) {
	// TODO
	return nil, nil
}

func (i *taskInteractor) FindTaskByID(input *FindTaskByIDInput) (*FindTaskByIDOutput, error) {
	// TODO
	return nil, nil
}

func (i *taskInteractor) FindTasksByUserID(input *FindTasksByUserIDInput) (*FindTasksByUserIDOutput, error) {
	// TODO
	return nil, nil
}

func (i *taskInteractor) UpdateTask(input *UpdateTaskInput) (*UpdateTaskOutput, error) {
	// TODO
	return nil, nil
}

func (i *taskInteractor) DeleteTask(input *DeleteTaskInput) (*DeleteTaskOutput, error) {
	// TODO
	return nil, nil
}

// type CreateTaskUseCase interface {
// 	Execute(*CreateTaskInput) (*CreateTaskOutput, error)
// }

// type CreateTaskInput struct{}

// type CreateTaskOutput struct{}

// type createTaskInteractor struct {
// 	repo domain.TaskRepository
// }

// func NewCreateTaskInteractor(repo domain.TaskRepository) CreateTaskUseCase {
// 	return &createTaskInteractor{
// 		repo: repo,
// 	}
// }

// func (i createTaskInteractor) Execute(input *CreateTaskInput) (*CreateTaskOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type FindTaskByIDUseCase interface {
// 	Execute(*FindTaskByIDInput) (*FindTaskByIDOutput, error)
// }

// type FindTaskByIDInput struct{}

// type FindTaskByIDOutput struct{}

// type findTaskByIDInteractor struct {
// 	repo domain.TaskRepository
// }

// func NewFindTaskByIDInteractor(repo domain.TaskRepository) FindTaskByIDUseCase {
// 	return &findTaskByIDInteractor{
// 		repo: repo,
// 	}
// }

// func (i *findTaskByIDInteractor) Execute(input *FindTaskByIDInput) (*FindTaskByIDOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type FindTasksByUserIDUseCase interface {
// 	Execute(*FindTasksByUserIDInput) (*FindTasksByUserIDOutput, error)
// }

// type FindTasksByUserIDInput struct{}

// type FindTasksByUserIDOutput struct{}

// type FindTasksByUserIDInteractor struct {
// 	repo domain.TaskRepository
// }

// func NewFindTasksByUserIDInteractor(repo domain.TaskRepository) FindTasksByUserIDUseCase {
// 	return &FindTasksByUserIDInteractor{
// 		repo: repo,
// 	}
// }

// func (i *FindTasksByUserIDInteractor) Execute(input *FindTasksByUserIDInput) (*FindTasksByUserIDOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type UpdateTaskNameUseCase interface {
// 	Execute(*UpdateTaskNameInput) (*UpdateTaskNameOutput, error)
// }

// type UpdateTaskNameInput struct{}

// type UpdateTaskNameOutput struct{}

// type updateTaskNameInteractor struct {
// 	repo domain.TaskRepository
// }

// func NewUpdateTaskNameInteractor(repo domain.TaskRepository) UpdateTaskNameUseCase {
// 	return &updateTaskNameInteractor{
// 		repo: repo,
// 	}
// }

// func (i *updateTaskNameInteractor) Execute(input *UpdateTaskNameInput) (*UpdateTaskNameOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type UpdateTaskStatusUseCase interface {
// 	Execute(*UpdateTaskStatusInput) (*UpdateTaskStatusOutput, error)
// }

// type UpdateTaskStatusInput struct{}

// type UpdateTaskStatusOutput struct{}

// type updateTaskStatusInteractor struct {
// 	repo domain.TaskRepository
// }

// func NewUpdateTaskStatusInteractor(repo domain.TaskRepository) UpdateTaskStatusUseCase {
// 	return &updateTaskStatusInteractor{
// 		repo: repo,
// 	}
// }

// func (i *updateTaskStatusInteractor) Execute(input *UpdateTaskStatusInput) (*UpdateTaskStatusOutput, error) {
// 	// TODO
// 	return nil, nil
// }

// type DeleteTaskUseCase interface {
// 	Execute(*DeleteTaskInput) (*DeleteTaskOutput, error)
// }

// type DeleteTaskInput struct{}

// type DeleteTaskOutput struct{}

// type deleteTaskInteractor struct {
// 	repo domain.TaskRepository
// }

// func NewDeleteTaskInteractor(repo domain.TaskRepository) DeleteTaskUseCase {
// 	return &deleteTaskInteractor{
// 		repo: repo,
// 	}
// }

// func (i *deleteTaskInteractor) Execute(input *DeleteTaskInput) (*DeleteTaskOutput, error) {
// 	// TODO
// 	return nil, nil
// }

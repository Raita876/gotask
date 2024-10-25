package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/raita876/gotask/internal/domain"
)

type TaskUseCase interface {
	CreateTask(*CreateTaskInput) (*CreateTaskOutput, error)
	FindTaskByID(*FindTaskByIDInput) (*FindTaskByIDOutput, error)
	FindTasksByUserID(*FindTasksByUserIDInput) (*FindTasksByUserIDOutput, error)
	UpdateTask(*UpdateTaskInput) (*UpdateTaskOutput, error)
	DeleteTask(*DeleteTaskInput) (*DeleteTaskOutput, error)
}

type TaskOutput struct {
	ID        uuid.UUID
	Name      string
	Status    uint16
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateTaskInput struct {
	Name   string
	Status uint16
	UserID uuid.UUID
}
type CreateTaskOutput struct {
	TaskOutput *TaskOutput
}

type FindTaskByIDInput struct {
	ID uuid.UUID
}
type FindTaskByIDOutput struct {
	TaskOutput *TaskOutput
}

type FindTasksByUserIDInput struct {
	UserID uuid.UUID
}
type FindTasksByUserIDOutput struct {
	TasksOutput []*TaskOutput
}

type UpdateTaskInput struct {
	ID     uuid.UUID
	Name   string
	Status uint16
}
type UpdateTaskOutput struct {
	TaskOutput *TaskOutput
}

type DeleteTaskInput struct {
	ID uuid.UUID
}
type DeleteTaskOutput struct {
	// no param
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
	task := &domain.Task{
		ID:        uuid.New(),
		Name:      input.Name,
		Status:    input.Status,
		UserID:    input.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := task.Validate(); err != nil {
		return nil, err
	}

	if _, err := i.repo.Create(task); err != nil {
		return nil, err
	}

	return &CreateTaskOutput{
		TaskOutput: &TaskOutput{
			ID:        task.ID,
			Name:      task.Name,
			Status:    task.Status,
			UserID:    task.UserID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		},
	}, nil
}

func (i *taskInteractor) FindTaskByID(input *FindTaskByIDInput) (*FindTaskByIDOutput, error) {
	task, err := i.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindTaskByIDOutput{
		TaskOutput: &TaskOutput{
			ID:        task.ID,
			Name:      task.Name,
			Status:    task.Status,
			UserID:    task.UserID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		},
	}, nil
}

func toTasksOutput(tasks []*domain.Task) []*TaskOutput {
	tasksOutput := make([]*TaskOutput, len(tasks))

	for i, task := range tasks {
		tasksOutput[i] = &TaskOutput{
			ID:        task.ID,
			Name:      task.Name,
			Status:    task.Status,
			UserID:    task.UserID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
	}

	return tasksOutput
}

func (i *taskInteractor) FindTasksByUserID(input *FindTasksByUserIDInput) (*FindTasksByUserIDOutput, error) {
	tasks, err := i.repo.FindByUserID(input.UserID)
	if err != nil {
		return nil, err
	}

	return &FindTasksByUserIDOutput{
		TasksOutput: toTasksOutput(tasks),
	}, nil
}

func (i *taskInteractor) UpdateTask(input *UpdateTaskInput) (*UpdateTaskOutput, error) {
	task, err := i.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	task.Name = input.Name
	task.Status = input.Status

	if err := task.Validate(); err != nil {
		return nil, err
	}

	if err := i.repo.Update(input.ID, input.Name, input.Status); err != nil {
		return nil, err
	}

	return &UpdateTaskOutput{
		TaskOutput: &TaskOutput{
			ID:        task.ID,
			Name:      task.Name,
			Status:    task.Status,
			UserID:    task.UserID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		},
	}, nil
}

func (i *taskInteractor) DeleteTask(input *DeleteTaskInput) (*DeleteTaskOutput, error) {
	_, err := i.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if err := i.repo.Delete(input.ID); err != nil {
		return nil, err
	}

	return &DeleteTaskOutput{}, nil
}

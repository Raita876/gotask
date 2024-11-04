package rest

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/raita876/gotask/internal/usecase"
)

type TaskController struct {
	uc usecase.TaskUseCase
}

func NewTaskController(e *echo.Echo, uc usecase.TaskUseCase) *TaskController {
	ctr := &TaskController{
		uc: uc,
	}

	e.POST("/api/v1/tasks", ctr.CreateTask)
	e.GET("/api/v1/tasks/:id", ctr.FindTaskByID)
	e.GET("/api/v1/tasks/:user_id", ctr.FindTasksByUserID)
	e.PUT("/api/v1/tasks", ctr.UpdateTask)
	e.DELETE("/api/v1/tasks", ctr.DeleteTask)

	return ctr
}

// @Summary Create task
// @Schemes http
// @Description Create task
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body CreateTaskRequest true "request body"
// @Success 200 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func (ctr *TaskController) CreateTask(c echo.Context) error {
	var createTaskRequest CreateTaskRequest

	if err := c.Bind(&createTaskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	userID, err := uuid.Parse(createTaskRequest.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.CreateTaskInput{
		Name:   createTaskRequest.Name,
		Status: createTaskRequest.Status,
		UserID: userID,
	}
	output, err := ctr.uc.CreateTask(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create task",
		})
	}

	response := &TaskResponse{
		ID:        output.TaskOutput.ID.String(),
		Name:      output.TaskOutput.Name,
		Status:    output.TaskOutput.Status,
		UserID:    output.TaskOutput.UserID.String(),
		CreatedAt: output.TaskOutput.CreatedAt,
		UpdatedAt: output.TaskOutput.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary Get task by id
// @Schemes http
// @Description Get task by id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "task id"
// @Success 200 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [get]
func (ctr *TaskController) FindTaskByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.FindTaskByIDInput{
		ID: id,
	}
	output, err := ctr.uc.FindTaskByID(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to find task",
		})
	}

	response := &TaskResponse{
		ID:        output.TaskOutput.ID.String(),
		Name:      output.TaskOutput.Name,
		Status:    output.TaskOutput.Status,
		UserID:    output.TaskOutput.UserID.String(),
		CreatedAt: output.TaskOutput.CreatedAt,
		UpdatedAt: output.TaskOutput.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response)
}

func toTasks(tasksOutput []*usecase.TaskOutput) []*TaskResponse {
	tasks := make([]*TaskResponse, len(tasksOutput))

	for i, t := range tasksOutput {
		tasks[i] = &TaskResponse{
			ID:        t.ID.String(),
			Name:      t.Name,
			Status:    t.Status,
			UserID:    t.UserID.String(),
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}
	}

	return tasks
}

// @Summary Get tasks by user id
// @Schemes http
// @Description get tasks by user id
// @Tags tasks
// @Accept json
// @Produce json
// @Param user_id path string true "user id"
// @Success 200 {object} TasksResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{user_id} [get]
func (ctr *TaskController) FindTasksByUserID(c echo.Context) error {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.FindTasksByUserIDInput{
		UserID: userID,
	}
	output, err := ctr.uc.FindTasksByUserID(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete task",
		})
	}

	response := &TasksResponse{
		Tasks: toTasks(output.TasksOutput),
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary Update task
// @Schemes http
// @Description Update task
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body UpdateTaskRequest true "request body"
// @Success 200 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [put]
func (ctr *TaskController) UpdateTask(c echo.Context) error {
	var updateTaskRequest UpdateTaskRequest

	if err := c.Bind(&updateTaskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	id, err := uuid.Parse(updateTaskRequest.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.UpdateTaskInput{
		ID:     id,
		Name:   updateTaskRequest.Name,
		Status: updateTaskRequest.Status,
	}
	output, err := ctr.uc.UpdateTask(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update task",
		})
	}

	response := &TaskResponse{
		ID:        output.TaskOutput.ID.String(),
		Name:      output.TaskOutput.Name,
		Status:    output.TaskOutput.Status,
		UserID:    output.TaskOutput.UserID.String(),
		CreatedAt: output.TaskOutput.CreatedAt,
		UpdatedAt: output.TaskOutput.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary Delete task
// @Schemes http
// @Description Delete task
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body DeleteTaskRequest true "request body"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [delete]
func (ctr *TaskController) DeleteTask(c echo.Context) error {
	var deleteTaskRequest DeleteTaskRequest

	if err := c.Bind(&deleteTaskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	id, err := uuid.Parse(deleteTaskRequest.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.DeleteTaskInput{
		ID: id,
	}
	_, err = ctr.uc.DeleteTask(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete task",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted user",
	})
}

type CreateTaskRequest struct {
	Name   string `json:"name"`
	Status uint16 `json:"status"`
	UserID string `json:"user_id"`
}

type UpdateTaskRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status uint16 `json:"status"`
}

type DeleteTaskRequest struct {
	ID string `json:"id"`
}

type TaskResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    uint16    `json:"status"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TasksResponse struct {
	Tasks []*TaskResponse
}

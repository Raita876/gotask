package rest

import (
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
	e.GET("/api/v1/tasks", ctr.FindTasksByUserID)
	e.PUT("/api/v1/tasks/:id", ctr.UpdateTask)
	e.DELETE("/api/v1/tasks/:id", ctr.DeleteTask)

	return ctr
}

func (ctr *TaskController) CreateTask(c echo.Context) error {
	// TODO
	return nil
}

func (ctr *TaskController) FindTaskByID(c echo.Context) error {
	// TODO
	return nil
}

func (ctr *TaskController) FindTasksByUserID(c echo.Context) error {
	// TODO
	return nil
}

func (ctr *TaskController) UpdateTask(c echo.Context) error {
	// TODO
	return nil
}

func (ctr *TaskController) DeleteTask(c echo.Context) error {
	// TODO
	return nil
}

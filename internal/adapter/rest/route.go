package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/raita876/gotask/internal/usecase"
)

func Setup(
	userUseCase usecase.UserUseCase,
	taskUseCase usecase.TaskUseCase,
	name, version, usage string,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	apiGroup := e.Group("/api/v1")

	NewUserController(apiGroup, userUseCase)
	NewTaskController(apiGroup, taskUseCase)
	NewSwagController(apiGroup, name, version, usage)

	return e
}

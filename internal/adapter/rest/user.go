package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/raita876/gotask/internal/usecase"
)

type UserController struct {
	uc usecase.UserUseCase
}

func NewUserController(e *echo.Echo, uc usecase.UserUseCase) *UserController {
	ctr := &UserController{
		uc: uc,
	}

	e.POST("/api/v1/users", ctr.CreateUser)
	e.GET("/api/v1/users/:id", ctr.FindUserByID)
	e.PUT("/api/v1/users/:id", ctr.UpdateUser)
	e.DELETE("/api/v1/users/:id", ctr.DeleteUser)

	return ctr
}

func (ctr *UserController) CreateUser(c echo.Context) error {
	// TODO
	return nil
}

func (ctr *UserController) FindUserByID(c echo.Context) error {
	// TODO
	return nil
}

func (ctr *UserController) UpdateUser(c echo.Context) error {
	// TODO
	return nil
}

func (ctr *UserController) DeleteUser(c echo.Context) error {
	// TODO
	return nil
}

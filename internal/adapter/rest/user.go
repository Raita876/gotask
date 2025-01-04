package rest

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/raita876/gotask/internal/usecase"
)

type UserController struct {
	uc usecase.UserUseCase
}

func NewUserController(group *echo.Group, uc usecase.UserUseCase) *UserController {
	ctr := &UserController{
		uc: uc,
	}

	group.POST("/users", ctr.CreateUser)
	group.GET("/users/:id", ctr.FindUserByID)
	group.PUT("/sers/:id", ctr.UpdateUser)
	group.DELETE("/users/:id", ctr.DeleteUser)
	group.POST("/users/login", ctr.Login)

	return ctr
}

// @Summary Create user
// @Schemes http
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "request body"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (ctr *UserController) CreateUser(c echo.Context) error {
	var createUserRequest CreateUserRequest

	if err := c.Bind(&createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	input := &usecase.CreateUserInput{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	output, err := ctr.uc.CreateUser(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user",
		})
	}

	response := &UserResponse{
		ID:        output.UserOutput.ID.String(),
		Name:      output.UserOutput.Name,
		Email:     output.UserOutput.Email,
		CreatedAt: output.UserOutput.CreatedAt,
		UpdatedAt: output.UserOutput.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, response)
}

// @Summary Get user by id
// @Schemes http
// @Description Get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (ctr *UserController) FindUserByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.FindUserByIDInput{
		ID: id,
	}
	output, err := ctr.uc.FindUserByID(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to find user",
		})
	}

	response := &UserResponse{
		ID:        output.UserOutput.ID.String(),
		Name:      output.UserOutput.Name,
		Email:     output.UserOutput.Email,
		CreatedAt: output.UserOutput.CreatedAt,
		UpdatedAt: output.UserOutput.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary Update user
// @Schemes http
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param request body UpdateUserRequest true "request body"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [put]
func (ctr *UserController) UpdateUser(c echo.Context) error {
	var updateUserRequest UpdateUserRequest

	if err := c.Bind(&updateUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	id, err := uuid.Parse(updateUserRequest.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.UpdatePasswordInput{
		ID:   id,
		Name: updateUserRequest.Name,
	}
	output, err := ctr.uc.UpdateUser(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update user",
		})
	}

	response := &UserResponse{
		ID:        output.UserOutput.ID.String(),
		Name:      output.UserOutput.Name,
		Email:     output.UserOutput.Email,
		CreatedAt: output.UserOutput.CreatedAt,
		UpdatedAt: output.UserOutput.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary Delete user
// @Schemes http
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param request body DeleteUserRequest true "request body"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [delete]
func (ctr *UserController) DeleteUser(c echo.Context) error {
	var deleteUserRequest DeleteUserRequest

	if err := c.Bind(&deleteUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	id, err := uuid.Parse(deleteUserRequest.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	input := &usecase.DeleteUserInput{
		ID: id,
	}
	_, err = ctr.uc.DeleteUser(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted user",
	})
}

// @Summary Login user
// @Schemes http
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "request body"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/login [post]
func (ctr *UserController) Login(c echo.Context) error {
	var loginRequest LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	input := &usecase.LoginUserInput{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	output, err := ctr.uc.LoginUser(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to login user",
		})
	}

	if !output.IsSuccessful {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Not allowed",
		})
	}

	signedToken, err := CreateJwt(output.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create jwt",
		})
	}

	response := &LoginResponse{
		Token: signedToken,
	}

	return c.JSON(http.StatusOK, response)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

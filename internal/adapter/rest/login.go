package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raita876/gotask/internal/usecase"
)

type LoginController struct {
	uc usecase.UserUseCase
}

func NewLoginController(group *echo.Group, uc usecase.UserUseCase) *LoginController {
	ctr := &LoginController{
		uc: uc,
	}

	loginApiGroup := group.Group("/login")

	loginApiGroup.POST("", ctr.Login)

	return ctr
}

// @Summary Login
// @Schemes http
// @Description Login
// @Tags login
// @Accept json
// @Produce json
// @Param request body LoginRequest true "request body"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (ctr *LoginController) Login(c echo.Context) error {
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

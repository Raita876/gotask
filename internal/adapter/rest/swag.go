package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/raita876/gotask/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type SwagController struct{}

func NewSwagController(e *echo.Echo, name, version, usage string) *SwagController {
	docs.SwaggerInfo.Title = name
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Description = usage
	docs.SwaggerInfo.BasePath = "/api/v1"

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &SwagController{}
}

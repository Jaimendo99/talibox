package AuthRoutes

import (
	"talibox/controllers/AuthController"

	"github.com/labstack/echo/v4"
)

func DefineAuthRoutes(e *echo.Echo) {
	e.POST("/users", AuthController.SignUp())
	e.POST("/login", AuthController.Login())
	e.GET("/login", AuthController.LoginPage())
}

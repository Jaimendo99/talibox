package AuthController

import (
	loginpage "talibox/views/LoginPage"
	uicomponents "talibox/views/UiComponents"

	"github.com/labstack/echo/v4"
)

func LoginAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request()

		return c.JSON(200, "Login Action")
	}
}

func LoginPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return uicomponents.MainLayout("Login Page", loginpage.LoginForm()).Render(c.Request().Context(), c.Response().Writer)
	}
}

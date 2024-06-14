package AuthController

import (
	loginpage "talibox/views/LoginPage"
	uicomponents "talibox/views/UiComponents"
	"talibox/views/uiModels"

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
		return uicomponents.MainLayout("Login Page",
			loginpage.LoginForm(
				uiModels.InputModel{
					Name:        "username",
					Label:       "Username",
					Id:          "username",
					Placeholder: "Enter your username",
					Type:        "text",
				},
				uiModels.InputModel{
					Name:        "password",
					Label:       "Password",
					Id:          "password",
					Placeholder: "Enter your password",
					Type:        "password",
				},
				[]uiModels.ErrorModel{},
			),
		).Render(c.Request().Context(), c.Response().Writer)
	}
}

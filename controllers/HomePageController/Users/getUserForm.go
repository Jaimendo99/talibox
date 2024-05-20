package users

import (
	"talibox/inits"
	"talibox/models"
	"talibox/views/Forms"

	"github.com/labstack/echo/v4"
)

func GetUserForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		return Forms.RegisterForm().Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetUpadteForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// Get the user from the database
		user := models.User{}

		// Find the user with the given id
		inits.DB.First(&user, id)

		return Forms.EditForm(user).Render(c.Request().Context(), c.Response().Writer)
	}
}

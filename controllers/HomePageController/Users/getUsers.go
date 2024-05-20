package users

import (
	"talibox/inits"
	"talibox/models"
	homepage "talibox/views/HomePage"

	"github.com/labstack/echo/v4"
)

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}
		id := c.Param("id")
		inits.DB.First(&user, id)

		return homepage.GenTableRow(user.GetFieldsValues(), user.GetInstanceName()).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := []models.User{}

		inits.DB.Find(&users)

		return homepage.GenericTable(users).Render(c.Request().Context(), c.Response().Writer)
	}
}

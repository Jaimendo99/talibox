package users

import (
	"fmt"
	"talibox/inits"
	"talibox/models"
	homepage "talibox/views/HomePage"

	"github.com/labstack/echo/v4"
)

func UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}

		inits.DB.First(&user, c.Param("id"))

		supUser := models.UserSignUp{}

		if err := c.Bind(&supUser); err != nil {
			fmt.Println("ERROR: " + err.Error())
		}

		user.UserBasic.FullName = supUser.FullName
		user.UserBasic.UserLogin.Username = supUser.UserLogin.Username
		if supUser.UserLogin.Password != "" {
			user.UserBasic.UserLogin.Password = supUser.UserLogin.Password
		}

		inits.DB.Save(&user)

		return homepage.GenTableRow(user.GetFieldsValues(), user.GetInstanceName()).Render(c.Request().Context(), c.Response().Writer)
	}
}

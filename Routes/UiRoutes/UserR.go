package uiroutes

import (
	users "talibox/controllers/HomePageController/Users"

	"github.com/labstack/echo/v4"
)

func DefineUserRoutes(e *echo.Echo, middleware ...echo.MiddlewareFunc) {
	e.GET("/users/:id", users.GetUser(), middleware...)
	e.GET("/users", users.GetUsers(), middleware...)
	e.GET("/user/form", users.GetUserForm(), middleware...)
	e.GET("/user/form/:id", users.GetUpadteForm(), middleware...)
}

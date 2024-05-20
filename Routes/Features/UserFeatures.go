package features

import (
	users "talibox/controllers/HomePageController/Users"

	"github.com/labstack/echo/v4"
)

func DefineUserFeaturesRoutes(e *echo.Echo, middleware ...echo.MiddlewareFunc) {
	e.PATCH("/user/:id", users.UpdateUser(), middleware...)
	e.DELETE("/user/:id", users.DeleteUser(), middleware...)
}

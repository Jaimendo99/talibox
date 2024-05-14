package uiroutes

import (
	homepagecontroller "talibox/controllers/HomePageController"

	"github.com/labstack/echo/v4"
)

func DefineHomeRoutes(e *echo.Echo, middleware ...echo.MiddlewareFunc) {

	if len(middleware) > 0 {
		e.GET("/", homepagecontroller.GetHomePage(), middleware...)
	} else {
		e.GET("/", homepagecontroller.GetHomePage())
	}
}

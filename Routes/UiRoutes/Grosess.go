package uiroutes

import (
	grossess "talibox/controllers/Grossess"

	"github.com/labstack/echo/v4"
)

func DefineGrossessRoutes(e *echo.Echo) {
	e.GET("/grossess", grossess.GetWeeklyTop())
}

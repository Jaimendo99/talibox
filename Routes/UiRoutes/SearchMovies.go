package uiroutes

import (
	grossess "talibox/controllers/Grossess"

	"github.com/labstack/echo/v4"
)

func DefineSearchRoutes(e *echo.Echo) {
	e.GET("/search", grossess.GetSearchPage())
	e.POST("/search", grossess.SearchMoviesByName())
}

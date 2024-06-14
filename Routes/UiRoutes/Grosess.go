package uiroutes

import (
	core "talibox/controllers/CORE"
	grossess "talibox/controllers/Grossess"

	"github.com/labstack/echo/v4"
)

func DefineGrossessRoutes(e *echo.Echo) {
	e.GET("/grossess", grossess.GetWeeklyTop())
	e.GET("/grossess/chart", grossess.GetWeeklyTopChart())
	e.GET("/grossess/historic", grossess.GetHistoricYearWeekly())
	e.GET("/movie", grossess.GetWeeklyByMovieTitle())
	e.GET("/add", grossess.SaveMovieData())
	e.GET("/stats", grossess.GetYearToYearComparason())
	e.POST("/core", core.GetActorsAndGrossByAWeekNet())
	e.GET("/actor", core.GetActorForm())
}

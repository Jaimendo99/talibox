package grossess

import (
	external "talibox/controllers/External"
	"talibox/views/SearchMovies"

	"github.com/labstack/echo/v4"
)

func GetWeeklyByMovieTitle() echo.HandlerFunc {
	return func(c echo.Context) error {

		movieTitle := c.QueryParam("id")

		weeklyGross, err := external.GetWeeklyByMovie(movieTitle)
		if err != nil {
			// return err
			c.Logger().Error(err)
			c.Echo().Logger.Error(err)

		}
		return SearchMovies.WeeklyGrossView(weeklyGross).Render(c.Request().Context(), c.Response().Writer)
	}
}

package grossess

import (
	external "talibox/controllers/External"
	groses "talibox/views/Grosess"
	"talibox/views/uiModels"

	"github.com/labstack/echo/v4"
)

func GetWeeklyTop() echo.HandlerFunc {
	return func(c echo.Context) error {
		topGrossess := uiModels.TopGrosses{}

		topGrossess, err := external.GetWeeklyTopGross()

		if err != nil {
			return err
		}

		moviesTop := topGrossess.Movies

		topGrossessParsed := []uiModels.MovieTopParsed{}

		for _, movie := range moviesTop {
			topGrossessParsed = append(topGrossessParsed, movie.ParseGross())
		}

		return groses.TopGrossList(topGrossessParsed).Render(c.Request().Context(), c.Response().Writer)
	}
}

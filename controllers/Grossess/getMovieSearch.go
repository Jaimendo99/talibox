package grossess

import (
	components "talibox/controllers/Components"
	external "talibox/controllers/External"
	"talibox/views/SearchMovies"
	"talibox/views/uiModels"

	"github.com/labstack/echo/v4"
)

func GetSearchPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return SearchMovies.SearchView().Render(c.Request().Context(), c.Response().Writer)
	}
}

func SearchMoviesByName() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := components.ReadJwt(c)

		logged := false

		if user.ID != 0 {
			logged = true
		}

		movieSearch := []uiModels.MovieSearch{}
		err := error(nil)

		movieName := c.FormValue("q")

		if movieName != "" {
			movieSearch, err = external.SearchMovieByName(movieName)
		}

		if err != nil {
			return err
		}

		return SearchMovies.SearchResults(movieSearch, logged).Render(c.Request().Context(), c.Response().Writer)
	}
}

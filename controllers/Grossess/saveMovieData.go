package grossess

import (
	"strconv"
	"strings"
	external "talibox/controllers/External"
	"talibox/inits"
	"talibox/models"
	"talibox/views/SearchMovies"

	"github.com/labstack/echo/v4"
)

func SaveMovieData() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := inits.DB

		title := c.QueryParam("title")
		titleId := c.QueryParam("id")
		yearQP := c.QueryParam("year")
		actorsQP := c.QueryParam("actors")
		genresQP := c.QueryParam("genres")

		if title == "" ||
			titleId == "" ||
			yearQP == "" ||
			actorsQP == "" ||
			genresQP == "" {
			return SearchMovies.FailMessage("Missing data").Render(c.Request().Context(), c.Response().Writer)
		}

		weeklyGross, err := external.GetWeeklyByMovie(titleId)

		if err != nil {
			c.Logger().Error(err)
			c.Echo().Logger.Error(err)
			return SearchMovies.FailMessage(err.Error()).Render(c.Request().Context(), c.Response().Writer)
		}

		genresList := strings.Split(genresQP, ", ")

		genres := []models.GenreCatalog{}
		for _, genre := range genresList {
			genres = append(genres, models.GenreCatalog{Genre: genre})
		}

		actorsList := strings.Split(actorsQP, ", ")

		year, err := strconv.ParseInt(yearQP, 10, 64)

		if err != nil {
			c.Logger().Error(err)
			c.Echo().Logger.Error(err)
			return SearchMovies.FailMessage(err.Error()).Render(c.Request().Context(), c.Response().Writer)
		}

		releaseDate := weeklyGross.WeeklyGross[0].WeekStart

		movie := models.Movie{
			Title:       title,
			RealeseId:   weeklyGross.DomesticRealese,
			Year:        int(year),
			ReleaseDate: releaseDate,
			Genres:      genres,
		}

		cast := []models.Cast{}
		for _, actor := range actorsList {
			cast = append(cast, models.Cast{
				Movie: movie,
				Role:  "Actor",
				Professional: models.Professional{
					ArtisticName: actor,
				},
			})
		}

		result := db.Create(&movie)

		if result.Error != nil {
			c.Logger().Error(result.Error)
			c.Echo().Logger.Error(result.Error)
			return SearchMovies.FailMessage("error creating movie").Render(c.Request().Context(), c.Response().Writer)
		}

		for _, genre := range genres {
			db.Create(&genre)
		}
		for _, actor := range cast {
			db.Create(&actor)
		}

		return SearchMovies.SucessMessage("Movie data saved").Render(c.Request().Context(), c.Response().Writer)
	}
}

package users

import (
	"net/http"
	"os/user"
	"talibox/inits"

	"github.com/labstack/echo/v4"
)

type id struct {
	ID string `json:"id"`
}

func DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := user.User{}

		results := inits.DB.Delete(&user, id{ID: c.Param("id")})

		if results.Error != nil {
			return results.Error
		}
		return c.NoContent(http.StatusOK)

	}
}

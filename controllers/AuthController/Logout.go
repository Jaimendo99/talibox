package AuthController

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func LogOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetCookie(&http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: false,
		})
		return c.Redirect(http.StatusFound, "/login")
	}
}

package components

import (
	"os"
	"talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ReadJwt(c echo.Context) models.User {
	cookie, err := c.Cookie("jwt")

	if err == nil {
		claims := &Models.JwtCustomClaims{}
		_, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.Logger().Debug("ERROR: " + err.Error())
		}

		var user models.User
		result := inits.DB.First(&user, "id = ?", claims.Sub)

		if result.Error != nil {
			c.Logger().Debug("ERROR: " + result.Error.Error())
		}

		return user
	}
	return models.User{}
}

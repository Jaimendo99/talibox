package inits

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"talibox/controllers/Models"
)

var Config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return &Models.JwtCustomClaims{}
	},

	SigningKey: []byte(os.Getenv("JWT_SECRET")),

	TokenLookup: "cookie:jwt",

	ErrorHandler: func(c echo.Context, err error) error {
		fmt.Printf("JWT Error: %v\n", err) // Log the error for debugging

		return c.Redirect(http.StatusPermanentRedirect, "/login")
	},
}

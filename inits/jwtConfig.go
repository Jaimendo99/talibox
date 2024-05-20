package inits

import (
	"fmt"
	"net/http"
	"os"
	"talibox/controllers/Models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var Config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		fmt.Println(os.Getenv("JWT_SECRET"))
		return &Models.JwtCustomClaims{}
	},

	SigningKey: []byte("SECRET"),

	TokenLookup: "cookie:jwt",

	ErrorHandler: func(c echo.Context, err error) error {
		fmt.Printf("JWT Error: %v\n", err) // Log the error for debugging

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized",
		})
	},
}

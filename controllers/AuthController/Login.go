package AuthController

import (
	"net/http"
	"os"
	"talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.UserLogin
		var user models.User

		if c.Bind(&body) != nil {
			return c.JSON(http.StatusBadRequest, Models.FailResponse{Message: "Invalid request"})
		}

		result := inits.DB.Where("username = ?", body.Username).First(&user)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, Models.FailResponse{Message: "User not found"})
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.UserBasic.UserLogin.Password), []byte(body.Password))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, Models.FailResponse{Message: "Invalid password"})
		}

		claims := &Models.JwtCustomClaims{
			Sub:   user.ID,
			Admin: user.Admin,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, Models.FailResponse{Message: err.Error()})
		}

		c.SetCookie(&http.Cookie{
			Name:     "jwt",
			Value:    signedToken,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: false,
		})

		return c.Redirect(http.StatusFound, "/")
	}
}

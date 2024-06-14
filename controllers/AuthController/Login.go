package AuthController

import (
	"net/http"
	"os"
	"talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"
	loginpage "talibox/views/LoginPage"
	"talibox/views/uiModels"
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
			return renderLogin(c, "", "", []uiModels.ErrorModel{
				{Message: "Invalid form data"}})
		}

		result := inits.DB.Where("username = ?", body.Username).First(&user)
		if result.Error != nil {
			return renderLogin(c, body.Username, body.Password, []uiModels.ErrorModel{
				{Message: "User not found"}})
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.UserBasic.UserLogin.Password), []byte(body.Password))
		if err != nil {
			return renderLogin(c, body.Username, body.Password, []uiModels.ErrorModel{
				{Message: "Invalid password"}})
		}

		signedToken, err := genToken(&user)

		if err != nil {
			return renderLogin(c, body.Username, body.Password, []uiModels.ErrorModel{
				{Message: "Error signing token"}})
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

func renderLogin(c echo.Context, userValue, passValue string, errors []uiModels.ErrorModel) error {

	usernameForm, passwordForm, _ := getDefultLoginForm(userValue, passValue)

	return loginpage.LoginForm(
		usernameForm,
		passwordForm,
		errors,
	).Render(c.Request().Context(), c.Response().Writer)
}

func getDefultLoginForm(userValue, passValue string) (uiModels.InputModel, uiModels.InputModel, []uiModels.ErrorModel) {

	usernameForm := uiModels.InputModel{
		Name:        "username",
		Label:       "Username",
		Id:          "username",
		Placeholder: "Enter your username",
		Type:        "text",
		Value:       userValue,
	}

	passwordForm := uiModels.InputModel{
		Name:        "password",
		Label:       "Password",
		Id:          "password",
		Placeholder: "Enter your password",
		Type:        "password",
		Value:       passValue,
	}
	return usernameForm, passwordForm, []uiModels.ErrorModel{}
}

func genToken(user *models.User) (string, error) {
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
		return "", err
	}

	return signedToken, nil

}

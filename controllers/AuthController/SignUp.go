package AuthController

import (
	"errors"
	"net/http"
	"talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"
	homepage "talibox/views/HomePage"
	"unicode"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.UserSignUp

		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, Models.FailResponse{Message: err.Error()})
		}

		if err := validateBody(&body); err != nil {
			return c.JSON(http.StatusBadRequest, Models.FailResponse{Message: err.Error()})
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.UserLogin.Password), 10)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Models.FailResponse{Message: err.Error()})
		}

		user := models.User{
			UserBasic: models.UserSignUp{
				FullName: body.FullName,
				UserLogin: models.UserLogin{
					Username: body.UserLogin.Username,
					Password: string(hash),
				},
			},
			Admin:  false,
			Movies: nil,
		}

		dbResult := inits.DB.Create(&user)
		if dbResult.Error != nil {
			return c.JSON(http.StatusInternalServerError, Models.FailResponse{Message: dbResult.Error.Error()})
		}

		return homepage.GenTableRow(user.GetFieldsValues(), user.GetInstanceName()).Render(c.Request().Context(), c.Response().Writer)
	}
}

func validateBody(body *models.UserSignUp) error {

	if body.FullName == "" {
		return errors.New("fullname is required")
	}

	if body.UserLogin.Username == "" {
		return errors.New("username is required")
	}

	if body.UserLogin.Password == "" {
		return errors.New("password is required")
	}

	sevenOrMore, number, upper := validatePass(body.UserLogin.Password)
	if !sevenOrMore {
		return errors.New("password must be at least 7 characters long")
	}

	if !number {
		return errors.New("password must contain at least one number")
	}

	if !upper {
		return errors.New("password must contain at least one uppercase letter")
	}
	return nil
}

func validatePass(password string) (sevenOrMore, number, upper bool) {
	letters := 0
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}

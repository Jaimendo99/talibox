package homepagecontroller

import (
	"fmt"
	"os"
	"talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"
	homepage "talibox/views/HomePage"
	uicomponents "talibox/views/UiComponents"
	"talibox/views/uiModels"

	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetHomePage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var name string
		var content templ.Component
		cookie, err := c.Cookie("jwt")
		if err == nil {
			claims := &Models.JwtCustomClaims{}
			_, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
			}

			var user models.User
			result := inits.DB.First(&user, "id = ?", claims.Sub)

			if result.Error != nil {
				fmt.Println("ERROR: " + result.Error.Error())
			}

			users := []models.User{}
			inits.DB.Find(&users)
			content = homepage.GenericTable(users)

			name = user.UserBasic.FullName

		} else {
			name = "Guest"
			content = homepage.HomeContent()

		}

		return uicomponents.MainLayout("Home Page", homepage.HomeContainer(
			content,
			name,
			[]uiModels.BarItem{
				{
					Title: "Home",
					Ref:   "/",
					Icon:  "home",
				},
				{
					Title: "Weekly Grosses",
					Ref:   "/grossess",
					Icon:  "movie",
				},
				{
					Title: "Login",
					Ref:   "/login",
					Icon:  "person",
				},
			},
		)).Render(c.Request().Context(), c.Response().Writer)
	}
}

package homepagecontroller

import (
	"fmt"
	"os"
	"strconv"
	"talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"
	homepage "talibox/views/HomePage"
	uicomponents "talibox/views/UiComponents"
	"talibox/views/uiModels"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetHomePage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var sub int
		var name string
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

			name = user.UserBasic.FullName
			sub = int(claims.Sub)

		} else {
			sub = 0
			name = "Guest"
			// err := c.Redirect(302, "/login")
			// if err != nil {
			// fmt.Println(err)
			// }
		}

		content := homepage.HomeContent(
			strconv.Itoa(sub),
		)

		return uicomponents.MainLayout("Home Page", homepage.HomeContainer(
			content,
			name,
			[]uiModels.BarItem{
				{
					Title: "Home",
					Ref:   "/",
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

package homepagecontroller

import (
	"fmt"
	"os"
	"talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"
	"talibox/views/Grosess"
	homepage "talibox/views/HomePage"
	uicomponents "talibox/views/UiComponents"
	"talibox/views/uiModels"

	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetHomePage() echo.HandlerFunc {
	return func(c echo.Context) error {

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

			if user.Admin {
				return PageDefaultContainer(
					true,
					homepage.GenericTable(users),
					"Home Page",
					user.UserBasic.FullName,
				).Render(c.Request().Context(), c.Response().Writer)
			}

			return PageDefaultContainer(
				false,
				Grosess.GetChart(),
				"Home Page",
				user.UserBasic.FullName,
			).Render(c.Request().Context(), c.Response().Writer)
		}

		return PageDefaultContainer(
			false,
			homepage.HomeContent(),
			"Home Page",
			"Guest",
		).Render(c.Request().Context(), c.Response().Writer)
	}
}

func PageDefaultContainer(isAdmin bool, content templ.Component, title, name string) templ.Component {

	barItems := []uiModels.BarItem{

		{
			Title:   "Weekly Grosses",
			Ref:     "/grossess/chart",
			Icon:    "movie",
			Swap:    "innerHTML transition:true",
			Target:  "#homecontent",
			PushUrl: "true",
		},
		{
			Title:   "Historic Weekly Grosses",
			Ref:     "/grossess/historic",
			Icon:    "tv",
			Swap:    "innerHTML transition:true",
			Target:  "#homecontent",
			PushUrl: "true",
		}, {
			Title:   "Insights",
			Ref:     "/stats",
			Icon:    "query_stats",
			Swap:    "innerHTML transition:true",
			Target:  "#homecontent",
			PushUrl: "true",
		},
		{
			Title:   "core",
			Ref:     "/actor",
			Icon:    "person",
			Swap:    "innerHTML transition:true",
			Target:  "#homecontent",
			PushUrl: "true",
		},
		{
			Title:   "Search Movies",
			Ref:     "/search",
			Icon:    "search",
			Swap:    "innerHTML transition:true",
			Target:  "#homecontent",
			PushUrl: "true",
		},

		{
			Title:   "Login",
			Ref:     "/login",
			Icon:    "person",
			Swap:    "innerHTML transition:true",
			Target:  "body",
			PushUrl: "true",
		},
	}

	if isAdmin {
		adminItem := []uiModels.BarItem{
			{
				Title:   "Home",
				Ref:     "/",
				Icon:    "home",
				Swap:    "innerHTML transition:true",
				Target:  "body",
				PushUrl: "true",
			},
		}

		barItems = append(
			adminItem,
			barItems...,
		)
	}

	return uicomponents.MainLayout(title,
		homepage.HomeContainer(
			content,
			name,
			barItems,
		))
}

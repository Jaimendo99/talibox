package main

import (
	"fmt"
	"os"
	"talibox/Routes/AuthRoutes"
	uiroutes "talibox/Routes/UiRoutes"
	"talibox/inits"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	inits.LoadEnvVariables()
	inits.InitDB()

	err := inits.Migrate(inits.DB)
	if err != nil {
		fmt.Println(err)
		panic("failed to migrate")
	}

}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	AuthRoutes.DefineAuthRoutes(e)

	m := echojwt.WithConfig(inits.Config)

	fmt.Print(m)

	uiroutes.DefineHomeRoutes(e)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}

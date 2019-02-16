package bootstrap

import (
	"go-account-api/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Initialize app
func App() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Auth(e.Group("/auth"))

	e.Logger.Fatal(e.Start(":1323"))
}

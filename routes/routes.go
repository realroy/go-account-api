package routes

import (
	"go-account-api/controllers"

	"github.com/labstack/echo"
)

func Auth(g *echo.Group) {
	g.POST("/login", controllers.Login)
	g.POST("/register", controllers.Register)
	g.POST("/logout", controllers.Logout)
}

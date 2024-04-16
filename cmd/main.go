package main

import (
	"github.com/Alexis3386/goPortfolio/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handler.NewRepo().Home)
	e.GET("/projects", handler.NewRepo().Projects)

	// serve static files
	e.Static("/static", "assets")

	e.Start(":3000")
}

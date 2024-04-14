package main

import (
	"github.com/Alexis3386/goPortfolio/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handler.NewRepo().Home)

	// serve static files
	e.Static("/static", "assets")

	e.Start(":3000")
}

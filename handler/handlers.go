package handler

import (
	"github.com/Alexis3386/goPortfolio/templates"
	"github.com/labstack/echo/v4"
)

var Repo *Repository

type Repository struct{}

func NewRepo() *Repository {
	return &Repository{}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (r *Repository) Home(c echo.Context) error {
	content := templates.Home()
	return render(c, templates.Layout("My website", true, content))
}

package handler

import (
	"bytes"
	"net/http"

	"github.com/Alexis3386/goPortfolio/templates"
	"github.com/Alexis3386/goPortfolio/templates/project"
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

func (r *Repository) Projects(c echo.Context) error {
	projectNum := c.QueryParam("id")
	var tpl bytes.Buffer
	switch projectNum {
	case "1":
		if err := project.Project1().Render(c.Request().Context(), &tpl); err != nil {
			return err
		}
	case "2":
		if err := project.Project2().Render(c.Request().Context(), &tpl); err != nil {
			return err
		}
	}
	content := tpl.String()
	return c.String(http.StatusOK, content)
}

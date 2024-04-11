package handler

import (
	"github.com/Alexis3386/goPortfolio/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}

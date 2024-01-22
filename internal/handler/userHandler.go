package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/ssss-tantalum/echo-kick/internal/core"
	"github.com/ssss-tantalum/echo-kick/internal/model"
	"github.com/ssss-tantalum/echo-kick/pkg/templ"
	"github.com/ssss-tantalum/echo-kick/web/views/pages/user"
)

type UserHandler struct {
	app *core.App
}

func NewUserHandler(app *core.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) HandleUserShow(c echo.Context) error {
	ctx := c.Request().Context()

	var u model.User
	err := h.app.DB().NewSelect().Model(&u).Where("id = ?", 1).Scan(ctx)
	if err != nil {
		return err
	}

	props := user.ShowProps{
		Title: "this is user page title",
		User:  u,
	}

	return templ.Render(c, user.Show(props))
}

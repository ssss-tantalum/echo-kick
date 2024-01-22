package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/echo-kick/internal/core"
	"github.com/ssss-tantalum/echo-kick/internal/model"
	"github.com/ssss-tantalum/echo-kick/pkg/templ"
)

type UserHandler struct {
	app *core.App
}

func NewUserHadler(app *core.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) ShowUser(c echo.Context) error {
	ctx := c.Request().Context()

	var user model.User
	err := h.app.DB().NewSelect().Model(&user).Where("id = ?", 1).Scan(ctx)
	if err != nil {
		return err
	}

	content := fmt.Sprintf(`
	<div class ="flex flex-col text-center">
		<h2>Yappi~🎉 You got HTMX-GET (/api/user) request!</h2>
		<ul>
			<li class="font-thin">Created by %s</li>
			<li class="font-thin">%s</li>
		</ul>
	</div>
	`, user.Username, user.Email)

	return templ.RenderHtmx(c, content)
}

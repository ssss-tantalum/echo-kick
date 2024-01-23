package view

import (
	"github.com/labstack/echo/v4"

	"github.com/ssss-tantalum/echo-kick/pkg/templ"
	"github.com/ssss-tantalum/echo-kick/web/views/pages/index"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h IndexHandler) IndexView(c echo.Context) error {
	props := index.IndexProps{
		Title: "Echo-Kick",
	}

	return templ.Render(c, index.Show(props))
}

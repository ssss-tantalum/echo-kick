package templ

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return component.Render(c.Request().Context(), c.Response())
}

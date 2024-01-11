package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"io"
)

type TemplRenderer struct {
}

func (t *TemplRenderer) Render(w io.Writer, s string, data any, c echo.Context) error {
	if template, ok := data.(templ.Component); ok {
		return layout(s).Render(templ.WithChildren(c.Request().Context(), template), w)
	}

	return echo.NewHTTPError(500, "could not render template "+s)
}

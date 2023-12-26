package main

import (
	"context"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"io"
)

type TemplRenderer struct {
}

func (t *TemplRenderer) Render(w io.Writer, s string, data any, _ echo.Context) error {
	if template, ok := data.(templ.Component); ok {
		return template.Render(context.Background(), w)
	}

	return echo.NewHTTPError(500, "could not render template "+s)
}

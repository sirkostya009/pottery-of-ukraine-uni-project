package main

import (
	"context"
	"errors"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

type TemplRenderer struct {
	Code      int
	Component templ.Component
}

func (t *TemplRenderer) Render(writer http.ResponseWriter) error {
	t.WriteContentType(writer)
	writer.WriteHeader(t.Code)
	if t.Component == nil {
		return errors.New("no component to render")
	}
	_ = t.Component.Render(context.Background(), writer)
	return nil
}

func (t *TemplRenderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *TemplRenderer) Instance(name string, data any) render.Render {
	if data, ok := data.(templ.Component); ok {
		return &TemplRenderer{t.Code, data}
	}

	return nil
}

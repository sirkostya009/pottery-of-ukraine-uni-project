package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	e := echo.New()
	e.Renderer = &TemplRenderer{}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${latency_human} | ${status} | ${method} @ [${remote_ip}] : ${path}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "home", home())
	})

	e.GET("/glossary", func(c echo.Context) error {
		return c.Render(200, "glossary", glossary())
	})

	e.GET("/decor", func(c echo.Context) error {
		return c.Render(200, "decor", decor())
	})

	e.GET("/process", func(c echo.Context) error {
		return c.Render(200, "process", process())
	})

	e.Static("/static", "static")

	e.File("/favicon.ico", "static/лого.ico")

	_ = e.Start(":" + os.Getenv("PORT"))
}

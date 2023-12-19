package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

var ignoredFiletypes = []string{".go", ".md", ".gitignore", ".sum", ".mod", ".templ"}

func main() {
	e := gin.Default()
	e.HTMLRender = &TemplRenderer{}

	e.GET("/", func(c *gin.Context) {
		c.HTML(200, "", index())
	})

	e.GET("/glossary", func(c *gin.Context) {
		c.HTML(200, "", glossary())
	})

	e.GET("/decor", func(c *gin.Context) {
		c.HTML(200, "", decor())
	})

	e.GET("/process", func(c *gin.Context) {
		c.HTML(200, "", process())
	})

	e.GET("/static/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		for _, filetype := range ignoredFiletypes {
			if strings.HasSuffix(filename, filetype) {
				c.AbortWithStatus(403)
				return
			}
		}
		c.File(filename)
	})

	e.GET("/favicon.ico", func(c *gin.Context) {
		c.File("лого.ico")
	})

	_ = e.Run(":" + os.Getenv("PORT"))
}

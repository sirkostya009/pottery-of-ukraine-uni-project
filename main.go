package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	r.HTMLRender = &TemplRenderer{}

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "", index())
	})

	r.GET("/glossary", func(c *gin.Context) {
		c.HTML(200, "", glossary())
	})

	_ = r.Run(":" + os.Getenv("PORT"))
}

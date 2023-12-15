package main

import (
	"fmt"
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

	r.GET("/stylesheet.css", func(c *gin.Context) {
		_, err := c.Writer.WriteString(stylesheet)
		if err != nil {
			fmt.Println("Failed to write stylesheet", err)
			_ = c.AbortWithError(500, err)
		}
	})

	_ = r.Run(":" + os.Getenv("PORT"))
}

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

	r.GET("/stylesheet.go", func(c *gin.Context) {
		_, err := c.Writer.WriteString(stylesheet)
		if err != nil {
			fmt.Println("Error writing stylesheet.go")
		}
	})

	_ = r.Run(":" + os.Getenv("PORT"))
}

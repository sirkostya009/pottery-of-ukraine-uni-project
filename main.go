package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = &TemplRenderer{}
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(200, "index.html", index())
	})
	_ = r.Run(":8080")
}

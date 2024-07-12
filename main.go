package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Index/index.tmpl", gin.H{
			"title": "Everyone",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Users/index.tmpl", gin.H{
			"user": "James",
		})
	})
	r.Run("localhost:9080")
}

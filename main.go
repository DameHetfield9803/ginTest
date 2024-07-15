package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
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
	r.Run("localhost:3000")
}

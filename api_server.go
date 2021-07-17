package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/freecracy/a/content"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func apiServer(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	r.GET("/getfiles", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"list": content.GetAllFile("./docs"),
		})
	})
	r.GET("/html", func(c *gin.Context) {
		c.Header("Content-Type", "text/html;charset=utf-8")
		file, _ := os.Open("./docs/2021年07月17日.md")
		body, _ := ioutil.ReadAll(file)
		content := blackfriday.Run(body)
		c.String(http.StatusOK, string(content))
	})
	r.NoRoute(func(c *gin.Context) {
		c.File("./ui/dist/index.html")
	})
}

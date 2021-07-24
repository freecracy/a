package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

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
		file, _ := os.Open("./docs/README.md")
		body, _ := ioutil.ReadAll(file)
		content := blackfriday.Run(body)
		c.String(http.StatusOK, string(content))
	})
	r.NoRoute(func(c *gin.Context) {
		c.File("./ui/dist/index.html")
	})
	r.GET("/login", func(c *gin.Context) {
		local, _ := time.LoadLocation("Asia/Shanghai")
		s := time.Now().Local().In(local).Format(time.RFC3339[:7])
		password := fmt.Sprintf("%x", md5.Sum([]byte(s+".tk")))
		if c.PostForm("username") != "root" ||
			c.PostForm("password") != password[0:4]+password[len(password)-4:] {
			c.JSON(http.StatusOK, gin.H{
				"code":    "10001",
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    "10000",
			"message": "success",
			"data": map[string]interface{}{
				"status": true,
			},
		})
	})
}

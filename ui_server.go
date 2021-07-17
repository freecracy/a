package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func uiServer(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.File("./ui/dist/index.html")
	})
	r.StaticFS("/assets", http.Dir("./ui/dist/assets"))
}

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

var result string

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error")
		}
	}()
	go crontab()
	server()
}

func crontab() {
	c := cron.New()
	// _ = c.AddFunc("00 * * * *", func() {
	// 	toutiao.Daily()
	// })
	c.Run()
}

func server() {
	r := gin.Default()
	apiServer(r)
	uiServer(r)
	_ = r.Run()
}

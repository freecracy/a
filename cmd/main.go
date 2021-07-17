package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Toutiao struct {
	gorm.Model
	Title string
	Url   string
}

func main() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//_ = db.AutoMigrate(&Toutiao{})
	//db.Create(&Toutiao{Title: "baidu", Url: "https://www.baidu.com"})
	var t Toutiao
	db.First(&t, 1)
	fmt.Println(t)

}

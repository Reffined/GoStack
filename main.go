package main

import (
	"GoStack/lib/html"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	title := html.H("Home", 1)

	title.AddClass("text")

	html.Page(r).
		AddName("homePage").
		AddChild(title).
		AddStyle("./assets/css/test.css")

	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

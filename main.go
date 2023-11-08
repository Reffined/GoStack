package main

import (
	"GoStack/lib/html"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	title := html.H("Home", 1).
		AddClass("col-fluid py-5")

	row := html.Div().
		AddClass("container text-center").
		AddChild(title)

	html.Page(r).
		AddName("homePage").
		AddChild(row).
		AddStyle("./assets/scss/main.css")

	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

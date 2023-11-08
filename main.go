package main

import (
	"GoStack/lib/html"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	title := html.H("Home", 1)
	title.AddStyle("./assets/css/test.css")
	title.AddClass("text")

	form := html.Form("/login").AddChild(html.Input("text", "username", ""))
	btn := html.Button("submit").AddClass("btn btn-primary")
	form.AddChild(btn)

	html.Page(r).
		AddName("homePage").
		AddChild(title).
		AddChild(form)
	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

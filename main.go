package main

import (
	"GoStack/lib/html"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	title := html.H("Home", 1)
	title.AddStyle("./test.css")
	title.AddClass("text")

	list := html.Ul(html.P("item1"),
		html.P("item2"),
		html.P("item3")).
		AddClass("text")
	img := html.Img("/assets/train-ride.png", "human")

	html.Page(r).
		AddName("homePage").
		AddChild(title).
		AddChild(list).
		AddChild(img)

	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

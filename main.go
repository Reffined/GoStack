package main

import (
	"GoStack/html"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	html.Page(r).
		AddName("home").
		AddChild(html.Div().
			AddChild(html.P("Hello World").
				AddClass("text").
				AddStyle("./test.css")))

	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"GoStack/html"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	out := html.Div().
		AddRouter(r).
		AddName("main").
		AddChild(html.Div().
			AddRouter(r).
			AddName("inner"))

	r.GET("/", func(context *gin.Context) {
		out.Render(context.Writer)
		return
	})

	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

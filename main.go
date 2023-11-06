package main

import (
	"GoStack/lib"
	"GoStack/lib/html"
	"GoStack/lib/htmx"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	h := html.H("Hello World", 1).
		AddRouter(r).
		AddName("msg").
		AddClass("text").
		AddStyle("./test.css")

	a := html.A("Home", "/")

	html.Page(r).
		AddName("home").
		AddChild(html.Div().
			AddName("section").
			AddChild(h).
			AddChild(a))

	h.AddEndpoint(&lib.Endpoint{
		Get: func(context *gin.Context) {
			html.H("Hello World", 2).
				AddRouter(r).
				AddName("msg").
				AddClass("text").
				Render(context.Writer, nil)
		},
		Post:   nil,
		Put:    nil,
		Delete: nil,
	}).
		AddHtmx(&htmx.Hx{
			Get:     "this",
			Trigger: "mouseenter",
			Swap:    "outerHTML",
			Target:  "this",
		})

	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

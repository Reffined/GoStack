package main

import (
	"GoStack/lib"
	"GoStack/lib/html"
	"GoStack/lib/htmx"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	p := html.P("Hello World").
		AddRouter(r).
		AddName("msg").
		AddClass("text").
		AddStyle("./test.css")

	html.Page(r).
		AddName("home").
		AddChild(html.Div().
			AddName("section").
			AddChild(p))

	p.AddEndpoint(&lib.Endpoint{
		Get: func(context *gin.Context) {
			html.P("Hello Alireza").
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

package main

import (
	"GoStack/lib"
	html2 "GoStack/lib/html"
	"GoStack/lib/htmx"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	hx := htmx.Hx{
		Get:     "/home/section/msg",
		Trigger: "mouseenter",
		Swap:    "outerHtml",
	}
	p := html2.P("Hello World", &hx).
		AddRouter(r).
		AddName("msg").
		AddClass("text").
		AddStyle("./test.css")

	html2.Page(r).
		AddName("home").
		AddChild(html2.Div(nil).
			AddName("section").
			AddChild(p))

	p.AddEndpoint(&lib.Endpoint{
		Get: func(context *gin.Context) {
			html2.P("Hello Alireza", nil).
				AddRouter(r).
				AddName("msg").
				AddClass("text").
				Render(context.Writer, nil)
		},
		Post:   nil,
		Put:    nil,
		Delete: nil,
	})
	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

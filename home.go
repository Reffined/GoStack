package main

import (
	"GoStack/lib"
	"github.com/gin-gonic/gin"
)

type Div struct {
	*lib.BaseComponent
}

func NewDiv(name string, router *gin.Engine, parent lib.Component, children []lib.Component, endPoint *lib.Endpoint) *Div {
	var n Div
	n.BaseComponent = lib.NewComponent(name, router, parent, children, endPoint)
	n.OuterHtml = "<div>{{.}}</div>"
	return &n
}

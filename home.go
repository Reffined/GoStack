package main

import (
	"github.com/gin-gonic/gin"
	"helphub/lib"
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

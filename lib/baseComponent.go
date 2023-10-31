package lib

import (
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"text/template"
)

type BaseComponent struct {
	name       string
	parent     Component
	router     *gin.Engine
	routes     []string
	endPoint   *Endpoint
	components []Component
	OuterHtml  string
}

func (b *BaseComponent) AddChild(c Component) {
	b.components = append(b.components, c)
}

func (b *BaseComponent) Endpoint() *Endpoint {
	return b.endPoint
}

func NewComponent(name string, router *gin.Engine, parent Component, children []Component, endPoint *Endpoint) *BaseComponent {
	var b BaseComponent
	b.router = router
	b.parent = parent
	b.name = name
	b.components = children
	b.endPoint = endPoint

	group := router.Group(b.Route())
	if endPoint != nil {
		if endPoint.Get != nil {
			group.GET("/", endPoint.Get)
		}

		if endPoint.Post != nil {
			group.POST("/", endPoint.Post)
		}

		if endPoint.Put != nil {
			group.PUT("/", endPoint.Put)
		}

		if endPoint.Delete != nil {
			group.DELETE("/", endPoint.Delete)
		}
	} else {
		group.GET("/", func(context *gin.Context) {
			b.Render(context.Writer, "")
		})
	}
	return &b
}

func (b *BaseComponent) Parent() Component {
	return b.parent
}

func (b *BaseComponent) Render(writer io.Writer, outerHtml string) {
	var tmpl string

	if b.OuterHtml == "" {
		tmpl = `{{ . }}`
	} else {
		tmpl = b.OuterHtml
	}

	base, err := template.New("base.html").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	sb := strings.Builder{}
	for _, v := range b.components {
		v.Render(&sb, "")
	}

	err = base.Execute(writer, sb.String())
	if err != nil {
		panic(err)
	}
}

func (b *BaseComponent) Name() string {
	return b.name
}

func (b *BaseComponent) Routes() []string {
	return b.routes
}

func (b *BaseComponent) Route() string {
	if b.parent != nil {
		return b.parent.Route() + "/" + b.name
	} else {
		return "/" + b.name
	}
}

func (b *BaseComponent) Components() []Component {
	return b.components
}

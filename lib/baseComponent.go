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

func (b *BaseComponent) AddOuterHtml(h string) *BaseComponent {
	b.OuterHtml = h
	return b
}

func (b *BaseComponent) AddEndpoint(e *Endpoint) *BaseComponent {
	b.endPoint = e
	if b.router == nil {
		panic("add a router before adding an endpoint")
	}
	g := b.router.Group(b.Route())

	if e.Get != nil {
		g.GET("", e.Get)
	}

	if e.Post != nil {
		g.POST("", e.Post)
	}

	if e.Put != nil {
		g.PUT("", e.Put)
	}

	if e.Delete != nil {
		g.DELETE("", e.Delete)
	}

	return b
}

func (b *BaseComponent) AddRouter(r *gin.Engine) *BaseComponent {
	b.router = r
	return b
}

func (b *BaseComponent) addParent(p Component) *BaseComponent {
	b.parent = p
	return b
}

func (b *BaseComponent) AddName(name string) *BaseComponent {
	b.name = name
	return b
}

func (b *BaseComponent) AddChild(c Component) *BaseComponent {
	bc, ok := c.(*BaseComponent)
	if ok {
		bc.addParent(b)
	}
	b.components = append(b.components, c)

	return b
}

func (b *BaseComponent) Endpoint() *Endpoint {
	return b.endPoint
}

func (b *BaseComponent) Parent() Component {
	return b.parent
}

func (b *BaseComponent) Render(writer io.Writer) {
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
		v.Render(&sb)
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

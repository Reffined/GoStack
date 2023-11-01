package html

import (
	"GoStack/lib"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"text/template"
)

var templatePage string

type GSPage struct {
	name     string
	children []lib.Component
	router   *gin.Engine
	parent   *GSPage
	endpoint *lib.Endpoint
}

func (p *GSPage) AddParent(component lib.Component) {
	v, ok := component.(*GSPage)
	if ok {
		p.parent = v
	} else {
		p.parent = nil
	}
}

func (p *GSPage) Parent() lib.Component {
	return p.parent
}

func (p *GSPage) Routes() []string {
	routes := make([]string, 0, 10)
	for _, v := range p.children {
		routes = append(routes, v.Route())
	}
	return routes
}

func (p *GSPage) Components() []lib.Component {
	return p.children
}

func (p *GSPage) Endpoint() *lib.Endpoint {
	return p.endpoint
}

func (p *GSPage) Name() string {
	return p.name
}

func (p *GSPage) Style() string {
	return ""
}

func (p *GSPage) AddName(name string) *GSPage {
	p.name = name
	return p
}

func (p *GSPage) AddChild(c lib.Component) *GSPage {
	p.children = append(p.children, c)
	c.AddParent(p)
	return p
}

func Page(router *gin.Engine) *GSPage {
	if templatePage == "" {
		file, err := os.Open("./lib/html/pageTemplate.html")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		stat, _ := file.Stat()
		size := stat.Size()
		buff := make([]byte, size)
		_, err = file.Read(buff)
		if err != nil {
			panic(err)
		}

		templatePage = string(buff)
	}

	var page GSPage
	page.router = router
	handler := func(ctx *gin.Context) {
		page.Render(ctx.Writer, nil)
	}
	page.endpoint = &lib.Endpoint{
		Get:    handler,
		Post:   nil,
		Put:    nil,
		Delete: nil,
	}
	router.GET(page.Route(), handler)
	return &page
}

func (p *GSPage) Route() string {
	if p.parent == nil {
		return "/" + p.name
	} else {
		return p.parent.Route() + "/" + p.name
	}
}

func (p *GSPage) Render(writer io.Writer, style io.Writer) {
	t, err := template.New("Page.html").Parse(templatePage)
	if err != nil {
		panic(err)
	}

	tags := strings.Builder{}
	styles := strings.Builder{}
	for _, v := range p.children {
		if _, ok := v.(*GSPage); ok {
			continue
		}
		v.Render(&tags, &styles)
	}
	data := struct {
		Title string
		Body  string
		Style string
	}{
		p.name,
		tags.String(),
		styles.String(),
	}

	err = t.Execute(writer, data)
	if err != nil {
		panic(err)
	}
}

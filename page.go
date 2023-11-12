package GoStack

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"text/template"
)

var templatePage = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <script crossorigin=\"anonymous\" integrity=\"sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni\" src=\"https://unpkg.com/htmx.org@1.9.6\"></script>\n    <script src=\"https://unpkg.com/htmx.org/dist/ext/json-enc.js\"></script>\n    <link crossorigin=\"anonymous\" href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css\" integrity=\"sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN\" rel=\"stylesheet\">\n    <script crossorigin=\"anonymous\" integrity=\"sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL\" src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js\"></script>\n    <meta charset=\"UTF-8\">\n    <title>{{.Title}}</title>\n</head>\n<header>\n\n</header>\n<body>\n    {{ .Body }}\n</body>\n<style>\n    {{ .Style }}\n</style>\n</html>"

// GSPage is a web page and automatically assigns a route to itself that can be visited by a get request.
type GSPage struct {
	name     string
	children []Component
	router   *gin.Engine
	parent   *GSPage
	endpoint *Endpoint
}

// AddStyle Takes a css file path and adds the content to the <style> tag of that page
func (p *GSPage) AddStyle(file string) *GSPage {
	if p.children == nil || len(p.children) == 0 {
		return p
	}
	baseComp, ok := p.children[0].(*BaseComponent)
	if ok {
		baseComp.AddStyle(file)
	}
	return p
}

// AddParent is used by both the AddChild from other pages and users of the package to add a parent
// for this GSPage and is used for forming a virtual DOM and routes
func (p *GSPage) AddParent(component Component) {
	v, ok := component.(*GSPage)
	if ok {
		p.parent = v
	} else {
		p.parent = nil
	}
}

func (p *GSPage) Parent() Component {
	return p.parent
}

func (p *GSPage) Routes() []string {
	routes := make([]string, 0, 10)
	for _, v := range p.children {
		routes = append(routes, v.Route())
	}
	return routes
}

func (p *GSPage) Components() []Component {
	return p.children
}

func (p *GSPage) Endpoint() *Endpoint {
	return p.endpoint
}

func (p *GSPage) Name() string {
	return p.name
}

func (p *GSPage) Style() string {
	return ""
}

// AddName adds a name that is used for routing and can be seen in the url
func (p *GSPage) AddName(name string) *GSPage {
	p.name = name
	return p
}

// AddChild adds a Component as a child to this GSPage and sets itself as the parent
func (p *GSPage) AddChild(c Component) *GSPage {
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
	page.endpoint = &Endpoint{
		Get:    handler,
		Post:   nil,
		Put:    nil,
		Delete: nil,
	}
	router.GET(page.Route(), handler)
	return &page
}

// Route generates the absolute path of the GSPage
func (p *GSPage) Route() string {
	if p.parent == nil {
		return "/" + p.name
	} else {
		return p.parent.Route() + "/" + p.name
	}
}

// Render creates the final html by recursively traversing the tree and writes the result to an io.Writer
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

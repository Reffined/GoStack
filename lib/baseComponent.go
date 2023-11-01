package lib

import (
	"GoStack/lib/htmx"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"reflect"
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
	style      []byte
	classes    []string
	id         string
	Htmx       *htmx.Hx
	TagName    string
	Body       string
}

func (b *BaseComponent) GetClass() []string {
	return b.classes
}
func (b *BaseComponent) AddClass(c ...string) *BaseComponent {
	b.classes = append(b.classes, c...)
	return b
}

func (b *BaseComponent) AddHtmx(hx *htmx.Hx) *BaseComponent {
	b.Htmx = hx
	return b
}

func (b *BaseComponent) AddId(i string) *BaseComponent {
	b.id = i
	return b
}

func (b *BaseComponent) AddStyle(cssPath string) *BaseComponent {
	file, err := os.Open(cssPath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	stat, _ := file.Stat()
	size := stat.Size()
	buff := make([]byte, size)
	_, err = file.Read(buff)
	if err != nil {
		panic(err)
	}
	b.style = buff
	return b
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
		b.routes = append(b.routes, "GET:"+b.Route())
	}

	if e.Post != nil {
		g.POST("", e.Post)
		b.routes = append(b.routes, "POST:"+b.Route())
	}

	if e.Put != nil {
		g.PUT("", e.Put)
		b.routes = append(b.routes, "PUT:"+b.Route())
	}

	if e.Delete != nil {
		g.DELETE("", e.Delete)
		b.routes = append(b.routes, "DELETE:"+b.Route())
	}

	return b
}

func (b *BaseComponent) AddRouter(r *gin.Engine) *BaseComponent {
	b.router = r
	return b
}

func (b *BaseComponent) AddParent(p Component) {
	b.parent = p
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
	c.AddParent(b)
	b.components = append(b.components, c)

	return b
}

func (b *BaseComponent) Endpoint() *Endpoint {
	return b.endPoint
}

func (b *BaseComponent) Parent() Component {
	return b.parent
}

func (b *BaseComponent) Render(writer io.Writer, styles io.Writer) {
	b.makeHtml(b.TagName, b.Body, b.Htmx)
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

	if b.style != nil {
		_, err = styles.Write(b.style)
		if err != nil {
			panic(err)
		}
		_, err = styles.Write([]byte("\n"))
		if err != nil {
			panic(err)
		}
	}
	for _, v := range b.components {
		v.Render(&sb, styles)
	}

	classes := strings.Builder{}
	for _, v := range b.classes {
		classes.WriteString(v)
		classes.WriteString(" ")
	}

	err = base.Execute(writer, struct {
		Id    string
		Class string
		Body  string
	}{b.id, classes.String(), sb.String()})
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

func (b *BaseComponent) Style() string {
	return string(b.style)
}

func (b *BaseComponent) makeHtml(tagName string, body string, htmx *htmx.Hx) string {
	hx := strings.Builder{}
	if htmx != nil {
		ty := reflect.TypeOf(*htmx)
		for i := 0; i < 3; i++ {
			field := ty.Field(i)
			fValue := reflect.ValueOf(*htmx).FieldByIndex([]int{i}).String()
			var attr string
			if fValue != "" {
				fName := strings.ToLower(field.Name)
				if fValue == "this" {
					attr = fmt.Sprintf(`hx-%s="%s" `, fName, b.Route())
				} else {
					attr = fmt.Sprintf(`hx-%s="%s" `, fName, fValue)
				}
				hx.WriteString(attr)
			}
		}
		for i := 3; i < ty.NumField(); i++ {
			field := ty.Field(i)
			fValue := reflect.ValueOf(*htmx).FieldByIndex([]int{i}).String()
			if fValue != "" {
				fName := strings.ToLower(field.Name)
				attr := fmt.Sprintf(`hx-%s="%s" `, fName, fValue)
				hx.WriteString(attr)
			}
		}
	}

	var str string
	if body == "" {
		if hx.String() == "" {
			str = fmt.Sprintf(`<%s class="{{ .Class }}" id="{{ .Id }}">{{ .Body }}</%s>`, tagName, tagName)
		} else {
			str = fmt.Sprintf(`<%s %s class="{{ .Class }}" id="{{ .Id }}">{{ .Body }}</%s>`, tagName, hx.String(), tagName)
		}
	} else {
		if hx.String() == "" {
			str = fmt.Sprintf(`<%s class="{{ .Class }}" id="{{ .Id }}">%s</%s>`, tagName, body, tagName)
		} else {
			str = fmt.Sprintf(`<%s %s class="{{ .Class }}" id="{{ .Id }}">%s</%s>`, tagName, hx.String(), body, tagName)
		}
	}
	b.AddOuterHtml(str)
	return str
}

package GoStack

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type BaseComponent struct {
	Bname       string
	Bparent     Component
	Brouter     *gin.Engine
	Broutes     []string
	BendPoint   *Endpoint
	Bcomponents []Component
	BOuterHtml  string
	Bstyle      []byte
	Bclasses    []string
	Bid         string
	BHtmx       *Hx
	TagName     string
	Body        string
	Attributes  []string
}

func (b *BaseComponent) AddAttributes(attrs ...string) {
	b.Attributes = attrs
}
func (b *BaseComponent) GetClass() []string {
	return b.Bclasses
}
func (b *BaseComponent) AddClass(c ...string) *BaseComponent {
	b.Bclasses = append(b.Bclasses, c...)
	return b
}
func (b *BaseComponent) AddHtmx(hx *Hx) *BaseComponent {
	b.BHtmx = hx
	return b
}
func (b *BaseComponent) AddId(i string) *BaseComponent {
	b.Bid = i
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
	b.Bstyle = buff
	return b
}
func (b *BaseComponent) AddOuterHtml(h string) *BaseComponent {
	b.BOuterHtml = h
	return b
}
func (b *BaseComponent) AddEndpoint(e *Endpoint) *BaseComponent {
	b.BendPoint = e
	if b.Brouter == nil {
		panic("add a Brouter before adding an endpoint")
	}
	g := b.Brouter.Group(b.Route())

	if e.Get != nil {
		g.GET("", e.Get)
		b.Broutes = append(b.Broutes, "GET:"+b.Route())
	}

	if e.Post != nil {
		g.POST("", e.Post)
		b.Broutes = append(b.Broutes, "POST:"+b.Route())
	}

	if e.Put != nil {
		g.PUT("", e.Put)
		b.Broutes = append(b.Broutes, "PUT:"+b.Route())
	}

	if e.Delete != nil {
		g.DELETE("", e.Delete)
		b.Broutes = append(b.Broutes, "DELETE:"+b.Route())
	}

	return b
}
func (b *BaseComponent) AddRouter(r *gin.Engine) *BaseComponent {
	b.Brouter = r
	return b
}
func (b *BaseComponent) AddParent(p Component) {
	b.Bparent = p
}
func (b *BaseComponent) addParent(p Component) *BaseComponent {
	b.Bparent = p
	return b
}
func (b *BaseComponent) AddName(name string) *BaseComponent {
	b.Bname = name
	return b
}

func (b *BaseComponent) AddChild(c Component) *BaseComponent {
	c.AddParent(b)
	b.Bcomponents = append(b.Bcomponents, c)

	return b
}

func (b *BaseComponent) Endpoint() *Endpoint {
	return b.BendPoint
}

func (b *BaseComponent) Parent() Component {
	return b.Bparent
}

func (b *BaseComponent) Render(writer io.Writer, styles io.Writer) {
	b.makeHtml(b.TagName, b.Body, b.BHtmx)
	var tmpl string

	if b.BOuterHtml == "" {
		tmpl = `{{ . }}`
	} else {
		tmpl = b.BOuterHtml
	}

	base, err := template.New("base.html").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	sb := strings.Builder{}

	if b.Bstyle != nil {
		_, err = styles.Write(b.Bstyle)
		if err != nil {
			panic(err)
		}
		_, err = styles.Write([]byte("\n"))
		if err != nil {
			panic(err)
		}
	}
	for _, v := range b.Bcomponents {
		v.Render(&sb, styles)
	}

	classes := strings.Builder{}
	for _, v := range b.Bclasses {
		classes.WriteString(v)
		classes.WriteString(" ")
	}

	err = base.Execute(writer, struct {
		Id    string
		Class string
		Body  string
	}{b.Bid, classes.String(), sb.String()})
	if err != nil {
		panic(err)
	}
}

func (b *BaseComponent) Name() string {
	return b.Bname
}

func (b *BaseComponent) Routes() []string {
	return b.Broutes
}

func (b *BaseComponent) Route() string {
	if b.Bparent != nil {
		return b.Bparent.Route() + "/" + b.Bname
	} else {
		return "/" + b.Bname
	}
}

func (b *BaseComponent) Components() []Component {
	return b.Bcomponents
}

func (b *BaseComponent) Style() string {
	return string(b.Bstyle)
}

func (b *BaseComponent) makeHtml(tagName string, body string, htmx *Hx) string {
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

	attr := strings.Builder{}
	if b.Attributes != nil && len(b.Attributes) > 0 {
		for _, v := range b.Attributes {
			attr.WriteString(v + " ")
		}
	}

	var str string
	if body == "" {
		str = fmt.Sprintf(`<%s %s %s class="{{ .Class }}" id="{{ .Id }}">{{ .Body }}</%s>`, tagName, attr.String(), hx.String(), tagName)
	} else {
		str = fmt.Sprintf(`<%s %s %s class="{{ .Class }}" id="{{ .Id }}">%s</%s>`, tagName, attr.String(), hx.String(), body, tagName)
	}
	b.AddOuterHtml(str)
	return str
}

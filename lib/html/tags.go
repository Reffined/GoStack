package html

import (
	"GoStack/lib"
	"GoStack/lib/htmx"
	"fmt"
	"reflect"
	"strings"
)

type Tag struct {
	lib.BaseComponent
	tagName string
}

func (t *Tag) makeHtml(tagName string, body string, htmx *htmx.Hx) string {
	t.tagName = tagName
	hx := strings.Builder{}
	if htmx != nil {
		ty := reflect.TypeOf(*htmx)
		for i := 0; i < ty.NumField(); i++ {
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
	t.AddOuterHtml(str)
	return str
}

func Div(hx *htmx.Hx) *Tag {
	tag := Tag{}
	tag.makeHtml("div", "", hx)
	return &tag
}

func P(text string, hx *htmx.Hx) *Tag {
	tag := Tag{}
	tag.makeHtml("p", text, hx)
	return &tag
}

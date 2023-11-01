package html

import (
	"GoStack/lib"
	"fmt"
)

type Tag struct {
	lib.BaseComponent
	tagName string
}

func (t *Tag) makeHtml(tagName string, body string) string {
	t.tagName = tagName
	var str string
	if body == "" {
		str = fmt.Sprintf(`<%s class="{{ .Class }}" id="{{ .Id }}">{{ .Body }}</%s>`, tagName, tagName)
	} else {
		str = fmt.Sprintf(`<%s class="{{ .Class }}" id="{{ .Id }}">%s</%s>`, tagName, body, tagName)
	}
	t.AddOuterHtml(str)
	return str
}

func Div() *Tag {
	tag := Tag{}
	tag.makeHtml("div", "")
	return &tag
}

func P(text string) *Tag {
	tag := Tag{}
	tag.makeHtml("p", text)
	return &tag
}

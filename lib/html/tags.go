package html

import (
	"GoStack/lib"
	"fmt"
)

func Div() *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "div"
	return &tag
}

func P(text string) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "p"
	tag.Body = text
	return &tag
}

func H(text string, lvl int) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = fmt.Sprintf("h%d", lvl)
	tag.Body = text
	return &tag
}

func A(text string, href string) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "a"
	tag.Body = text
	tag.AddAttributes("href=" + href)
	return &tag
}

func Li(child *lib.BaseComponent) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "li"
	if child != nil {
		tag.AddChild(child)
	}
	return &tag
}

func Ul(children ...*lib.BaseComponent) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "ul"
	for _, v := range children {
		tag.AddChild(v)
	}
	return &tag
}

func Ol(children ...*lib.BaseComponent) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "ol"
	for _, v := range children {
		tag.AddChild(v)
	}
	return &tag
}

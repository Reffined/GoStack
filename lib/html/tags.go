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
		li := Li(v)
		tag.AddChild(li)
	}
	return &tag
}

func Ol(children ...*lib.BaseComponent) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "ol"
	for _, v := range children {
		li := Li(v)
		tag.AddChild(li)
	}
	return &tag
}

func Img(src string, alt string) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "img"
	tag.AddAttributes(fmt.Sprintf("src=%s alt=%s", src, alt))
	return &tag
}

func Form(action string) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "form"
	tag.AddAttributes(fmt.Sprintf("action=\"%s\"", action))
	return &tag
}

func Input(inputType string, name string, value string) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "input"
	tag.AddAttributes(fmt.Sprintf("type=\"%s\" name=\"%s\" value=\"%s\"", inputType, name, value))
	return &tag
}

func Button(text string) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "button"
	tag.Body = text
	return &tag
}

func Option(text string) *lib.BaseComponent {
	tag := lib.BaseComponent{}
	tag.TagName = "option"
	tag.Body = text
	return &tag
}

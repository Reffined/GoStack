package GoStack

import (
	"fmt"
)

func Div() *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "div"
	return &tag
}

func P(text string) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "p"
	tag.Body = text
	return &tag
}

func H(text string, lvl int) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = fmt.Sprintf("h%d", lvl)
	tag.Body = text
	return &tag
}

func A(text string, href string) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "a"
	tag.Body = text
	tag.AddAttributes("href=" + href)
	return &tag
}

func Li(child *BaseComponent) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "li"
	if child != nil {
		tag.AddChild(child)
	}
	return &tag
}

func Ul(children ...*BaseComponent) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "ul"
	for _, v := range children {
		li := Li(v)
		tag.AddChild(li)
	}
	return &tag
}

func Ol(children ...*BaseComponent) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "ol"
	for _, v := range children {
		li := Li(v)
		tag.AddChild(li)
	}
	return &tag
}

func Img(src string, alt string) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "img"
	tag.AddAttributes(fmt.Sprintf("src=%s alt=%s", src, alt))
	return &tag
}

func Form(action string) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "form"
	tag.AddAttributes(fmt.Sprintf("action=\"%s\"", action))
	return &tag
}

func Input(inputType string, name string, value string) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "input"
	tag.AddAttributes(fmt.Sprintf("type=\"%s\" name=\"%s\" value=\"%s\"", inputType, name, value))
	return &tag
}

func Button(text string) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "button"
	tag.Body = text
	return &tag
}

func Option(text string) *BaseComponent {
	tag := BaseComponent{}
	tag.TagName = "option"
	tag.Body = text
	return &tag
}

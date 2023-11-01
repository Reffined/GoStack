package html

import "GoStack/lib"

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

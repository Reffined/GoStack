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

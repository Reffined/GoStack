package html

import "GoStack/lib"

type GSDiv struct {
	lib.BaseComponent
}

func Div() *GSDiv {
	var d GSDiv
	d.AddOuterHtml(`<div class="{{ .Class }}">{{ .Body }}</div>`)
	return &d
}

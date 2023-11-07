package html

import (
	"strings"
	"testing"
)

func TestDiv(t *testing.T) {
	tag := Div().AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<div   class=\"\" id=\"\"></div>" {
			t.Fatal(buff.String())
		}
	})
}

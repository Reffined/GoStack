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

func TestP(t *testing.T) {
	tag := P("Hello").AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<p   class=\"\" id=\"\">Hello</p>" {
			t.Fatal(buff.String())
		}
	})
}

func TestH(t *testing.T) {
	tag := H("Hello", 1).AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<h1   class=\"\" id=\"\">Hello</h1>" {
			t.Fatal(buff.String())
		}
	})
}

func TestA(t *testing.T) {
	tag := A("Hello", "/").AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<a href=/   class=\"\" id=\"\">Hello</a>" {
			t.Fatal(buff.String())
		}
	})
}

func TestLi(t *testing.T) {
	child := P("hello")
	tag := Li(child).AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<li   class=\"\" id=\"\"><p   class=\"\" id=\"\">hello</p></li>" {
			t.Fatal(buff.String())
		}
	})
}

func TestUl(t *testing.T) {
	child := Li(P("hello"))
	tag := Ul(child).AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<ul   class=\"\" id=\"\"><li   class=\"\" id=\"\"><li   class=\"\" id=\"\"><p   class=\"\" id=\"\">hello</p></li></li></ul>" {
			t.Fatal(buff.String())
		}
	})
}

func TestOl(t *testing.T) {
	child := Li(P("hello"))
	tag := Ol(child).AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<ol   class=\"\" id=\"\"><li   class=\"\" id=\"\"><li   class=\"\" id=\"\"><p   class=\"\" id=\"\">hello</p></li></li></ol>" {
			t.Fatal(buff.String())
		}
	})
}

func TestImg(t *testing.T) {
	tag := Img("/assets/img.png", "pic").AddName("main")
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<img src=/assets/img.png alt=pic   class=\"\" id=\"\"></img>" {
			t.Fatal(buff.String())
		}
	})
}

func TestForm(t *testing.T) {
	child := Input("text", "username", "username")
	tag := Form("/login").AddName("main").AddChild(child)
	buff := strings.Builder{}
	tag.Render(&buff, nil)
	t.Run("outerHTML is valid", func(t *testing.T) {
		if buff.String() != "<form action=\"/login\"   class=\"\" id=\"\"><input type=\"text\" name=\"username\" value=\"username\"   class=\"\" id=\"\"></input></form>" {
			t.Fatal(buff.String())
		}
	})
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	out := NewDiv("outer", r, nil, nil, nil)
	out.AddChild(NewDiv("inner", r, out, nil, nil))

	err := r.Run("0.0.0.0:4040")
	if err != nil {
		panic(err)
	}
}

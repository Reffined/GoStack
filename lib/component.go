package lib

import (
	"io"
)

type Component interface {
	Parent() Component
	Render(writer io.Writer)
	Routes() []string
	Components() []Component
	Route() string
	Endpoint() *Endpoint
	Name() string
}

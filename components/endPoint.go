package components

import "github.com/gin-gonic/gin"

type Endpoint struct {
	Get    gin.HandlerFunc
	Post   gin.HandlerFunc
	Put    gin.HandlerFunc
	Delete gin.HandlerFunc
}

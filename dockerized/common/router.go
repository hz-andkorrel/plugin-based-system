package common

import "github.com/gin-gonic/gin"

type Router interface {
	AddGetRoute(path string, handler func(context *gin.Context))
	Run(port string)
}

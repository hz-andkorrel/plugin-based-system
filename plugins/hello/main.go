package main

import (
	"plugins/common"

	"github.com/gin-gonic/gin"
)

type Plugin struct{}

func NewPlugin() common.Plugin {
	return &Plugin{}
}

func (plugin *Plugin) RegisterRoutes(router common.Router) {
	router.AddGetRoute("/hello", func(context *gin.Context) {
		context.String(200, "Hello from Plugin!")
	})
}

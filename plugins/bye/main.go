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
	router.AddGetRoute("/bye", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Goodbye, World!",
		})
	})
}

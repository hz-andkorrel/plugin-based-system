package domain

import (
	"plugins/common"

	"github.com/gin-gonic/gin"
)

type RouteManager struct {
	router *gin.Engine
}

// NewRouteManager creates a new instance of RouteManager.
func NewRouteManager() common.Router {
	return &RouteManager{
		router: gin.Default(),
	}
}

// AddPostRoute adds a new POST route to the routing table.
func (manager *RouteManager) AddPostRoute(path string, handler func(c *gin.Context)) {
	manager.router.POST(path, handler)
}

// Run starts the HTTP server on the specified port.
func (manager *RouteManager) Run(port string) {
	manager.router.Run(port)
}

package main

import (
	"plugins/broker/domain"
)

// The entry point of the application.
// It uses the PluginRegistry to load and interact with plugins.
// First, the router and plugin registry are created.
// The plugins are auto-discovered from the plugins directory before being initialized.
// Finally, the router is started to listen for incoming requests.
func main() {
	router := domain.NewRouteManager()

	registry := domain.NewPluginRegistry()
	registry.AutoDiscoverPlugins()
	registry.InitializePlugins(router)

	router.Run(":8080")
}

package main

import (
	"fmt"
	"plugins/broker/domain"
)

// The entry point of the application.
// It uses the PluginRegistry to load and interact with plugins.
// In the current iteration, this means loading a single plugin and calling its Hello method.
func main() {
	registry := domain.NewPluginRegistry()
	registry.AutoDiscoverPlugins()

	for _, plugin := range registry.Plugins {
		fmt.Println("Loaded plugin:", plugin.Hello())
	}
}

package domain

import "plugins/common"

// PluginRegistry is a minimal registry used by the dockerized broker.
// It intentionally omits legacy .so plugin discovery and only holds
// an in-memory list of plugins (container-backed or otherwise).
type PluginRegistry struct {
	Plugins []common.Plugin
}

// NewPluginRegistry returns an empty registry ready to receive plugins.
func NewPluginRegistry() *PluginRegistry {
	return &PluginRegistry{Plugins: []common.Plugin{}}
}

// InitializePlugins initializes all registered plugins by calling their RegisterRoutes method.
func (registry *PluginRegistry) InitializePlugins(router common.Router) {
	for _, plugin := range registry.Plugins {
		plugin.RegisterRoutes(router)
	}
}

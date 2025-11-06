package domain

import (
	"eventbus/common"
	"os"
	"plugin"
	"strings"
)

// PluginRegistry represents a collection of plugins available in the system.
// It manages the lifecycle and metadata of each plugin.
type PluginRegistry struct {
	PluginDirectory string
	Plugins         []common.Plugin
}

// The PluginRegistry shall be initialized with an empty list of plugins.
// The plugin directory is set to a predefined constant.
func NewPluginRegistry() *PluginRegistry {
	return &PluginRegistry{
		PluginDirectory: "./plugins/",
		Plugins:         []common.Plugin{},
	}
}

// AddPlugin add a new plugin to the registry.
// This is based on the file name of the plugin.
// The plugin directory is assumed to be a predefined constant.
// The .so extension should be provided with the plugin name.
func (registry *PluginRegistry) AddPlugin(fileName string) {
	filePath := registry.PluginDirectory + fileName
	pluginFile, _ := plugin.Open(filePath)

	NewPluginSymbol, _ := pluginFile.Lookup("NewPlugin")
	NewPluginFunc := NewPluginSymbol.(func() common.Plugin)

	registry.Plugins = append(registry.Plugins, NewPluginFunc())
}

// AutoDiscoverPlugins scans the plugin directory for available plugins.
// Only top-level .so within the plugins directory are considered valid plugins.
func (registry *PluginRegistry) AutoDiscoverPlugins() {
	files, _ := os.ReadDir(registry.PluginDirectory)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".so") {
			registry.AddPlugin(file.Name())
		}
	}
}

// InitializePlugins initializes all registered plugins by calling their RegisterRoutes method.
func (registry *PluginRegistry) InitializePlugins(bus common.Bus) {
	for _, plugin := range registry.Plugins {
		plugin.Initialize(bus)
	}
}

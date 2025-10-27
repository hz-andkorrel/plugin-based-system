package domain

import (
	"os"
	"plugin"
	"plugins/common"
	"strings"
)

// PluginRegistry represents a collection of plugins available in the system.
// It manages the lifecycle and metadata of each plugin.
type PluginRegistry struct {
	Plugins         []common.Plugin
	PluginDirectory string
}

// The PluginRegistry shall be initialized with an empty list of plugins.
// The plugin directory is set to a predefined constant.
func NewPluginRegistry() *PluginRegistry {
	return &PluginRegistry{
		Plugins:         []common.Plugin{},
		PluginDirectory: "./plugins/",
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

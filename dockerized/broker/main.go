package main

import (
	"fmt"
	"os"
	"plugins/broker/domain"
)

// The entry point of the application.
// It uses the PluginRegistry to load and interact with plugins.
// First, the router and plugin registry are created.
// The plugins are auto-discovered from the plugins directory before being initialized.
// Additionally, if a containers manifest exists it will be loaded and registered.
// Finally, the router is started to listen for incoming requests.
func main() {
	router := domain.NewRouteManager()

	registry := domain.NewPluginRegistry()

	// Load containerized plugin manifest if present. Check multiple locations and support an env override.
	manifestPath := os.Getenv("PLUGINS_MANIFEST")
	if manifestPath == "" {
		// prefer repo-local path (for running from source), fall back to path used in Docker build image
		manifestPath = "containers.yml"
		if _, err := os.Stat(manifestPath); err != nil {
			// try the Docker builder location where the source was copied during image build
			manifestPath = "/src/containers.yml"
		}
	}

	if _, err := os.Stat(manifestPath); err == nil {
		if containerPlugins, err := domain.LoadContainerPlugins(manifestPath); err == nil {
			// append loaded container plugins to registry
			registry.Plugins = append(registry.Plugins, containerPlugins...)
			fmt.Println("Loaded container plugins from", manifestPath)
		} else {
			fmt.Println("Failed to load container plugins from", manifestPath, ":", err)
		}
	}

	registry.InitializePlugins(router)

	router.Run(":8080")
}

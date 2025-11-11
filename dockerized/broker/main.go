package main

import (
	"fmt"
	"os"
	"plugins/broker/domain"
	"plugins/common"
)

// Entry point for dockerized broker. Uses a simple slice for plugins instead of a registry type.
func main() {
	router := domain.NewRouteManager()

	var registry []common.Plugin

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
			// append loaded container plugins to registry slice
			registry = append(registry, containerPlugins...)
			fmt.Println("Loaded container plugins from", manifestPath)
		} else {
			fmt.Println("Failed to load container plugins from", manifestPath, ":", err)
		}
	}

	// Initialize plugins by registering their routes on the router.
	for _, p := range registry {
		p.RegisterRoutes(router)
	}

	// Expose admin registration route which appends to the registry slice.
	domain.RegisterAdminRoutes(&registry, router)

	router.Run(":8080")
}

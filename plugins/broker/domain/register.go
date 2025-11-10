package domain

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"plugins/common"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
)

type RegisterRequest struct {
	ID      string            `json:"id" yaml:"id"`
	Name    string            `json:"name" yaml:"name"`
	BaseURL string            `json:"base_url" yaml:"base_url"`
	Route   string            `json:"route" yaml:"route"`
	Meta    map[string]string `json:"meta,omitempty" yaml:"meta,omitempty"`
}

// persistContainerPlugins writes the currently registered container plugins to disk
// at the path `plugins/containers.yml` (relative).
func persistContainerPlugins(path string, registry *PluginRegistry) error {
	wrapper := struct {
		Plugins []ContainerPlugin `yaml:"plugins"`
	}{Plugins: []ContainerPlugin{}}

	for _, p := range registry.Plugins {
		if cp, ok := p.(*ContainerPlugin); ok {
			wrapper.Plugins = append(wrapper.Plugins, *cp)
		}
	}

	data, err := yaml.Marshal(wrapper)
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := ioutil.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

// RegisterAdminRoutes installs the registration endpoint on the provided router.
// It uses the given registry to append container plugins on successful registration.
func RegisterAdminRoutes(registry *PluginRegistry, router common.Router) {
	router.AddPostRoute("/plugins/register", func(c *gin.Context) {
		var req RegisterRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
			return
		}

		if req.Name == "" || req.BaseURL == "" || req.Route == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name, base_url and route are required"})
			return
		}

		cp := &ContainerPlugin{
			Name:    req.Name,
			BaseURL: req.BaseURL,
			Route:   req.Route,
		}

		registry.Plugins = append(registry.Plugins, cp)
		cp.RegisterRoutes(router)

		// persist to disk (best-effort)
		manifestPath := "plugins/containers.yml"
		if err := persistContainerPlugins(manifestPath, registry); err != nil {
			// log to stdout (visible in container logs)
			fmt.Println("warning: failed to persist containers manifest:", err)
		}

		c.JSON(http.StatusCreated, gin.H{"id": req.ID, "name": req.Name})
	})
}

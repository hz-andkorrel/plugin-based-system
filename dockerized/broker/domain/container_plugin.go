package domain

import (
	"io"
	"net/http"
	"strings"
	"time"

	"plugins/common"

	"github.com/gin-gonic/gin"
)

// ContainerPlugin is a minimal plugin representation for a single-route container.
type ContainerPlugin struct {
	Name    string `yaml:"name"`
	BaseURL string `yaml:"base_url"`
	Route   string `yaml:"route"`
}

// Ensure ContainerPlugin implements common.Plugin at compile time.
var _ common.Plugin = (*ContainerPlugin)(nil)

// RegisterRoutes registers a single GET proxy route on the provided router.
// This minimal implementation assumes the plugin exposes the same path at its base URL.
func (p *ContainerPlugin) RegisterRoutes(router common.Router) {
	base := strings.TrimRight(p.BaseURL, "/")
	path := p.Route

	router.AddGetRoute(path, func(c *gin.Context) {
		target := base + path
		if q := c.Request.URL.RawQuery; q != "" {
			target = target + "?" + q
		}

		client := &http.Client{Timeout: 5 * time.Second}
		req, err := http.NewRequest(http.MethodGet, target, nil)
		if err != nil {
			c.Status(http.StatusBadGateway)
			return
		}
		req.Header.Set("User-Agent", "plugin-proxy/1.0")

		resp, err := client.Do(req)
		if err != nil {
			c.Status(http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// copy headers
		for k, v := range resp.Header {
			for _, vv := range v {
				c.Writer.Header().Add(k, vv)
			}
		}
		c.Status(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)
	})
}

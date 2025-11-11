package domain

import (
	"io/ioutil"
	"os"

	"plugins/common"

	"github.com/goccy/go-yaml"
)

type manifestWrapper struct {
	Plugins []ContainerPlugin `yaml:"plugins"`
}

// LoadContainerPlugins parses a simplified YAML manifest file and returns
// a slice of ContainerPlugin instances implementing common.Plugin.
func LoadContainerPlugins(path string) ([]common.Plugin, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var wrapper manifestWrapper
	if err := yaml.Unmarshal(bytes, &wrapper); err != nil {
		return nil, err
	}

	var plugins []common.Plugin
	for _, p := range wrapper.Plugins {
		cp := p
		plugins = append(plugins, &cp)
	}

	return plugins, nil
}

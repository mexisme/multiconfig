package multiconfig

import (
	"github.com/mexisme/multiconfig/marshal"
)

// Configs contains map of "path" --> BodyEnv structs
type Configs map[string]*marshal.Config

// MultiConfig is simply the struct method-wrapper for the "multiconfig" package
type MultiConfig struct {
	configs Configs
}

// New creates a new MultiConfigs struct
func New() *MultiConfig {
	return &MultiConfig{configs: make(Configs)}
}

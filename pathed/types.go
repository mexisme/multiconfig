package pathed

import (
	"path/filepath"
	"strings"

	"github.com/mexisme/multiconfig/common"
)

// ConfigFormats defines the type for the *Format enum.
type ConfigFormats int

const (
	// UnknownFormat is when the config file format isn't recognised.
	UnknownFormat ConfigFormats = iota
	// EnvFormat is when the config file format is ".env".
	EnvFormat
	// TOMLFormat is when the config file format is TOML.
	TOMLFormat
	// YAMLFormat is when the config file format is YAML.
	YAMLFormat
	// JSONFormat is when the config file format is JSON.
	JSONFormat
)

/*
Config contains the given "path", "extn" and "body" of a
config, as well as the parsed config --> KV map.
*/
type Config struct {
	path, extn, body string
	parsed           common.BodyMap
}

// New object
func New() *Config {
	return &Config{}
}

// SetPath adds Path string to the Config struct.
func (s *Config) SetPath(path string) *Config {
	s.path = path
	s.extn = strings.ToLower(filepath.Ext(s.path))
	// Explicitly reset it:
	s.parsed = nil

	return s
}

// SetBody adds Body string to the Config struct.
func (s *Config) SetBody(body string) *Config {
	s.body = body
	// Explicitly reset it:
	s.parsed = nil

	return s
}

// Key returns s.path as a key, for use by the configs package.
func (s *Config) Key() string {
	return s.path
}

// Body returns s.path as a key, for use by the configs package.
func (s *Config) Body() string {
	return s.body
}

// ToBodyMap returns the parsed map, for use by the configs package.
func (s *Config) ToBodyMap() (common.BodyMap, error) {
	if s.parsed == nil {
		if err := s.Unmarshal(); err != nil {
			return nil, err
		}
	}
	return s.parsed, nil
}

// ConfigFormat returns an enum representing the format of s.body.
func (s *Config) ConfigFormat() (ConfigFormats, error) {
	switch s.extn {
	case ".env":
		return EnvFormat, nil

	case ".toml":
		return TOMLFormat, nil

	case ".yaml", ".yml":
		return YAMLFormat, nil

	case ".json":
		return JSONFormat, nil

		// case ".properties", ".props", ".prop":
	}

	return UnknownFormat, UnsupportedFormatError(s.extn)
}

func (s *Config) checkAttributes() error {
	if s.path == "" {
		return EmptyAttributeError("path")
	}
	if s.body == "" {
		return EmptyAttributeError("body")
	}

	return nil
}

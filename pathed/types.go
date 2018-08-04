/*
Package pathed parses a config file into a common.BodyMap struct.

Rather than simply provide a path to a file, which would make it harder to support arbitrary storage systems (like AWS S3) it expects "path" and "body" strings to be provided.

The path is an arbitrary string, used by the wrapping multiconfig package as the Key() for merge-order and by the pathed package to determine config-format.

The file-extension of the path defines the format:

- .env -- Dot-env file
- .toml -- TOML file
- .yaml, .yml -- YAML file
- .json -- JSON file

The body is the contents of the above file, as string type.

NOTE: The body is not parsed until the ToBodyMap() method is run.
*/
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

// SetPath adds a Path string to the Config struct.
func (s *Config) SetPath(path string) *Config {
	s.path = path
	s.extn = strings.ToLower(filepath.Ext(s.path))
	// Explicitly reset it:
	s.parsed = nil

	return s
}

// SetBody adds a Body string to the Config struct.
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

// Body returns s.body returns the contents of the body, as a string.
func (s *Config) Body() string {
	return s.body
}

/*
ToBodyMap parses the body, for use by the multiconfig package.
On first execution, it will parse the body and store it. On subsequent, it will return the previously stored parse.

If there's an error when parsing, it will be returned instead of the parsed body.

NOTE: SetPath() and SetBody() above will reset the stored parse.
*/
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

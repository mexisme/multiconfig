package env

import (
	"os"
)

// BodyMap is type of the post-unmarshaled / parsed copy of the KV pairs.
type BodyMap map[string]string

// Envs is the type of an os.Environ-provided env-list.
type Envs []string

/*
Config contains the given "path", "extn" and "body" of a
config, as well as the parsed config --> KV map.
*/
type Config struct {
	env    Envs
	parsed BodyMap
}

// New object.
func New() *Config {
	return &Config{}
}

// SetEnv sets s.env struct, with the given env.
func (s *Config) SetEnv(env Envs) *Config {
	s.env = env
	// We don't want to parse it here, partly so we don't have to deal with
	// returning an error, but also because it feels more-appropriate to
	// produce errors when asking to get the value:
	s.parsed = nil

	return s
}

// SetBodyMap creates a new env.Configs object, with the contents of a Maps object added to its EnvMaps.
func (s *Config) SetBodyMap(env BodyMap) *Config {
	s.parsed = env
	// We don't want to parse it here, partly so we don't have to deal with
	// returning an error, but also because it feels more-appropriate to
	// produce errors when asking to get the value:
	s.env = nil

	return s
}

// Key returns a sort key, for use by the configs package.
// This makes little sense for os.Environ, so we always return "".
func (s *Config) Key() string {
	return ""
}

// Body returns s.path as a key, for use by the configs package.
func (s *Config) Body() Envs {
	if e, err := s.ToOsEnviron(); err == nil {
		return e
	}
	return nil
}

// FromOsEnviron sets s.env to the current os.Environ.
func (s *Config) FromOsEnviron() *Config {
	return s.SetEnv(os.Environ())
}

// ToOsEnviron returns the config, converted to a list of "K=V" lines
func (s *Config) ToOsEnviron() (Envs, error) {
	var err error
	if s.env == nil {
		s.env, err = mapToEnv(s.parsed)
	}
	return s.env, err
}

// ToBodyMap returns the parsed map, for use by the configs package.
func (s *Config) ToBodyMap() (BodyMap, error) {
	var err error
	if s.parsed == nil {
		s.parsed, err = envToMap(s.env)
	}
	return s.parsed, err
}

package marshal

// ConfigMap is a convenience type for the internal EnvVar map
type ConfigMap map[string]string

// Config contains the given "path" and "body" of a config,
// as well as the parsed config --> KV map
type Config struct {
	path   string
	body   string
	parsed ConfigMap
}

// New object
func New() *Config {
	return &Config{}
}

// AddPathBody creates a new Config with Path and Body added, and the Env updated
func (s *Config) AddPathBody(path, body string) *Config {
	s.path = path
	s.body = body
	// Explicitly reset it:
	s.parsed = nil

	return s
}

// Map returns the parsed map
func (s *Config) Map() (ConfigMap, error) {
	if s.parsed == nil {
		if err := s.checkAttributes(); err != nil {
			return nil, err
		}

		if err := s.unmarshallIntoParsed(); err != nil {
			return nil, err
		}
	}
	return s.parsed, nil
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

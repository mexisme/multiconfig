package env

// Map is the type used to store $NAME="Val" KV pairs
type Map map[string]string

// Maps is a list of Map objects
type Maps []Map

// Details is simply the struct method-wrapper for the "env" package
type Details struct {
	EnvMaps Maps
}

// MultiConfig interfaces to a type supporting the Combine() method -- usually the "multiconfig" package
type MultiConfig interface {
	Combine() (map[string]string, error)
}

// New creates a new env.Details struct
func New() *Details {
	return &Details{}
}

// AddEnvMaps creates a new env.Details object, with the contents of a Maps object added to its EnvMaps
func (s *Details) AddEnvMaps(envs *Maps) *Details {
	clone := *s // This does a shallow clone

	if clone.EnvMaps == nil {
		clone.EnvMaps = make(Maps, 0)
	}
	clone.EnvMaps = append(clone.EnvMaps, *envs...)

	return &clone
}

// AddEnvMap creates a new env.Details object, with each Map object arg to its EnvMaps
func (s *Details) AddEnvMap(env ...Map) *Details {
	newEnvs := append(make(Maps, 0), env...)
	return s.AddEnvMaps(&newEnvs)
}

// AddMultiConfig creates a new env.Details struct, with the dotEnv KV map object added to its EnvMaps
func (s *Details) AddMultiConfig(configs MultiConfig) *Details {
	combinedConfigs, _ := configs.Combine()
	return s.AddEnvMap(combinedConfigs)
}

// AddOsEnviron creates a new env.Details struct, with the given env (usually an os.Environ array) added to its EnvMaps
func (s *Details) AddOsEnviron(env []string) *Details {
	return s.AddEnvMap(s.envToMap(env))
}

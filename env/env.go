package env

import (
	log "github.com/sirupsen/logrus"
)

// Combine merges the EnvMaps objects into a single env.Map, and returns that
func (s *Details) Combine() Map {
	newEnvMap := make(Map)
	for _, envMap := range s.EnvMaps {
		log.WithFields(log.Fields{"ToEnv": newEnvMap, "fromEnv": envMap}).Debug("Adding env...")
		s.mergeEnv(envMap, newEnvMap)
	}

	log.WithFields(log.Fields{"new": newEnvMap}).Debug("Combined envs")

	return newEnvMap
}

// ToOsEnviron combines the EnvMaps objects, and returns it in the os.Environ format, suitable for passing to the "exec" package
func (s *Details) ToOsEnviron() []string {
	return s.mapToEnv(s.Combine())
}

/*
Merge the first map *into* the second

Because maps are implicit references, updating the second map *updates* the caller's map in-place
*/
func (s *Details) mergeEnv(from, to map[string]string) {
	for k, v := range from {
		to[k] = v
	}
}

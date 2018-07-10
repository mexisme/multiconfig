package env

import (
	"fmt"
	"strings"
)

func (s *Details) envToMap(env []string) map[string]string {
	newEnv := make(map[string]string)

	for _, envLine := range env {
		kv := strings.SplitN(envLine, "=", 2)
		newEnv[kv[0]] = kv[1]
	}

	return newEnv
}

func (s *Details) mapToEnv(env map[string]string) []string {
	newEnv := make([]string, 0)

	for name, val := range env {
		line := fmt.Sprintf("%s=%s", name, val)
		newEnv = append(newEnv, line)
	}

	return newEnv
}

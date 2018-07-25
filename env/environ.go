package env

import (
	"fmt"
	"strings"
)

func envToMap(env Envs) (BodyMap, error) {
	if len(env) < 1 {
		return nil, EmptyAttributeError("env")
	}

	newEnv := make(BodyMap)

	for _, envLine := range env {
		kv := strings.SplitN(envLine, "=", 2)

		if len(kv) != 2 {
			return nil, ParseSplitError(envLine)
		}
		newEnv[kv[0]] = kv[1]
	}

	return newEnv, nil
}

func mapToEnv(bodyMap BodyMap) (Envs, error) {
	if bodyMap == nil {
		return nil, EmptyAttributeError("bodyMap")
	}
	newEnv := make(Envs, 0)

	for name, val := range bodyMap {
		line := fmt.Sprintf("%s=%s", name, val)
		newEnv = append(newEnv, line)
	}

	return newEnv, nil
}

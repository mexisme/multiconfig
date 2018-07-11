package marshal

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	// "github.com/hashicorp/hcl"
	// "github.com/magiconair/properties"
	toml "github.com/pelletier/go-toml"
	// "github.com/spf13/cast"
	"gopkg.in/yaml.v2"
)

func (s *Config) mapToConfigMap(tree interface{}) ConfigMap {
	parsed := make(ConfigMap)

	switch tree.(type) {
	case gotenv.Env:
		for k, v := range tree.(gotenv.Env) {
			parsed[k] = v
		}
	case map[string]interface{}:
		for k, v := range tree.(map[string]interface{}) {
			parsed[k] = fmt.Sprintf("%v", v)
		}
	}

	return parsed
}

func (s *Config) unmarshallIntoParsed() error {
	extn := s.extn()

	log.WithFields(log.Fields{"Path": s.path, "Extension": extn}).Debug("Reading file...")

	switch extn {
	case ".env":
		env, err := gotenv.StrictParse(strings.NewReader(s.body))
		if err != nil {
			return s.ParseError(err)
		}
		s.parsed = s.mapToConfigMap(env)

	case ".toml":
		tree, err := toml.Load(s.body)
		if err != nil {
			return s.ParseError(err)
		}

		s.parsed = s.mapToConfigMap(tree.ToMap())

	case ".yaml", ".yml":
		if err := yaml.Unmarshal([]byte(s.body), &s.parsed); err != nil {
			return s.ParseError(err)
		}

	case ".json":
		var tree map[string]interface{}
		if err := json.Unmarshal([]byte(s.body), &tree); err != nil {
			return s.ParseError(err)
		}

		s.parsed = s.mapToConfigMap(tree)

	// case ".properties", ".props", ".prop":
	default:
		log.WithField("extension", extn).Debug("Unsupported format")
		return UnsupportedFormatError(extn)
	}

	return nil
}

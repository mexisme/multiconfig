package pathed

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mexisme/multiconfig/common"
	"github.com/subosito/gotenv"
	// "github.com/hashicorp/hcl"
	// "github.com/magiconair/properties"
	toml "github.com/pelletier/go-toml"
	// "github.com/spf13/cast"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes/parses the given s.body string into the s.parsed field.
// TODO: Use Viper, instead:
func (s *Config) Unmarshal() error {
	if err := s.checkAttributes(); err != nil {
		return err
	}

	configFormat, err := s.ConfigFormat()
	if err != nil {
		return err
	}

	switch configFormat {
	case EnvFormat:
		env, err := gotenv.StrictParse(strings.NewReader(s.body))
		if err != nil {
			return s.ParseError(err)
		}
		s.parsed = s.mapToConfigMap(env)

	case TOMLFormat:
		tree, err := toml.Load(s.body)
		if err != nil {
			return s.ParseError(err)
		}
		s.parsed = s.mapToConfigMap(tree.ToMap())

	case YAMLFormat:
		if err := yaml.Unmarshal([]byte(s.body), &s.parsed); err != nil {
			return s.ParseError(err)
		}

	case JSONFormat:
		var tree map[string]interface{}
		if err := json.Unmarshal([]byte(s.body), &tree); err != nil {
			return s.ParseError(err)
		}
		s.parsed = s.mapToConfigMap(tree)

	default:
		return s.ParseError(fmt.Errorf("Unexpected parse failure for %#v: format %#v", s.Body(), configFormat))
	}

	return nil
}

func (s *Config) mapToConfigMap(tree interface{}) common.BodyMap {
	parsed := make(common.BodyMap)

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

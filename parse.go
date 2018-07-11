package multiconfig

import (
	"sort"

	"github.com/mexisme/multiconfig/marshal"
)

/*
AddFromString allows to add new *.env string to the list of MultiConfigs

- Need to provide the body-content as a string.

- Need to provide a "path" to where the .env was read from, to allow for combination based on lexical-order of this path.

- This gets parsed immediately, not deferred until combined.
*/
// TODO: Is "path" the wrong name for this?
func (s *MultiConfig) AddFromString(path string, body string) error {

	// log.WithFields(log.Fields{"path": path, "body": body}).Debug("Parsing multiconfig file...")
	config := marshal.New().AddPathBody(path, body)
	if _, err := config.Map(); err != nil {
		return err
	}
	s.configs[path] = config
	return nil
}

/*
Combine Combine all the .env's provided into a single .env map

Returns a new .env map (gotenv.Env) with the combined .env's.
The .env's are combined by lexical-ordering of the "path" field;

i.e. an .env with a "path" of "z..." will be added after those with a "path" of "a..."
*/
// TODO: Need to make sure file's named ".env" (only) are parsed last, somehow...
func (s *MultiConfig) Combine() (map[string]string, error) {
	joined := make(map[string]string)

	for _, path := range s.sortedPaths() {
		config, err := s.configs[path].Map()
		if err != nil {
			return nil, err
		}
		s.merge(config, joined)
	}

	return joined, nil
}

func (s *MultiConfig) merge(from, to map[string]string) {
	for k, v := range from {
		to[k] = v
	}
}

func (s *MultiConfig) sortedPaths() []string {
	names := make([]string, 0)

	for name := range s.configs {
		names = append(names, name)
	}
	sort.Strings(names)

	return names
}

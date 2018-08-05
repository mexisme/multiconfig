/*
Package multiconfig can take config from multiple config files and environment variables.

It can merge the list of configs together into a single KV-style string (i.e. "map[string]string").

It can then convert that map into a array of strings in the same format as "os.Environ()" provides, and compatible with the "Env" field of the [os/exec.Cmd](https://golang.org/pkg/os/exec/#Cmd) type and the "envv" arg [syscall.Exec](https://golang.org/pkg/syscall/#Exec) function.

Each config object is expected to have an associated "Key" that's used to ascii-sort the list of configs into an "merge order".
The "merge order" determines the order that the list of configs will be merged together by the Merge() method.
*/
package multiconfig

import "github.com/mexisme/multiconfig/common"

/*
ItemInterface is the interface for a config map, providing the methods needed to sort a set of configs,
as well as for extracting the BodyMap (for merging, later)
*/
type ItemInterface interface {
	Key() string
	ToBodyMap() (common.BodyMap, error)
}

// Map is the type of config list.
// This is a list of sort-keys associated to a config item (i.e. ItemInterface)
type Map struct {
	items []ItemInterface
}

// New struct
func New() *Map {
	return &Map{}
}

// AddItem is for adding another item to the config list.
func (s *Map) AddItem(item ...ItemInterface) *Map {
	s.items = append(s.items, item...)
	return s
}

// Items returns a copy of the current config items.
func (s *Map) Items() []ItemInterface {
	return s.items[:]
}

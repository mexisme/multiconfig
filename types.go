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

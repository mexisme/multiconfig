package multiconfig

import "sort"

// Len is part of sort.Interface.
func (s *Map) Len() int {
	return len(s.items)
}

// Swap is part of sort.Interface.
func (s *Map) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

// Less is part of sort.Interface.
func (s *Map) Less(i, j int) bool {
	return s.items[i].Key() < s.items[j].Key()
}

// Sorted returns sorted _copy_ of the Map.
func (s *Map) Sorted() *Map {
	clone := s.clone()
	sort.Sort(clone)

	return clone
}

func (s *Map) clone() *Map {
	clone := *s
	return &clone
}

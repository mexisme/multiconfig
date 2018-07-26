package multiconfig

import "github.com/mexisme/multiconfig/common"

/*
Merge all the BodyMap's provided into a single BodyMap, return a new BodyMap.

They are combined by first lexically-ordering all the configs with the Sorted()
method, then merging thme together in lexical-order -- whereby the later keys
in a BodyMap will override earlier ones.

i.e. a BodyMap with a Key() of "z..." will be merged-in after those with a Key()
of "a..."
*/
func (s *Map) Merge() (common.BodyMap, error) {
	sorted := s.Sorted()
	joined := make(common.BodyMap)

	for _, item := range (*sorted).items {
		body, err := item.ToBodyMap()
		if err != nil {
			// Wrap error:
			return nil, MergeError(err, item)
		}
		s.mergeBodies(body, joined)
	}

	return joined, nil
}

func (s *Map) mergeBodies(from, to common.BodyMap) {
	for k, v := range from {
		to[k] = v
	}
}

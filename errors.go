package multiconfig

import (
	"fmt"

	"github.com/pkg/errors"
)

// type parseError struct {
// 	error
// 	path string
// }

type mergeError struct {
	error
	item ItemInterface
}

// MergeError is for when a required attribute is empty
func MergeError(err error, item ItemInterface) error {
	return errors.WithStack(&mergeError{err, item})
}

// Error returns the formatted configuration error.
func (s *mergeError) Error() string {
	return fmt.Sprintf("Failed to merge %#v", s.item.Key())
}

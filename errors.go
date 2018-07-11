package multiconfig

import (
	"fmt"

	"github.com/pkg/errors"
)

type parseError struct {
	error
	path string
}

type emptyAttributeError struct {
	attr string
}

// EmptyAttributeError is for when a required attribute is empty
func EmptyAttributeError(attr string) error {
	return errors.WithStack(&emptyAttributeError{attr})
}

// Error returns the formatted configuration error.
func (s *emptyAttributeError) Error() string {
	return fmt.Sprintf("%#v is empty", s.attr)
}

package env

import (
	"fmt"

	"github.com/pkg/errors"
)

type emptyAttributeError struct {
	attr string
}

// EmptyAttributeError is for when a required attribute is empty.
func EmptyAttributeError(attr string) error {
	return errors.WithStack(&emptyAttributeError{attr})
}

func (s *emptyAttributeError) Error() string {
	return fmt.Sprintf("%#v was empty", s.attr)
}

type parseSplitError struct {
	env string
}

// ParseSplitError is for when failing to parse a env-line.
func ParseSplitError(env string) error {
	return errors.WithStack(&parseSplitError{env})
}

// Error returns the formatted configuration error.
func (s *parseSplitError) Error() string {
	return fmt.Sprintf("%#v couldn't be parsed by 'K=V'", s.env)
}

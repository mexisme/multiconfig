package marshal

import (
	"fmt"

	"github.com/pkg/errors"
)

type parseError struct {
	error
	path string
}

// ParseError is for when failing to parse an config file
func (s *Config) ParseError(err error) error {
	return &parseError{errors.WithStack(err), s.path}
}

func (s *parseError) Cause() error { return s.error }

// Error returns the formatted configuration error.
func (s *parseError) Error() string {
	return fmt.Sprintf("While parsing config %#v: %s", s.path, s.error)
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

type unsupportedFormatError struct {
	format string
}

// UnsupportedFormatError is for when asked to parse a format-type that's not supported
func UnsupportedFormatError(format string) error {
	return errors.WithStack(&unsupportedFormatError{format})
}

// Error returns the formatted configuration error.
func (s *unsupportedFormatError) Error() string {
	return fmt.Sprintf("%#v format is not yet supported", s.format)
}

package pathed

import (
	"fmt"

	"github.com/pkg/errors"
)

type parseError struct {
	error
	path string
}

// ParseError is for when failing to parse a config.
func (s *Config) ParseError(err error) error {
	return &parseError{errors.WithStack(err), s.path}
}

func (s *parseError) Cause() error { return s.error }

// Error returns the formatted configuration error.
func (s *parseError) Error() string {
	return fmt.Sprintf("While parsing config %#v: %s", s.path, s.error)
}

type attributeError struct {
	message string
	attr    string
}

// EmptyAttributeError is for when a required attribute is empty.
func EmptyAttributeError(attr string) error {
	// NOTE: We repeat the following line, to avoid the stack-trace having
	//       the extra function-call to AttributeError() in it
	return errors.WithStack(&attributeError{"%#v is empty", attr})
}

// AttributeError is for when there is an error in using a required attribute.
func AttributeError(message, attr string) error {
	return errors.WithStack(&attributeError{message, attr})
}

// Error returns the formatted configuration error.
func (s *attributeError) Error() string {
	return fmt.Sprintf(s.message, s.attr)
}

type unsupportedFormatError struct {
	format string
}

// UnsupportedFormatError is for when asked to parse a format-type that's not supported.
func UnsupportedFormatError(format string) error {
	return errors.WithStack(&unsupportedFormatError{format})
}

// Error returns the formatted configuration error.
func (s *unsupportedFormatError) Error() string {
	return fmt.Sprintf("%#v format is not yet supported", s.format)
}

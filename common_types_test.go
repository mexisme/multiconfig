package multiconfig_test

import (
	. "github.com/mexisme/multiconfig"
)

type ItemArr struct {
	Body map[string]string
}

func (s *ItemArr) Key() string {
	return ""
}

func (s *ItemArr) ToBodyMap() (BodyMap, error) {
	return s.Body, nil
}

type ItemMap struct {
	K    string
	Body map[string]string
}

func (s *ItemMap) Key() string {
	return s.K
}

func (s *ItemMap) ToBodyMap() (BodyMap, error) {
	return s.Body, nil
}

package multiconfig_test

import "github.com/mexisme/multiconfig/common"

// . "github.com/mexisme/multiconfig"

type ItemArr struct {
	Body map[string]string
}

func (s *ItemArr) Key() string {
	return ""
}

func (s *ItemArr) ToBodyMap() (common.BodyMap, error) {
	return s.Body, nil
}

type ItemMap struct {
	K    string
	Body map[string]string
}

func (s *ItemMap) Key() string {
	return s.K
}

func (s *ItemMap) ToBodyMap() (common.BodyMap, error) {
	return s.Body, nil
}

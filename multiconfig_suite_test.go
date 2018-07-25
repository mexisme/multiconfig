package multiconfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConfigmap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "multiconfig Suite")
}

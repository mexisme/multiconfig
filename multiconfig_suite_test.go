package multiconfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMultiConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultiConfig Suite")
}

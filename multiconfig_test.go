package multiconfig_test

import (
	. "github.com/mexisme/multiconfig"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The MultiConfig type", func() {
	It("initialises an empty map", func() {
		env := New()

		Expect(env.Combine()).To(BeEmpty())
	})
})

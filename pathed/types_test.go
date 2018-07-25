package pathed_test

import (
	. "github.com/mexisme/multiconfig/pathed"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("multiconfig/pathed/types", func() {
	Context("when creating new object", func() {
		It("should create an object with blank Key() and Body()", func() {
			c := New()
			Expect(c.Key()).To(Equal(""))
			Expect(c.Body()).To(Equal(""))
		})

		It("should correctly add a new path", func() {
			c := New().SetPath("a/b/c")
			Expect(c.Key()).To(Equal("a/b/c"))
		})

		It("should correctly add a new body", func() {
			c := New().SetBody("booya")
			Expect(c.Body()).To(Equal("booya"))
		})
	})

	Context("when calculating config format", func() {
		DescribeTable("",
			func(path string, expected ConfigFormats) {
				c := New().SetPath(path)
				Expect(c.ConfigFormat()).To(Equal(expected))
			},
			Entry("should calculate a .env format", "a/b/c.env", EnvFormat),
			Entry("should calculate a .toml format", "a/b/c.toml", TOMLFormat),
			Entry("should calculate a .yaml format", "a/b/c.yaml", YAMLFormat),
			Entry("should calculate a .yml format", "a/b/c.yml", YAMLFormat),
			Entry("should calculate a .json format", "a/b/c.json", JSONFormat),
		)

		It("should calculate an unknown format", func() {
			c := New().SetPath("a/b/c.meh")
			_, err := c.ConfigFormat()
			Expect(err).To(MatchError(Equal(`".meh" format is not yet supported`)))
		})
	})
})

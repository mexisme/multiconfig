package env_test

import (
	. "github.com/mexisme/multiconfig/env"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("multiconfig/env/types", func() {
	Context("when creating new struct", func() {
		It("should create an object with blank Key() and Body()", func() {
			c := New()
			Expect(c.Key()).To(Equal(""))
			Expect(c.Body()).To(Equal(Envs(nil)))
		})
	})

	Context("when updating the struct", func() {
		It("should never add a new Key()", func() {
			c := New().SetEnv(Envs{
				"A=  A1A1  ",
				"B=12",
				"C=over there",
				"D= 15.1 ",
			})
			Expect(c.Key()).To(Equal(""))
		})

		It("should correctly add a new env body", func() {
			c := New().SetEnv(Envs{
				"A=  A1A1  ",
				"B=12",
				"C=over there",
				"D= 15.1 ",
			})
			Expect(c.Body()).To(Equal(Envs{
				"A=  A1A1  ",
				"B=12",
				"C=over there",
				"D= 15.1 ",
			}))
		})
	})
})

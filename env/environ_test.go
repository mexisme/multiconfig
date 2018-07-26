package env_test

import (
	"os"

	"github.com/mexisme/multiconfig/common"
	. "github.com/mexisme/multiconfig/env"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("multiconfig/env/environ", func() {
	Describe("when getting an os.Environ", func() {
		BeforeEach(func() {
			os.Clearenv()
		})

		It("should return an empty environ, if it's cleared", func() {
			c := New().FromOsEnviron()
			Expect(c.ToOsEnviron()).To(Equal(Envs{}))
		})

		It("should return the environ", func() {
			os.Setenv("A", "  A1A1  ")
			os.Setenv("B", "12")

			c := New().FromOsEnviron()
			e, err := c.ToOsEnviron()
			Expect(err).To(Succeed())
			Expect(e).To(Equal(Envs{"A=  A1A1  ", "B=12"}))
		})

		It("should return the environ", func() {
			os.Setenv("A", "  A1A1  ")
			os.Setenv("B", "12")

			c := New().FromOsEnviron()
			e, err := c.ToOsEnviron()
			Expect(err).To(Succeed())
			Expect(e).To(Equal(Envs{"A=  A1A1  ", "B=12"}))
		})
	})

	Describe("when getting a BodyMap", func() {
		BeforeEach(func() {
			os.Clearenv()
		})

		It("should return an empty BodyMap, if it's cleared", func() {
			c := New().FromOsEnviron()
			e, err := c.ToBodyMap()
			Expect(err).To(HaveOccurred())
			Expect(e).To(Equal(common.BodyMap(nil)))
		})

		It("should return the environ", func() {
			os.Setenv("A", "  A1A1  ")
			os.Setenv("B", "12")

			c := New().FromOsEnviron()
			e, err := c.ToBodyMap()
			Expect(err).To(Succeed())
			Expect(e).To(Equal(common.BodyMap{
				"A": "  A1A1  ",
				"B": "12",
			}))
		})

		It("should return the environ", func() {
			os.Setenv("A", "  A1A1  ")
			os.Setenv("B", "12")

			c := New().FromOsEnviron()
			e, err := c.ToBodyMap()
			Expect(err).To(Succeed())
			Expect(e).To(Equal(common.BodyMap{
				"A": "  A1A1  ",
				"B": "12",
			}))
		})

	})

	Describe("when setting a BodyMap", func() {
		It("should return the same BodyMap", func() {
			c := New().SetBodyMap(common.BodyMap{
				"A": "  A1A1  ",
				"B": "12",
			})
			e, err := c.ToBodyMap()
			Expect(err).To(Succeed())
			Expect(e).To(Equal(common.BodyMap{
				"A": "  A1A1  ",
				"B": "12",
			}))
		})

		It("should return it as an environ", func() {
			c := New().SetBodyMap(common.BodyMap{
				"A": "  A1A1  ",
				"B": "12",
			})
			e, err := c.ToOsEnviron()
			Expect(err).To(Succeed())
			Expect(e).To(ConsistOf(Envs{"A=  A1A1  ", "B=12"}))
		})
	})
})

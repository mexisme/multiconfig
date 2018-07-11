package multiconfig_test

import (
	. "github.com/mexisme/multiconfig"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The MultiConfig type, parsing functions,", func() {
	Describe("intialises a simple map", func() {
		It("from .env", func() {
			env := New()
			env.AddFromString("lol.env", `A=1
B=2
C=3`)

			Expect(env.Combine()).To(Equal(map[string]string{
				"A": "1",
				"B": "2",
				"C": "3",
			}))
		})

		It("from .toml", func() {
			env := New()
			env.AddFromString("lol.toml", `D=4
E=5
F=6`)

			Expect(env.Combine()).To(Equal(map[string]string{
				"D": "4",
				"E": "5",
				"F": "6",
			}))
		})

		It("from .yaml", func() {
			env := New()
			env.AddFromString("lol.yaml", `---
G: 7
H: 8
I: 9`)

			Expect(env.Combine()).To(Equal(map[string]string{
				"G": "7",
				"H": "8",
				"I": "9",
			}))
		})

		It("from .json", func() {
			env := New()
			env.AddFromString("lol.json", `{
  "J": 10,
  "K": 11,
  "L": 12
}`)

			Expect(env.Combine()).To(Equal(map[string]string{
				"J": "10",
				"K": "11",
				"L": "12",
			}))
		})

		It("from .properties", func() {
			env := New()
			Expect(env.AddFromString("lol.properties", `M: 13
N: 14
O: 15`)).To(HaveOccurred())

			// TODO: Restore when supported:
			// Expect(env.Combine()).To(Equal(map[string]string{
			// 	"M": "14",
			// 	"N": "15",
			// 	"O": "16",
			// }))
		})
	})

	It("combines multiple envs", func() {
		env := New()
		env.AddFromString("lol2.env", `A=1
B=2
C=3`)
		env.AddFromString("lol.env", `D=4
E=5`)

		Expect(env.Combine()).To(Equal(map[string]string{
			"A": "1",
			"B": "2",
			"C": "3",
			"D": "4",
			"E": "5",
		}))
	})

	It("combines multiple envs based on lexical-order of the path", func() {
		env := New()
		env.AddFromString("lol2.env", `A=11
B=2
C=3`)
		env.AddFromString("lol.env", `A=1
D=4`)

		Expect(env.Combine()).To(Equal(map[string]string{
			"A": "11",
			"B": "2",
			"C": "3",
			"D": "4",
		}))
	})
})

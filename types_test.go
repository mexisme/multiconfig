package multiconfig_test

import (
	. "github.com/mexisme/multiconfig"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("multiconfig/types", func() {
	Describe("with multiple items", func() {
		m := New()
		m.AddItem(&ItemArr{
			Body: map[string]string{
				"a": "b",
				"c": "d",
			},
		})
		m.AddItem(&ItemMap{
			K: "a/b/c",
			Body: map[string]string{
				"e": "f",
				"g": "h",
			},
		})
		m.AddItem(&ItemMap{
			K: "d/e/f",
			Body: map[string]string{
				"h": "i",
				"j": "k",
			},
		})

		It("has the correct length", func() {
			Expect(m.Len()).To(Equal(3))
		})

		It("can use multiple types", func() {
			Expect(m.Items()).To(ConsistOf(
				[]ItemInterface{
					&ItemMap{
						K: "a/b/c",
						Body: map[string]string{
							"e": "f",
							"g": "h",
						},
					},
					&ItemMap{
						K: "d/e/f",
						Body: map[string]string{
							"h": "i",
							"j": "k",
						},
					},
					&ItemArr{
						Body: map[string]string{
							"a": "b",
							"c": "d",
						},
					},
				},
			))
		})
	})
})

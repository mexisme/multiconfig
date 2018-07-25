package multiconfig_test

import (
	. "github.com/mexisme/multiconfig"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("multiconfig/sort", func() {
	It("sorts by Key()", func() {
		m := Map{}
		m.AddItem(&ItemMap{
			K: "d/e/f",
			Body: map[string]string{
				"h": "i",
				"j": "k",
			},
		})
		m.AddItem(&ItemMap{
			K: "a/b/c",
			Body: map[string]string{
				"e": "f",
				"g": "h",
			},
		})
		m.AddItem(&ItemArr{
			Body: map[string]string{
				"a": "b",
				"c": "d",
			},
		})

		n := m.Sorted()

		Expect((*n).Items()[0]).To(Equal(
			&ItemArr{
				Body: map[string]string{
					"a": "b",
					"c": "d",
				},
			}))
		Expect((*n).Items()[1]).To(Equal(
			&ItemMap{
				K: "a/b/c",
				Body: map[string]string{
					"e": "f",
					"g": "h",
				},
			}))
		Expect((*n).Items()[2]).To(Equal(
			&ItemMap{
				K: "d/e/f",
				Body: map[string]string{
					"h": "i",
					"j": "k",
				},
			}))
	})
})

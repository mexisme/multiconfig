package multiconfig_test

import (
	. "github.com/mexisme/multiconfig"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The configs package", func() {
	Describe("when Combining the Bodies", func() {
		It("combines them correctly", func() {
			e := &Map{}

			e.AddItem(&ItemMap{
				K: "a",
				Body: BodyMap{
					"A": "1",
					"B": "2",
				},
			},
				&ItemMap{
					K: "b",
					Body: BodyMap{
						"C": "3",
						"D": "4",
					},
				},
				&ItemMap{
					K: "c",
					Body: BodyMap{
						"E": "5",
						"F": "6",
					},
				})

			Expect(e.Merge()).To(Equal(BodyMap{
				"A": "1",
				"B": "2",
				"C": "3",
				"D": "4",
				"E": "5",
				"F": "6",
			}))
		})
	})
})

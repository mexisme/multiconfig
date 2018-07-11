package env_test

import (
	. "github.com/mexisme/multiconfig/env"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The env package", func() {
	Describe("when exporting to os.Environ", func() {
		It("exports correctly", func() {
			envMap := Map{
				"A": "1",
				"B": "2",
			}

			e := New().AddEnvMap(envMap)
			Expect(e.ToOsEnviron()).To(ConsistOf([]string{
				"A=1",
				"B=2",
			}))
		})
	})

	Describe("when Combining the EnvMaps", func() {
		It("combines them correctly", func() {
			envMap := Maps{
				{
					"A": "1",
					"B": "2",
				},
				{
					"C": "3",
					"D": "4",
				},
				{
					"E": "5",
					"F": "6",
				},
			}

			e := New().AddEnvMaps(&envMap)
			Expect(e.Combine()).To(Equal(Map{
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

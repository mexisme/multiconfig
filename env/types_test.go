package env_test

import (
	"github.com/mexisme/multiconfig"
	. "github.com/mexisme/multiconfig/env"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The env type", func() {
	Describe("when adding EnvMaps", func() {
		It("copies the single provided", func() {
			envMap := Map{
				"A": "1",
				"B": "2",
			}

			e := New().AddEnvMap(envMap)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"A": "1",
					"B": "2",
				},
			}))
		})

		It("copies several individually provided", func() {
			envMap1 := Map{
				"A": "1",
				"B": "2",
			}
			envMap2 := Map{
				"C": "3",
				"D": "4",
			}
			envMap3 := Map{
				"E": "5",
				"F": "6",
			}

			e := New().AddEnvMap(envMap1).AddEnvMap(envMap2).AddEnvMap(envMap3)
			Expect(e.EnvMaps).To(Equal(Maps{
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
			}))
		})

		It("copies a provided group", func() {
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
			Expect(e.EnvMaps).To(Equal(Maps{
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
			}))
		})
	})

	Describe("when adding from os.Environ", func() {
		It("copies the single provided", func() {
			environ := []string{
				"AA=11",
				"BB=22",
			}

			e := New().AddOsEnviron(environ)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"AA": "11",
					"BB": "22",
				},
			}))
		})
	})

	Describe("when adding from MultiConfig", func() {
		It("copies the single provided", func() {
			env := multiconfig.New()
			env.AddFromString("lol.env", `A=1
B=2
C=3`)

			e := New().AddMultiConfig(env)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"A": "1",
					"B": "2",
					"C": "3",
				},
			}))
		})
	})
})

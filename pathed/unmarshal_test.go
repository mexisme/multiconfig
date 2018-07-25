package pathed_test

import (
	. "github.com/mexisme/multiconfig/pathed"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("multiconfig/pathed/unmarshal", func() {
	Context("when unmarshalling", func() {
		DescribeTable("",
			func(path, body string, bodyMap BodyMap) {
				if bodyMap == nil {
					bodyMap = BodyMap{
						"A": "  A1A1  ",
						"B": "12",
						"C": "over there",
						"D": "15.1",
					}
				}
				c := New().SetPath(path).SetBody(body)
				m, err := c.ToBodyMap()
				Expect(err).To(Succeed())
				Expect(m).To(Equal(bodyMap))
			},
			Entry("should correctly map a .env", "a/b/c.env", `
A="  A1A1  "
B=12
C=over there
D= 15.1 `, nil),
			Entry("should correctly map a .toml", "a/b/c.toml", `
A="  A1A1  "
B=12
C="over there"
D=15.1
`, nil),
			Entry("should correctly map a .yaml", "a/b/c.yaml", `---
A: "  A1A1  "
B: 12
C:    over there
D: 15.1
`, nil),
			Entry("should correctly map a .json", "a/b/c.json", `{
	"A": "  A1A1  ",
	"B": 12,
	"C": "over there",
	"D": 15.1
}`, nil),
		)
	})
})

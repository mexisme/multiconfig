package multiconfig_test

import (
	"fmt"
	"sort"

	"github.com/mexisme/multiconfig"
	"github.com/mexisme/multiconfig/env"
	"github.com/mexisme/multiconfig/pathed"
)

func Example_merged() {
	env := env.New().SetEnv(env.Envs{
		"A=  A1A1  ",
		"B=12",
		"C=over there",
		"D= 15.1 ",
	})
	c1 := pathed.New().SetPath("/a1.toml").SetBody(`A = "A2"`)
	c2 := pathed.New().SetPath("/b2.yaml").SetBody(`B: "not 12"`)
	config := multiconfig.New().AddItem(env, c1, c2)
	merged, _ := config.Merge()
	fmt.Printf("Merged config: %#v", merged)
	// Output: Merged config: common.BodyMap{"A":"A2", "B":"not 12", "C":"over there", "D":" 15.1 "}
}

func Example_environ() {
	e := env.New().SetEnv(env.Envs{
		"A=  A1A1  ",
		"B=12",
		"C=over there",
		"D= 15.1 ",
	})
	c1 := pathed.New().SetPath("/a1.toml").SetBody(`
A = "A2"
`)
	c2 := pathed.New().SetPath("/b2.yaml").SetBody(`
B: "not 12"
`)
	config := multiconfig.New().AddItem(e, c1, c2)
	merged, _ := config.Merge()

	newEnv := env.New().SetBodyMap(merged)
	newOsEnviron, _ := newEnv.ToOsEnviron()
	sort.Strings(newOsEnviron)
	fmt.Printf("New os.Environ: %#v", newOsEnviron)
	// Output: New os.Environ: env.Envs{"A=A2", "B=not 12", "C=over there", "D= 15.1 "}
}

package env_test

import (
	"fmt"

	"github.com/mexisme/multiconfig/common"
	"github.com/mexisme/multiconfig/env"
)

func Example_fromEnv() {
	e := env.New().SetEnv(env.Envs{
		"A=  A1A1  ",
		"B=12",
		"C=over there",
		"D= 15.1 ",
	})
	bodyMap, _ := e.ToBodyMap()
	fmt.Printf("Parsed to map: %#v", bodyMap)
	// Output: Parsed to map: common.BodyMap{"A":"  A1A1  ", "B":"12", "C":"over there", "D":" 15.1 "}
}

func Example_fromMap() {
	e := env.New().SetBodyMap(common.BodyMap{
		"A": "A1A1  ",
		"B": "12",
		"C": "over there",
		"D": " 15.1 ",
	})
	bodyMap, _ := e.ToOsEnviron()
	fmt.Printf("Parsed to os.Environ: %#v", bodyMap)
	// Output: Parsed to os.Environ: env.Envs{"A=A1A1  ", "B=12", "C=over there", "D= 15.1 "}
}

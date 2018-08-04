package pathed_test

import (
	"fmt"

	"github.com/mexisme/multiconfig/pathed"
)

func Example() {
	path := "/from/this/path/.env"
	body := `
A="  A1A1  "
B=12
C=over there
D= 15.1 `
	config := pathed.New().SetPath(path).SetBody(body)
	bodyMap, _ := config.ToBodyMap()

	fmt.Printf("Read config: %#v", bodyMap)
	// Output: Read config: common.BodyMap{"D":"15.1", "A":"  A1A1  ", "B":"12", "C":"over there"}
}

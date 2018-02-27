package bar2

import (
	"fmt"

	_ "github.com/rogpeppe/test/foo/v2/bar1"
)

func init() {
	fmt.Println("bar2")
}

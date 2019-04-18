package init2

import "fmt"

func init() {
	fmt.Println("init2.init2.go.init1")
}

func init() {
	fmt.Println("init2.init2.go.init2")
}

var Init2 = "var:init2.init2.go.init"

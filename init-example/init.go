package main

import (
	"fmt"

	_ "github.com/yangwenmai/examples/init-example/init0"
	_ "github.com/yangwenmai/examples/init-example/init1"
	"github.com/yangwenmai/examples/init-example/init2"
)

func init() {
	fmt.Println(initStr)
	initStr = "init2"
}

func init() {
	fmt.Println(initStr)
	initStr = "init4"
}

func init() {
	fmt.Println(initStr)
	initStr = "init3"
}

func main() {
	fmt.Println(initStr)
}

var initStr = init2.Init2

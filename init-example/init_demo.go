package main

import "fmt"

var name string

func init() {
	fmt.Println("start...")
	name = "init"
}

func main() {
	fmt.Println(name)
}

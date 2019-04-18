package main

import "fmt"

var _ int64 = a()

func init() {
	fmt.Println("init in a.go")
}
func a() int64 {
	fmt.Println("calling a() in a.go")
	return 2
}

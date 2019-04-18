package main

import "fmt"

var _ int64 = s()

func init() {
	fmt.Println("init in sandbox.go")
}
func s() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}
func main() {
	fmt.Println("main")
}

// 情况一：
// $ go run sandbox.go a.go z.go
// calling s() in sandbox.go
// calling a() in a.go
// calling z() in z.go
// init in sandbox.go
// init in a.go
// init in z.go
// main

// 情况二：
// $ go run a.go z.go sandbox.go
// calling a() in a.go
// calling z() in z.go
// calling s() in sandbox.go
// init in a.go
// init in z.go
// init in sandbox.go
// main

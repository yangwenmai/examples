package main

import "fmt"

func main() {
	m := map[string]string{
		"ok1": "1",
		"ok2": "2",
		"ok3": "3",
		"ok4": "4",
	}
	fmt.Println(len(m), m)
	for k:=range m {
		fmt.Println(k)
	}
	cnt := 1
	for k, _ := range m {
		fmt.Printf("delete ok%d\n", cnt)
		delete(m, k)
		fmt.Println(len(m))
		cnt++
	}
	fmt.Println(len(m), m)
}

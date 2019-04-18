package main

import "fmt"

var precomputed = [20]float64{}

func init() {
	var current float64 = 1
	precomputed[0] = current
	for i := 1; i < len(precomputed); i++ {
		precomputed[i] = precomputed[i-1] * 1.2
	}
}

func main() {
	fmt.Println(precomputed)
}

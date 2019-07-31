package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	fmt.Println(randInt(1, 100))

	v := map[string]string{}
	v["1"] = "1"
	fmt.Println(v["2"])
}

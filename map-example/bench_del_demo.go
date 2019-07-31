package main

import (
"fmt"
"runtime"
"testing"
)

func main() {
	fmt.Printf("last %v\n", testing.Benchmark(benchDelete(getLastIndex)))
	fmt.Printf("some %v\n", testing.Benchmark(benchDelete(getSomeIndex)))
}

const N = 10000

func benchDelete(getIndex func(m map[int]int) int) func(b *testing.B){
	return func(b *testing.B){
		for i := b.N - 1; i >= 0; i-- {
			b.StopTimer()
			runtime.GC()
			m := make(map[int]int)
			for j := 0; j < N; j++ {
				m[j] = 0
			}
			b.StartTimer()
			for j := 0; j < N; j++ {
				// 在旧的 go 版本下才能通过编译
				// m[getIndex(m)] = 0, false
			}
		}
	}
}

func getSomeIndex(m map[int]int) int {
	for k := range m {
		return k
	}
	return 0
}

func getLastIndex(m map[int]int) int {
	return len(m)-1
}

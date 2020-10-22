package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	a := make(map[int]int)
	
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("Before: %5.2fk\n", float64(stats.Alloc)/1024)
	previous := stats.Alloc
	
	for i := 1000000; i > 0; i-- {
		a[i] = 0
	}
	
	for k := range a {
		delete(a, k)
	}
	
	// runtime.GC()
	debug.FreeOSMemory()
	fmt.Printf("len(a)=%v\n", len(a))
	
	runtime.ReadMemStats(stats)
	fmt.Printf("After: %5.2fk    Added: %5.2fk\n", float64(stats.Alloc)/1024, float64(stats.Alloc-previous)/1024)
	
	runtime.GC()
	runtime.ReadMemStats(stats)
	fmt.Printf("After: %5.2fk    Added: %5.2fk\n", float64(stats.Alloc)/1024, float64(stats.Alloc-previous)/1024)
	
	return
}

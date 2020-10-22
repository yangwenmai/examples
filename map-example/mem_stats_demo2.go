package main

import (
	"runtime"
	"log"
	"sync"
)

var intMap *sync.Map
var cnt = 8192

func main() {
	printMemStats()
	
	initMap()
	runtime.GC()
	printMemStats()
	
	log.Println(intMap)
	for i := 0; i < cnt; i++ {
		intMap.Delete(i)
	}
	log.Println(intMap)
	
	runtime.GC()
	printMemStats()
	
	// time.Sleep(1 *time.Minute)
	// intMap = nil
	printMemStats()
	runtime.GC()
	printMemStats()
}

func initMap() {
	intMap = new(sync.Map)
	for i := 0; i < cnt; i++ {
		intMap.Store(i, i)
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
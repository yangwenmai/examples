package main

import (
	"runtime"
	"log"
)

var intMapMap map[int]map[int]int
var cnt = 8192

func main() {
	printMemStats()
	
	initMapMap()
	runtime.GC()
	printMemStats()
	
	fillMapMap()
	runtime.GC()
	printMemStats()
	log.Println(len(intMapMap))
	for i := 0; i < cnt; i++ {
		delete(intMapMap, i)
	}
	log.Println(len(intMapMap))
	
	runtime.GC()
	printMemStats()
	
	intMapMap = nil
	printMemStats()
	runtime.GC()
	printMemStats()
}

func initMapMap() {
	intMapMap = make(map[int]map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		intMapMap[i] = make(map[int]int, cnt)
	}
}

func fillMapMap() {
	for i := 0; i < cnt; i++ {
		for j := 0; j < cnt; j++ {
			intMapMap[i][j] = j
		}
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
package main

var intMap map[int]int
var cnt =100

func initMap() {
	intMap = make(map[int]int, cnt)
	
	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func main() {
	initMap()
	for i:=0;i<100;i++ {
		go delete(intMap, i)
	}
}
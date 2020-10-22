package main

import "testing"

func TestUniq(t *testing.T) {
	list := []string{"fdksfjlds", "1", "3", "4", "flkdsjkf", "2", "3"}
	ret := uniq(list)
	t.Log(ret)
}

// goos: darwin
// goarch: amd64
// pkg: github.com/sundayfun/daycam-server
// BenchmarkUniq
// BenchmarkUniq-8   	 3138302	       333 ns/op
// PASS
func BenchmarkUniq(b *testing.B) {
	list := []string{"fdksfjlds", "1", "3", "4", "flkdsjkf", "2", "3"}
	for i := 0; i < b.N; i++ {
		uniq(list)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/sundayfun/daycam-server
// BenchmarkUniq2
// BenchmarkUniq2-8   	 1715863	       675 ns/op
// PASS
func BenchmarkUniq2(b *testing.B) {
	list := []string{"fdksfjlds", "1", "3", "4", "flkdsjkf", "2", "3"}
	for i := 0; i < b.N; i++ {
		uniq2(list)
	}
}

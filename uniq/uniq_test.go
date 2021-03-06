package main

import "testing"

var list = []string{"fdksfjlds", "1", "fdksfjlds", "3", "4", "flkdsjkf", "2", "3"}

func TestUniq(t *testing.T) {
	ret := uniq(list)
	t.Log(ret)
	ret = uniq2(list)
	t.Log(ret)
	ret = uniq3(list)
	t.Log(ret)
	ret = uniq4(list)
	t.Log(ret)
}

// goos: darwin
// goarch: amd64
// pkg: github.com/sundayfun/daycam-server
// BenchmarkUniq
// BenchmarkUniq-8   	 3138302	       333 ns/op
// PASS
func BenchmarkUniq(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		uniq2(list)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/sundayfun/daycam-server
// BenchmarkUniq3
// BenchmarkUniq3-8   	 2683063	       458 ns/op
// PASS
func BenchmarkUniq3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniq3(list)
	}
}

// BenchmarkUniq4-8   	 3277647	       325 ns/op
func BenchmarkUniq4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniq4(list)
	}
}

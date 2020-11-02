package main

import "testing"

func BenchmarkAdd1(b *testing.B) {
	var vec1 = vec1{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		vec1.add(vec1)
	}
}

func BenchmarkAdd2(b *testing.B) {
	var vec2 = &vec2{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		vec2.add(vec2)
	}
}

func BenchmarkAdd3(b *testing.B) {
	var vec3 = &vec3{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		vec3.add(vec3)
	}
}

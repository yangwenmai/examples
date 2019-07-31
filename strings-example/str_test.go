package main

import (
	"bytes"
	"testing"
)

const (
	shortStr = "short str"
	longStr  = "loooooooooonnnnnnnnnngggggggg string"
)

func BenchmarkShortString(b *testing.B) {
	b.StopTimer()
	var str String
	data := []byte(shortStr)
	r := bytes.NewReader(data)
	l := uint(len(data))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		r.Reset(data)
		str.fill(r, l)
		str.String()
	}
}

func BenchmarkLongString(b *testing.B) {
	b.StopTimer()
	var str String
	data := []byte(longStr)
	r := bytes.NewReader(data)
	l := uint(len(data))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		r.Reset(data)
		str.fill(r, l)
		str.String()
	}
}

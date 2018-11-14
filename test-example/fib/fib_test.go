package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}

func ExampleFib() {
	fmt.Println("ok")
	// output:
	// ok
}

func TestFib2(t *testing.T) {
	var fibTests = []struct {
		in       int // input
		expected int // expected result
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

// func TestFib3(t *testing.T) {
// 	t.Fail()
// 	t.Log("1")
// 	t.Fail()
// 	t.Log("3")
// 	t.FailNow()
// 	t.Log("2")
// }

func TestFib4(t *testing.T) {
	t.Log("0")
	t.Skip("skip")
	t.SkipNow()
	t.Log("1")
}

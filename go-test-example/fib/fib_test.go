package fib

import (
	"fmt"
	"testing"
	"time"
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
// 	// t.Fail()
// 	t.Log("3")
// 	t.FailNow()
// 	t.Log("2")
// }

func TestFib4(t *testing.T) {
	// t.Log("0")
	// t.Skip("skip", "fj")
	// t.SkipNow()
	// t.Fail()
	// // t.FailNow()
	// t.Log("1")

	// t.Parallel()
	t.Parallel()
}

func TestTime(t *testing.T) {
	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Asia/Chongqing", "20:31"},
		{"12:31", "America/New_York", "07:31"},
		{"08:08", "Australia/Sydney", "18:08"},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("%s in %s", tc.gmt, tc.loc)
		t.Log(name)
		t.Run(name, func(t *testing.T) {
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

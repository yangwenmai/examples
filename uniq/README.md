# uniq


```
$ go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out uniq_test.go uniq.go
goos: darwin
goarch: amd64
BenchmarkUniq-8    	 3333010	       307 ns/op	     144 B/op	       2 allocs/op
BenchmarkUniq2-8   	 2068416	       540 ns/op	     240 B/op	       4 allocs/op
BenchmarkUniq3-8   	 2641138	       453 ns/op	     240 B/op	       4 allocs/op
BenchmarkUniq4-8   	 3413756	       354 ns/op	      96 B/op	       1 allocs/op
PASS
ok  	command-line-arguments	6.464s
```

为什么 Uniq4 是最快的呢？


from 

https://go-review.googlesource.com/c/go/+/243941/8/src/go/build/build.go


# uniq


```
$ go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out uniq_test.go uniq.go
goos: darwin
goarch: amd64
BenchmarkUniq-8    	 3948081	       310 ns/op	     144 B/op	       2 allocs/op
BenchmarkUniq2-8   	 2101695	       565 ns/op	     240 B/op	       4 allocs/op
PASS
ok  	command-line-arguments	3.962s
```


from 

https://go-review.googlesource.com/c/go/+/243941/8/src/go/build/build.go


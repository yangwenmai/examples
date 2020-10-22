# uniq


```
$ go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out uniq_test.go uniq.go
goarch: amd64
BenchmarkUniq-8    	 4069275	       291 ns/op	     144 B/op	       2 allocs/op
BenchmarkUniq2-8   	 2184715	       527 ns/op	     240 B/op	       4 allocs/op
BenchmarkUniq3-8   	 2700417	       441 ns/op	     240 B/op	       4 allocs/op
PASS
ok  	command-line-arguments	5.466s
```


from 

https://go-review.googlesource.com/c/go/+/243941/8/src/go/build/build.go


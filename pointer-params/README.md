# pointer-params

go 1.14

```sh
goos: darwin
goarch: amd64
BenchmarkAdd1-8   	1000000000	         0.266 ns/op
BenchmarkAdd2-8   	486800916	         2.37 ns/op
BenchmarkAdd3-8   	485679025	         2.38 ns/op
PASS
ok  	command-line-arguments	3.765s
```

go 1.14

```sh
goos: darwin
goarch: amd64
BenchmarkAdd1-8   	1000000000	         0.283 ns/op
BenchmarkAdd2-8   	469893759	         2.36 ns/op
BenchmarkAdd3-8   	484364077	         2.44 ns/op
PASS
ok  	command-line-arguments	3.546s
```

# Reference

- https://github.com/golang-design/research/issues/1

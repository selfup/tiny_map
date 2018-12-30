[![GoDoc](https://godoc.org/github.com/selfup/tinymap?status.svg)](https://godoc.org/github.com/selfup/tinymap)
[![Build Status](https://travis-ci.org/selfup/tinymap.svg?branch=master)](https://travis-ci.org/selfup/tinymap)

# tinymap

Provides a simple to use interface that behaves like a `HashMap`.

Multiple `Maps` are exposed for use as well as structs that behave like `Tuples`.

This package is intended to be used with `tinygo`.

However it has some nice little structs, so feel free to use this as you wish.

It is very small, but helps with embedded programming.

:tada:

### Basic usage

Just one of the many Maps :smile:

```go
strMap := new(StrMap)

strMap.Set("foo", "bar")
strMap.Get("foo")
strMap.Delete("foo")
```

### Benchmarks

```ocaml
$ ./scripts/bench.sh
+ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/selfup/tinymap
Benchmark_ByteMap_Get_Single_Lower_Bound-4              50000000                25.3 ns/op
Benchmark_ByteMap_Get_Single_Expected_Bound-4           30000000                40.7 ns/op
Benchmark_ByteMap_Get_Max_Size_Upper_Bound-4             3000000               520 ns/op
Benchmark_IntMap_Get_Single_Lower_Bound-4               1000000000               2.51 ns/op
Benchmark_IntMap_Get_Single_Expected_Bound-4            300000000                4.36 ns/op
Benchmark_IntMap_Get_Max_Size_Upper_Bound-4             30000000                44.2 ns/op
Benchmark_IntStrMap_Get_Single_Lower_Bound-4            1000000000               2.70 ns/op
Benchmark_IntStrMap_Get_Single_Expected_Bound-4         300000000                4.46 ns/op
Benchmark_IntStrMap_Get_Max_Size_Upper_Bound-4          30000000                50.3 ns/op
Benchmark_StrMap_Get_Single_Lower_Bound-4               300000000                4.85 ns/op
Benchmark_StrMap_Get_Single_Expected_Bound-4            300000000                5.42 ns/op
Benchmark_StrMap_Get_Max_Size_Upper_Bound-4              5000000               382 ns/op
PASS
ok      github.com/selfup/tinymap       24.219s
```

### Details

1. Using an int as an index is fastest for lookups/comparisons.
1. Using []byte as key is the slowest (makes sense)
1. Upper Bound means the value being grabbed is the 100th element in a slice of 100 elements.
1. Under the hood TinyMap uses slices to store Tuples.
1. I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:

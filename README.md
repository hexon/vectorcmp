# vectorcmp

[![Go Reference](https://pkg.go.dev/badge/github.com/Jille/vectorcmp.svg)](https://pkg.go.dev/github.com/Jille/vectorcmp)

This library provides functions to do comparions on many rows using AVX2's VPCMP instruction. Fallback code is transparantly used if AVX2 is not available.

## VectorEquals

The easiest way to explain is to just show the fallback code.

```go
func goVectorEquals[T uint8 | uint16 | uint32 | uint64](dst []byte, search T, rows []T) {
  for i, v := range rows {
    if search == v {
      dst[i/8] |= 1 << (i % 8)
    }
  }
}
```

Similar to VectorEquals we also implement VectorGreaterThan, VectorLessThan, VectorGreaterEquals and VectorLesserEquals.

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/Jille/vectorcmp
cpu: 13th Gen Intel(R) Core(TM) i9-13900
                  │     purego    │                 asm                  │
                  │    sec/op     │    sec/op     vs base                │
VectorEquals8-32    164.600µ ± 2%   1.208µ ± 16%  -99.27% (p=0.000 n=10)
VectorEquals16-32   125.375µ ± 2%   2.429µ ± 16%  -98.06% (p=0.000 n=10)
VectorEquals32-32   163.694µ ± 3%   4.645µ ± 12%  -97.16% (p=0.000 n=10)
VectorEquals64-32   127.200µ ± 2%   8.495µ ± 15%  -93.32% (p=0.000 n=10)
geomean               144.0µ        3.280µ        -97.72%
```

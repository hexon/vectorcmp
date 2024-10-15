# vectorcmp

[![Go Reference](https://pkg.go.dev/badge/github.com/Jille/vectorcmp.svg)](https://pkg.go.dev/github.com/Jille/vectorcmp)

This library provides functions to do comparions on many rows using AVX2's VPCMP instruction. Fallback code is transparently used if AVX(2) is not available.

## VectorEquals

The easiest way to explain is to just show the fallback code.

```go
func goVectorEquals[T uint8 | uint16 | uint32 | uint64 | float32 | float64](dst []byte, search T, rows []T) {
  clear(dst)
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
pkg: github.com/hexon/vectorcmp
cpu: 13th Gen Intel(R) Core(TM) i9-13900
                              │    purego     │                avx2                 │
                              │    sec/op     │   sec/op     vs base                │
VectorEquals8-32                191.766µ ± 2%   1.240µ ± 4%  -99.35% (p=0.000 n=10)
VectorGreaterThan8-32           189.871µ ± 1%   1.198µ ± 5%  -99.37% (p=0.000 n=10)
VectorLessThan8-32              189.940µ ± 2%   1.234µ ± 3%  -99.35% (p=0.000 n=10)
VectorGreaterEquals8-32         200.345µ ± 1%   1.278µ ± 3%  -99.36% (p=0.000 n=10)
VectorLesserEquals8-32          198.190µ ± 2%   1.302µ ± 3%  -99.34% (p=0.000 n=10)
VectorEquals16-32               187.738µ ± 2%   2.391µ ± 2%  -98.73% (p=0.000 n=10)
VectorGreaterThan16-32          191.861µ ± 1%   2.405µ ± 4%  -98.75% (p=0.000 n=10)
VectorLessThan16-32             189.357µ ± 2%   2.406µ ± 3%  -98.73% (p=0.000 n=10)
VectorGreaterEquals16-32        201.101µ ± 2%   2.485µ ± 3%  -98.76% (p=0.000 n=10)
VectorLesserEquals16-32         197.853µ ± 1%   2.547µ ± 2%  -98.71% (p=0.000 n=10)
VectorEquals32-32               192.640µ ± 1%   4.725µ ± 2%  -97.55% (p=0.000 n=10)
VectorEqualsFloat32-32          209.408µ ± 1%   3.275µ ± 2%  -98.44% (p=0.000 n=10)
VectorGreaterThan32-32          192.983µ ± 1%   3.970µ ± 3%  -97.94% (p=0.000 n=10)
VectorGreaterThanFloat32-32     210.209µ ± 2%   3.264µ ± 1%  -98.45% (p=0.000 n=10)
VectorLessThan32-32             188.972µ ± 1%   3.870µ ± 3%  -97.95% (p=0.000 n=10)
VectorLessThanFloat32-32        211.101µ ± 3%   3.264µ ± 1%  -98.45% (p=0.000 n=10)
VectorGreaterEquals32-32        199.521µ ± 1%   3.955µ ± 9%  -98.02% (p=0.000 n=10)
VectorGreaterEqualsFloat32-32   217.290µ ± 1%   3.270µ ± 1%  -98.50% (p=0.000 n=10)
VectorLesserEquals32-32         200.432µ ± 1%   4.103µ ± 3%  -97.95% (p=0.000 n=10)
VectorLesserEqualsFloat32-32    215.699µ ± 2%   3.285µ ± 1%  -98.48% (p=0.000 n=10)
VectorEquals64-32               194.138µ ± 2%   6.831µ ± 5%  -96.48% (p=0.000 n=10)
VectorEqualsFloat64-32          228.983µ ± 1%   6.431µ ± 3%  -97.19% (p=0.000 n=10)
VectorGreaterThan64-32          191.010µ ± 2%   7.629µ ± 2%  -96.01% (p=0.000 n=10)
VectorGreaterThanFloat64-32     205.197µ ± 1%   6.540µ ± 2%  -96.81% (p=0.000 n=10)
VectorLessThan64-32             192.766µ ± 3%   7.661µ ± 2%  -96.03% (p=0.000 n=10)
VectorLessThanFloat64-32        212.995µ ± 2%   6.423µ ± 1%  -96.98% (p=0.000 n=10)
VectorGreaterEquals64-32        203.198µ ± 2%   7.548µ ± 3%  -96.29% (p=0.000 n=10)
VectorGreaterEqualsFloat64-32   220.453µ ± 2%   6.535µ ± 1%  -97.04% (p=0.000 n=10)
VectorLesserEquals64-32         201.495µ ± 2%   7.744µ ± 4%  -96.16% (p=0.000 n=10)
VectorLesserEqualsFloat64-32    215.303µ ± 3%   6.538µ ± 1%  -96.96% (p=0.000 n=10)
geomean                           201.1µ        3.548µ       -98.24%
```

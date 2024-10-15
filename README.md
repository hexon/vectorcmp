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
                              │     purego     │                avx2                 │
                              │     sec/op     │   sec/op     vs base                │
VectorEquals8-32                191765.5n ± 2%   912.5n ± 2%  -99.52% (p=0.000 n=10)
VectorGreaterThan8-32           189871.0n ± 1%   899.4n ± 1%  -99.53% (p=0.000 n=10)
VectorLessThan8-32              189939.5n ± 2%   903.3n ± 1%  -99.52% (p=0.000 n=10)
VectorGreaterEquals8-32         200345.0n ± 1%   969.2n ± 3%  -99.52% (p=0.000 n=10)
VectorLesserEquals8-32          198189.5n ± 2%   973.8n ± 3%  -99.51% (p=0.000 n=10)
VectorEquals16-32                187.738µ ± 2%   1.892µ ± 7%  -98.99% (p=0.000 n=10)
VectorGreaterThan16-32           191.861µ ± 1%   1.890µ ± 6%  -99.01% (p=0.000 n=10)
VectorLessThan16-32              189.357µ ± 2%   1.894µ ± 4%  -99.00% (p=0.000 n=10)
VectorGreaterEquals16-32         201.101µ ± 2%   2.003µ ± 5%  -99.00% (p=0.000 n=10)
VectorLesserEquals16-32          197.853µ ± 1%   1.991µ ± 5%  -98.99% (p=0.000 n=10)
VectorEquals32-32                192.640µ ± 1%   3.685µ ± 5%  -98.09% (p=0.000 n=10)
VectorEqualsFloat32-32           209.408µ ± 1%   3.156µ ± 1%  -98.49% (p=0.000 n=10)
VectorGreaterThan32-32           192.983µ ± 1%   3.717µ ± 2%  -98.07% (p=0.000 n=10)
VectorGreaterThanFloat32-32      210.209µ ± 2%   3.156µ ± 1%  -98.50% (p=0.000 n=10)
VectorLessThan32-32              188.972µ ± 1%   3.729µ ± 4%  -98.03% (p=0.000 n=10)
VectorLessThanFloat32-32         211.101µ ± 3%   3.137µ ± 2%  -98.51% (p=0.000 n=10)
VectorGreaterEquals32-32         199.521µ ± 1%   3.859µ ± 4%  -98.07% (p=0.000 n=10)
VectorGreaterEqualsFloat32-32    217.290µ ± 1%   3.179µ ± 1%  -98.54% (p=0.000 n=10)
VectorLesserEquals32-32          200.432µ ± 1%   3.907µ ± 3%  -98.05% (p=0.000 n=10)
VectorLesserEqualsFloat32-32     215.699µ ± 2%   3.167µ ± 2%  -98.53% (p=0.000 n=10)
VectorEquals64-32                194.138µ ± 2%   6.628µ ± 4%  -96.59% (p=0.000 n=10)
VectorEqualsFloat64-32           228.983µ ± 1%   6.341µ ± 1%  -97.23% (p=0.000 n=10)
VectorGreaterThan64-32           191.010µ ± 2%   7.453µ ± 3%  -96.10% (p=0.000 n=10)
VectorGreaterThanFloat64-32      205.197µ ± 1%   6.357µ ± 1%  -96.90% (p=0.000 n=10)
VectorLessThan64-32              192.766µ ± 3%   7.443µ ± 3%  -96.14% (p=0.000 n=10)
VectorLessThanFloat64-32         212.995µ ± 2%   6.323µ ± 1%  -97.03% (p=0.000 n=10)
VectorGreaterEquals64-32         203.198µ ± 2%   7.407µ ± 3%  -96.35% (p=0.000 n=10)
VectorGreaterEqualsFloat64-32    220.453µ ± 2%   6.391µ ± 2%  -97.10% (p=0.000 n=10)
VectorLesserEquals64-32          201.495µ ± 2%   7.541µ ± 3%  -96.26% (p=0.000 n=10)
VectorLesserEqualsFloat64-32     215.303µ ± 3%   6.388µ ± 2%  -97.03% (p=0.000 n=10)
geomean                            201.1µ        3.160µ       -98.43%
```

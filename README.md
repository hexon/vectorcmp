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

Similar to VectorEquals we also implement VectorNotEquals, VectorGreaterThan, VectorLessThan, VectorGreaterEquals and VectorLesserEquals.

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/hexon/vectorcmp
cpu: 14th Gen Intel(R) Core(TM) i9-14900
                              │     purego     │                avx2                 │
                              │     sec/op     │   sec/op     vs base                │
VectorEquals8-32                145594.5n ± 2%   628.0n ± 2%  -99.57% (p=0.000 n=10)
VectorNotEquals8-32             141155.0n ± 1%   679.7n ± 4%  -99.52% (p=0.000 n=10)
VectorGreaterThan8-32           146456.5n ± 2%   856.9n ± 2%  -99.41% (p=0.000 n=10)
VectorLessThan8-32              144820.5n ± 1%   858.8n ± 2%  -99.41% (p=0.000 n=10)
VectorGreaterEquals8-32         139554.0n ± 2%   928.1n ± 1%  -99.33% (p=0.000 n=10)
VectorLesserEquals8-32          141419.5n ± 1%   917.0n ± 3%  -99.35% (p=0.000 n=10)
VectorEquals16-32                151.078µ ± 1%   1.325µ ± 2%  -99.12% (p=0.000 n=10)
VectorNotEquals16-32             139.751µ ± 3%   1.359µ ± 4%  -99.03% (p=0.000 n=10)
VectorGreaterThan16-32           150.466µ ± 2%   1.682µ ± 1%  -98.88% (p=0.000 n=10)
VectorLessThan16-32              151.316µ ± 1%   1.700µ ± 2%  -98.88% (p=0.000 n=10)
VectorGreaterEquals16-32         139.931µ ± 2%   1.728µ ± 2%  -98.77% (p=0.000 n=10)
VectorLesserEquals16-32          138.835µ ± 1%   1.760µ ± 3%  -98.73% (p=0.000 n=10)
VectorEquals32-32                146.441µ ± 2%   2.626µ ± 1%  -98.21% (p=0.000 n=10)
VectorEqualsFloat32-32           153.134µ ± 1%   2.215µ ± 1%  -98.55% (p=0.000 n=10)
VectorNotEquals32-32             140.962µ ± 2%   2.663µ ± 2%  -98.11% (p=0.000 n=10)
VectorNotEqualsFloat32-32        160.325µ ± 2%   2.195µ ± 1%  -98.63% (p=0.000 n=10)
VectorGreaterThan32-32           144.302µ ± 3%   3.333µ ± 2%  -97.69% (p=0.000 n=10)
VectorGreaterThanFloat32-32      146.653µ ± 3%   2.209µ ± 2%  -98.49% (p=0.000 n=10)
VectorLessThan32-32              143.874µ ± 2%   3.348µ ± 3%  -97.67% (p=0.000 n=10)
VectorLessThanFloat32-32         147.214µ ± 1%   2.221µ ± 1%  -98.49% (p=0.000 n=10)
VectorGreaterEquals32-32         139.414µ ± 2%   3.506µ ± 0%  -97.49% (p=0.000 n=10)
VectorGreaterEqualsFloat32-32    151.919µ ± 1%   2.226µ ± 1%  -98.54% (p=0.000 n=10)
VectorLesserEquals32-32          142.315µ ± 3%   3.447µ ± 3%  -97.58% (p=0.000 n=10)
VectorLesserEqualsFloat32-32     152.132µ ± 2%   2.188µ ± 1%  -98.56% (p=0.000 n=10)
VectorEquals64-32                146.577µ ± 1%   4.468µ ± 3%  -96.95% (p=0.000 n=10)
VectorEqualsFloat64-32           152.734µ ± 2%   4.426µ ± 2%  -97.10% (p=0.000 n=10)
VectorNotEquals64-32             141.699µ ± 2%   4.506µ ± 5%  -96.82% (p=0.000 n=10)
VectorNotEqualsFloat64-32        164.610µ ± 2%   4.443µ ± 1%  -97.30% (p=0.000 n=10)
VectorGreaterThan64-32           146.472µ ± 2%   6.420µ ± 2%  -95.62% (p=0.000 n=10)
VectorGreaterThanFloat64-32      149.198µ ± 2%   4.423µ ± 1%  -97.04% (p=0.000 n=10)
VectorLessThan64-32              145.543µ ± 2%   6.190µ ± 3%  -95.75% (p=0.000 n=10)
VectorLessThanFloat64-32         146.221µ ± 2%   4.373µ ± 1%  -97.01% (p=0.000 n=10)
VectorGreaterEquals64-32         139.420µ ± 2%   6.651µ ± 2%  -95.23% (p=0.000 n=10)
VectorGreaterEqualsFloat64-32    151.540µ ± 2%   4.426µ ± 1%  -97.08% (p=0.000 n=10)
VectorLesserEquals64-32          139.964µ ± 1%   6.508µ ± 3%  -95.35% (p=0.000 n=10)
VectorLesserEqualsFloat64-32     151.295µ ± 1%   4.482µ ± 2%  -97.04% (p=0.000 n=10)
VectorIsNaNFloat32-32             23.555µ ± 1%   2.228µ ± 1%  -90.54% (p=0.000 n=10)
VectorIsNaNFloat64-32             23.687µ ± 1%   4.410µ ± 1%  -81.38% (p=0.000 n=10)
geomean                            133.0µ        2.492µ       -98.13%
```

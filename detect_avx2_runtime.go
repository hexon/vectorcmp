//go:build amd64 && !amd64.v3 && !purego && !vectorcmp_testing_avx

package vectorcmp

import "golang.org/x/sys/cpu"

func hasAVX() bool {
	return cpu.X86.HasAVX
}

func hasAVX2() bool {
	return cpu.X86.HasAVX2
}

func hasBMI2() bool {
	return cpu.X86.HasBMI2
}

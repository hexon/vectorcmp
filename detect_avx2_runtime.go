//go:build amd64 && !amd64.v3 && !purego && !vectorcmp_testing_avx

package vectorcmp

import "golang.org/x/sys/cpu"

var axv2AndBmi2 = cpu.X86.HasAVX2 && cpu.X86.HasBMI2

func hasAVX() bool {
	return cpu.X86.HasAVX
}

func hasAVX2AndBMI2() bool {
	return axv2AndBmi2
}

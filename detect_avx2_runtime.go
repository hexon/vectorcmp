//go:build amd64 && !amd64.v3 && !purego

package vectorcmp

import "golang.org/x/sys/cpu"

func hasAVX() bool {
	return cpu.X86.HasAVX
}

func hasAVX2() bool {
	return cpu.X86.HasAVX2
}

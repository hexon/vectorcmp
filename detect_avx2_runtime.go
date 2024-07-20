//go:build amd64 && !amd64.v3

package vectorcmp

import "golang.org/x/sys/cpu"

func hasAVX2() bool {
	return cpu.X86.HasAVX2
}

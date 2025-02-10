//go:build (!amd64 || purego) && !vectorcmp_testing_avx

package vectorcmp

func hasAVX() bool {
	return false
}

func hasAVX2() bool {
	return false
}

func hasBMI2() bool {
	return false
}

//go:build (!amd64 || purego) && !vectorcmp_testing_avx

package vectorcmp

func hasAVX() bool {
	return false
}

func hasAVX2AndBMI2() bool {
	return false
}

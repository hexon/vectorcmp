//go:build vectorcmp_testing_avx

package vectorcmp

func hasAVX() bool {
	return true
}

func hasAVX2() bool {
	return false
}

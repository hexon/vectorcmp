//go:build amd64.v3 && !purego && !vectorcmp_testing_avx

package vectorcmp

func hasAVX() bool {
	return true
}

func hasAVX2() bool {
	return true
}

func hasBMI2() bool {
	return true
}

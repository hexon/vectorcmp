//go:build !amd64 || purego

package vectorcmp

func hasAVX() bool {
	return false
}

func hasAVX2() bool {
	return false
}

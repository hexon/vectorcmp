//go:build !amd64

package vectorcmp

func hasAVX() bool {
	return false
}

func hasAVX2() bool {
	return false
}

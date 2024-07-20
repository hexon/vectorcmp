//go:build !amd64

package vectorcmp

func hasAVX2() bool {
	return false
}

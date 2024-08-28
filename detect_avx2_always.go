//go:build amd64.v3

package vectorcmp

func hasAVX() bool {
	return true
}

func hasAVX2() bool {
	return true
}

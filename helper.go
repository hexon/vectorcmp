package vectorcmp

func boundsCheck[T any](dst []byte, rows []T) {
	if len(dst)*8 < len(rows) {
		panic("vectorcmp: given bitmap is not large enough to hold bits for all rows")
	}
}

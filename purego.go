package vectorcmp

func goVectorEquals[T uint8 | uint16 | uint32 | uint64](dst []byte, search T, rows []T) {
	zero(dst)
	for i, v := range rows {
		if search == v {
			dst[i/8] |= 1 << (i % 8)
		}
	}
}

func goVectorGreaterThan[T uint8 | uint16 | uint32 | uint64](dst []byte, search T, rows []T) {
	zero(dst)
	for i, v := range rows {
		if search > v {
			dst[i/8] |= 1 << (i % 8)
		}
	}
}

func goVectorLessThan[T uint8 | uint16 | uint32 | uint64](dst []byte, search T, rows []T) {
	zero(dst)
	for i, v := range rows {
		if search < v {
			dst[i/8] |= 1 << (i % 8)
		}
	}
}

func goVectorGreaterEquals[T uint8 | uint16 | uint32 | uint64](dst []byte, search T, rows []T) {
	zero(dst)
	for i, v := range rows {
		if search >= v {
			dst[i/8] |= 1 << (i % 8)
		}
	}
}

func goVectorLesserEquals[T uint8 | uint16 | uint32 | uint64](dst []byte, search T, rows []T) {
	zero(dst)
	for i, v := range rows {
		if search <= v {
			dst[i/8] |= 1 << (i % 8)
		}
	}
}

func zero(dst []byte) {
	for i := range dst {
		dst[i] = 0
	}
}

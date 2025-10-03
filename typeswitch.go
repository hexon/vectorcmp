package vectorcmp

// VectorEquals compares each entry in rows to search and enables the corresponding bit in dst if equal.
// Prefer calling VectorEquals8/16/32/64 directly.
//
//	clear(dst)
//	for i, v := range rows {
//	  if search == v {
//	    dst[i/8] |= 1 << (i % 8)
//	  }
//	}
func VectorEquals[T uint8 | uint16 | uint32 | uint64 | float32 | float64](dst []byte, search T, rows []T) {
	switch rows := any(rows).(type) {
	case []uint8:
		VectorEquals8(dst, uint8(search), rows)
	case []uint16:
		VectorEquals16(dst, uint16(search), rows)
	case []uint32:
		VectorEquals32(dst, uint32(search), rows)
	case []uint64:
		VectorEquals64(dst, uint64(search), rows)
	case []float32:
		VectorEqualsFloat32(dst, float32(search), rows)
	case []float64:
		VectorEqualsFloat64(dst, float64(search), rows)
	default:
		panic("generic type not listed in type switch")
	}
}

// VectorNotEquals compares each entry in rows to search and enables the corresponding bit in dst if not equal.
// Prefer calling VectorNotEquals8/16/32/64 directly.
//
//	clear(dst)
//	for i, v := range rows {
//	  if search != v {
//	    dst[i/8] |= 1 << (i % 8)
//	  }
//	}
func VectorNotEquals[T uint8 | uint16 | uint32 | uint64 | float32 | float64](dst []byte, search T, rows []T) {
	switch rows := any(rows).(type) {
	case []uint8:
		VectorNotEquals8(dst, uint8(search), rows)
	case []uint16:
		VectorNotEquals16(dst, uint16(search), rows)
	case []uint32:
		VectorNotEquals32(dst, uint32(search), rows)
	case []uint64:
		VectorNotEquals64(dst, uint64(search), rows)
	case []float32:
		VectorNotEqualsFloat32(dst, float32(search), rows)
	case []float64:
		VectorNotEqualsFloat64(dst, float64(search), rows)
	default:
		panic("generic type not listed in type switch")
	}
}

// VectorGreaterThan compares each entry in rows to search and enables the corresponding bit in dst if search > rows[i].
// Prefer calling VectorGreaterThan8/16/32/64 directly.
//
//	clear(dst)
//	for i, v := range rows {
//	  if search > v {
//	    dst[i/8] |= 1 << (i % 8)
//	  }
//	}
func VectorGreaterThan[T uint8 | uint16 | uint32 | uint64 | float32 | float64](dst []byte, search T, rows []T) {
	switch rows := any(rows).(type) {
	case []uint8:
		VectorGreaterThan8(dst, uint8(search), rows)
	case []uint16:
		VectorGreaterThan16(dst, uint16(search), rows)
	case []uint32:
		VectorGreaterThan32(dst, uint32(search), rows)
	case []uint64:
		VectorGreaterThan64(dst, uint64(search), rows)
	case []float32:
		VectorGreaterThanFloat32(dst, float32(search), rows)
	case []float64:
		VectorGreaterThanFloat64(dst, float64(search), rows)
	default:
		panic("generic type not listed in type switch")
	}
}

// VectorLessThan compares each entry in rows to search and enables the corresponding bit in dst if search < rows[i].
// Prefer calling VectorLessThan8/16/32/64 directly.
//
//	clear(dst)
//	for i, v := range rows {
//	  if search < v {
//	    dst[i/8] |= 1 << (i % 8)
//	  }
//	}
func VectorLessThan[T uint8 | uint16 | uint32 | uint64 | float32 | float64](dst []byte, search T, rows []T) {
	switch rows := any(rows).(type) {
	case []uint8:
		VectorLessThan8(dst, uint8(search), rows)
	case []uint16:
		VectorLessThan16(dst, uint16(search), rows)
	case []uint32:
		VectorLessThan32(dst, uint32(search), rows)
	case []uint64:
		VectorLessThan64(dst, uint64(search), rows)
	case []float32:
		VectorLessThanFloat32(dst, float32(search), rows)
	case []float64:
		VectorLessThanFloat64(dst, float64(search), rows)
	default:
		panic("generic type not listed in type switch")
	}
}

// VectorGreaterEquals compares each entry in rows to search and enables the corresponding bit in dst if search >= rows[i].
// Prefer calling VectorGreaterEquals8/16/32/64 directly.
//
//	clear(dst)
//	for i, v := range rows {
//	  if search >= v {
//	    dst[i/8] |= 1 << (i % 8)
//	  }
//	}
func VectorGreaterEquals[T uint8 | uint16 | uint32 | uint64 | float32 | float64](dst []byte, search T, rows []T) {
	switch rows := any(rows).(type) {
	case []uint8:
		VectorGreaterEquals8(dst, uint8(search), rows)
	case []uint16:
		VectorGreaterEquals16(dst, uint16(search), rows)
	case []uint32:
		VectorGreaterEquals32(dst, uint32(search), rows)
	case []uint64:
		VectorGreaterEquals64(dst, uint64(search), rows)
	case []float32:
		VectorGreaterEqualsFloat32(dst, float32(search), rows)
	case []float64:
		VectorGreaterEqualsFloat64(dst, float64(search), rows)
	default:
		panic("generic type not listed in type switch")
	}
}

// VectorLesserEquals compares each entry in rows to search and enables the corresponding bit in dst if search <= rows[i].
// Prefer calling VectorLesserEquals8/16/32/64 directly.
//
//	clear(dst)
//	for i, v := range rows {
//	  if search <= v {
//	    dst[i/8] |= 1 << (i % 8)
//	  }
//	}
func VectorLesserEquals[T uint8 | uint16 | uint32 | uint64 | float32 | float64](dst []byte, search T, rows []T) {
	switch rows := any(rows).(type) {
	case []uint8:
		VectorLesserEquals8(dst, uint8(search), rows)
	case []uint16:
		VectorLesserEquals16(dst, uint16(search), rows)
	case []uint32:
		VectorLesserEquals32(dst, uint32(search), rows)
	case []uint64:
		VectorLesserEquals64(dst, uint64(search), rows)
	case []float32:
		VectorLesserEqualsFloat32(dst, float32(search), rows)
	case []float64:
		VectorLesserEqualsFloat64(dst, float64(search), rows)
	default:
		panic("generic type not listed in type switch")
	}
}

// Code generated by avo.go. DO NOT EDIT.

//go:build amd64 && !purego

package vectorcmp

func VectorEquals8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmAVX2EqualsUint8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	} else if hasAVX() && len(rows) >= 32 {
		asmAVXEqualsUint8(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmAVX2GreaterThanUint8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	} else if hasAVX() && len(rows) >= 32 {
		asmAVXGreaterThanUint8(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmAVX2LessThanUint8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	} else if hasAVX() && len(rows) >= 32 {
		asmAVXLessThanUint8(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmAVX2GreaterEqualsUint8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	} else if hasAVX() && len(rows) >= 32 {
		asmAVXGreaterEqualsUint8(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmAVX2LesserEqualsUint8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	} else if hasAVX() && len(rows) >= 32 {
		asmAVXLesserEqualsUint8(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorEquals16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmAVX2EqualsUint16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	} else if hasAVX() && len(rows) >= 16 {
		asmAVXEqualsUint16(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmAVX2GreaterThanUint16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	} else if hasAVX() && len(rows) >= 16 {
		asmAVXGreaterThanUint16(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmAVX2LessThanUint16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	} else if hasAVX() && len(rows) >= 16 {
		asmAVXLessThanUint16(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmAVX2GreaterEqualsUint16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	} else if hasAVX() && len(rows) >= 16 {
		asmAVXGreaterEqualsUint16(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmAVX2LesserEqualsUint16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	} else if hasAVX() && len(rows) >= 16 {
		asmAVXLesserEqualsUint16(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorEquals32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2EqualsUint32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXEqualsUint32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorEqualsFloat32(dstMask []byte, b float32, rows []float32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2EqualsFloat32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXEqualsFloat32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterThanUint32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterThanUint32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorGreaterThanFloat32(dstMask []byte, b float32, rows []float32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterThanFloat32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterThanFloat32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LessThanUint32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLessThanUint32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorLessThanFloat32(dstMask []byte, b float32, rows []float32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LessThanFloat32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLessThanFloat32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterEqualsUint32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterEqualsUint32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorGreaterEqualsFloat32(dstMask []byte, b float32, rows []float32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterEqualsFloat32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterEqualsFloat32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LesserEqualsUint32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLesserEqualsUint32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorLesserEqualsFloat32(dstMask []byte, b float32, rows []float32) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LesserEqualsFloat32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLesserEqualsFloat32(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorEquals64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2EqualsUint64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXEqualsUint64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorEqualsFloat64(dstMask []byte, b float64, rows []float64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2EqualsFloat64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXEqualsFloat64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterThanUint64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterThanUint64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorGreaterThanFloat64(dstMask []byte, b float64, rows []float64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterThanFloat64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterThanFloat64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LessThanUint64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLessThanUint64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorLessThanFloat64(dstMask []byte, b float64, rows []float64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LessThanFloat64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLessThanFloat64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterEqualsUint64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterEqualsUint64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorGreaterEqualsFloat64(dstMask []byte, b float64, rows []float64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2GreaterEqualsFloat64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXGreaterEqualsFloat64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LesserEqualsUint64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLesserEqualsUint64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorLesserEqualsFloat64(dstMask []byte, b float64, rows []float64) {
	if hasAVX2() && len(rows) >= 16 {
		asmAVX2LesserEqualsFloat64(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	} else if hasAVX() && len(rows) >= 8 {
		asmAVXLesserEqualsFloat64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}

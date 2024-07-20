// Code generated by avo.go. DO NOT EDIT.

//go:build amd64 && !purego

package vectorcmp

func VectorEquals8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmBulkEquals8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmBulkGreaterThan8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmBulkLessThan8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmBulkGreaterEquals8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals8(dstMask []byte, b uint8, rows []uint8) {
	if hasAVX2() && len(rows) >= 64 {
		asmBulkLesserEquals8(dstMask, b, rows[:len(rows) & ^63])
		dstMask = dstMask[(len(rows) & ^63)/8:]
		rows = rows[len(rows) & ^63:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorEquals16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmBulkEquals16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmBulkGreaterThan16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmBulkLessThan16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmBulkGreaterEquals16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals16(dstMask []byte, b uint16, rows []uint16) {
	if hasAVX2() && len(rows) >= 32 {
		asmBulkLesserEquals16(dstMask, b, rows[:len(rows) & ^31])
		dstMask = dstMask[(len(rows) & ^31)/8:]
		rows = rows[len(rows) & ^31:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorEquals32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmBulkEquals32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmBulkGreaterThan32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmBulkLessThan32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmBulkGreaterEquals32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals32(dstMask []byte, b uint32, rows []uint32) {
	if hasAVX2() && len(rows) >= 16 {
		asmBulkLesserEquals32(dstMask, b, rows[:len(rows) & ^15])
		dstMask = dstMask[(len(rows) & ^15)/8:]
		rows = rows[len(rows) & ^15:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
func VectorEquals64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 8 {
		asmBulkEquals64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorEquals(dstMask, b, rows)
}
func VectorGreaterThan64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 8 {
		asmBulkGreaterThan64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterThan(dstMask, b, rows)
}
func VectorLessThan64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 8 {
		asmBulkLessThan64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLessThan(dstMask, b, rows)
}
func VectorGreaterEquals64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 8 {
		asmBulkGreaterEquals64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorGreaterEquals(dstMask, b, rows)
}
func VectorLesserEquals64(dstMask []byte, b uint64, rows []uint64) {
	if hasAVX2() && len(rows) >= 8 {
		asmBulkLesserEquals64(dstMask, b, rows[:len(rows) & ^7])
		dstMask = dstMask[(len(rows) & ^7)/8:]
		rows = rows[len(rows) & ^7:]
	}
	goVectorLesserEquals(dstMask, b, rows)
}
package vectorcmp

import (
	"bytes"
	"math/rand"
	"testing"
)

func randomBuffer[T uint8 | uint16 | uint32 | uint64]() []T {
	buf := make([]T, 64*1024+rand.Intn(2048))
	for i := range buf {
		buf[i] = T('a' + rand.Intn(5))
	}
	return buf
}

func destinationBuffer[T uint8 | uint16 | uint32 | uint64](buf []T) []byte {
	return make([]byte, (len(buf)+7)/8)
}

func TestVectorEquals8(t *testing.T) {
	buf := randomBuffer[uint8]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorEquals8(dstAsm, 'a', buf)
	goVectorEquals(dstOld, 'a', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%b\n%b", dstAsm, dstOld)
	}
}

func TestVectorEquals16(t *testing.T) {
	buf := randomBuffer[uint16]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorEquals16(dstAsm, 'a', buf)
	goVectorEquals(dstOld, 'a', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%x\n%x", dstAsm, dstOld)
	}
}

func TestVectorEquals32(t *testing.T) {
	buf := randomBuffer[uint32]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorEquals32(dstAsm, 'a', buf)
	goVectorEquals(dstOld, 'a', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%b\n%b", dstAsm, dstOld)
	}
}

func TestVectorEquals64(t *testing.T) {
	buf := randomBuffer[uint64]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorEquals64(dstAsm, 'a', buf)
	goVectorEquals(dstOld, 'a', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%b\n%b", dstAsm, dstOld)
	}
}

func TestVectorGreaterThan16(t *testing.T) {
	buf := randomBuffer[uint16]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorGreaterThan16(dstAsm, 'b', buf)
	goVectorGreaterThan(dstOld, 'b', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%x\n%x", dstAsm, dstOld)
	}
}

func TestVectorLessThan16(t *testing.T) {
	buf := randomBuffer[uint16]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorLessThan16(dstAsm, 'b', buf)
	goVectorLessThan(dstOld, 'b', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%x\n%x", dstAsm, dstOld)
	}
}

func TestVectorGreaterEquals16(t *testing.T) {
	buf := randomBuffer[uint16]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorGreaterEquals16(dstAsm, 'b', buf)
	goVectorGreaterEquals(dstOld, 'b', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%x\n%x", dstAsm, dstOld)
	}
}

func TestVectorLesserEquals16(t *testing.T) {
	buf := randomBuffer[uint16]()
	dstAsm := destinationBuffer(buf)
	dstOld := destinationBuffer(buf)
	VectorLesserEquals16(dstAsm, 'b', buf)
	goVectorLesserEquals(dstOld, 'b', buf)
	if !bytes.Equal(dstAsm, dstOld) {
		t.Fatalf("ASM version returned a different result:\n%x\n%x", dstAsm, dstOld)
	}
}

func BenchmarkAsm8(b *testing.B) {
	buf := randomBuffer[uint8]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		VectorEquals8(dst, 'a', buf)
	}
}

func BenchmarkGo8(b *testing.B) {
	buf := randomBuffer[uint8]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		goVectorEquals(dst, 'a', buf)
	}
}

func BenchmarkAsm16(b *testing.B) {
	buf := randomBuffer[uint16]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		VectorEquals16(dst, 'a', buf)
	}
}

func BenchmarkGo16(b *testing.B) {
	buf := randomBuffer[uint16]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		goVectorEquals(dst, 'a', buf)
	}
}

func BenchmarkAsm32(b *testing.B) {
	buf := randomBuffer[uint32]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		VectorEquals32(dst, 'a', buf)
	}
}

func BenchmarkGo32(b *testing.B) {
	buf := randomBuffer[uint32]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		goVectorEquals(dst, 'a', buf)
	}
}

func BenchmarkAsm64(b *testing.B) {
	buf := randomBuffer[uint64]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		VectorEquals64(dst, 'a', buf)
	}
}

func BenchmarkGo64(b *testing.B) {
	buf := randomBuffer[uint64]()
	dst := destinationBuffer(buf)
	b.ResetTimer()
	for i := 0; b.N > i; i++ {
		goVectorEquals(dst, 'a', buf)
	}
}

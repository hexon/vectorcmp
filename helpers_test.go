package vectorcmp

// Test helpers

import (
	"math"
	"math/rand"
)

func randomBuffer[T uint8 | uint16 | uint32 | uint64 | float32 | float64]() []T {
	buf := make([]T, 64*1024+rand.Intn(2048))
	for i := range buf {
		buf[i] = T('a' + rand.Intn(3))
	}
	x := T(0)
	switch b := any(buf).(type) {
	case []float32:
		for i := 0; 20 > i; i++ {
			b[rand.Intn(len(b))] = float32(math.NaN())
		}
	case []float64:
		for i := 0; 20 > i; i++ {
			b[rand.Intn(len(b))] = math.NaN()
		}
	default:
		for i := 0; 20 > i; i++ {
			buf[rand.Intn(len(buf))] = x - 1 // The maximum value this uint supports.
		}
	}
	return buf
}

func destinationBuffer[T uint8 | uint16 | uint32 | uint64 | float32 | float64](buf []T) []byte {
	return make([]byte, (len(buf)+7)/8)
}

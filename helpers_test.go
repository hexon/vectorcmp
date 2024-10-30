package vectorcmp

import (
	"math"
	"math/rand"
)

func randomBuffer[T uint8 | uint16 | uint32 | uint64 | float32 | float64]() []T {
	buf := make([]T, 64*1024+rand.Intn(2048))
	for i := range buf {
		buf[i] = T('a' + rand.Intn(3))
	}
	switch buf := any(buf).(type) {
	case []float32:
		for i := 0; 20 > i; i++ {
			buf[rand.Intn(len(buf))] = float32(math.NaN())
		}
	case []float64:
		for i := 0; 20 > i; i++ {
			buf[rand.Intn(len(buf))] = math.NaN()
		}
	}
	return buf
}

func destinationBuffer[T uint8 | uint16 | uint32 | uint64 | float32 | float64](buf []T) []byte {
	return make([]byte, (len(buf)+7)/8)
}

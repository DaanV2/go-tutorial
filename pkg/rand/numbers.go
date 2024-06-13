package rand

import "math/rand"

func Int64(n int) int64 {
	return rand.Int63n(int64(n))
}

func Float64(n float64) float64 {
	return rand.Float64() * n
}
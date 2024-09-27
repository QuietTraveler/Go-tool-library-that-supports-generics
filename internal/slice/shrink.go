package slice

const (
	minCap = 64    // minimum capacity
	maxCap = 2048  // maximum capacity
	factor = 0.625 // shrink factor
)

func calCapacity(c, l int) (int, bool) {
	if c <= minCap {
		return c, false
	}
	if c > maxCap && (c/l >= 2) {
		return int(float32(c) * float32(factor)), true
	}
	if c <= maxCap && (c/l >= 4) {
		return c / 2, true
	}
	return c, false
}

// Shrink Reduce capacity
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, ok := calCapacity(c, l)
	if !ok {
		return src
	}
	newSlice := make([]T, 0, n)
	return append(newSlice, src...)
}

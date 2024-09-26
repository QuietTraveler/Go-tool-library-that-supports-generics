package slice

import "basic-go/gtool/internal/slice"

// Delete the element at the specified index
func Delete[T any](src []T, index int) ([]T, T, error) {
	return slice.Delete[T](src, index)
}

// FilterDelete Delete elements that match the criteria
// Returns a slice of elements that do not match the criteria
// All operations are performed on the original slices and no new slices will be created.
func FilterDelete[T any](src []T, condition func(index int, src T) bool) []T {
	emptyPos := 0 // The position of the empty element
	for index := range src {
		if condition(index, src[index]) {
			continue
		}
		// Move the non-empty element to the empty position
		src[emptyPos] = src[index]
		emptyPos++
	}
	return src[:emptyPos]
}

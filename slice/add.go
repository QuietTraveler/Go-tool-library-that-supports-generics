package slice

import "gtool/internal/slice"

// Add adds an element at index
// The range of index should be [0,len(src)]
// if index == len(src), then add element at the end of the slice
func Add[T any](src []T, element T, index int) ([]T, error) {
	return slice.Add[T](src, element, index)
}

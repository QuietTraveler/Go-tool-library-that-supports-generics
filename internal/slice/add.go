package slice

import "basic-go/gtool/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	//Boundary judgment
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	// If index is equal to length, add to the end
	if index == length {
		return append(src, element), nil
	} else { // If index is not equal to length, add to the middle
		latterSlice := append([]T{element}, src[index:]...)
		return append(src[:index], latterSlice...), nil
	}
}

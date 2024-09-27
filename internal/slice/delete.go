package slice

import "gtool/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	var target T // Declare a variable to store the deleted element
	if index < 0 || index >= length {
		return nil, target, errs.NewErrIndexOutOfRange(length, index)
	}
	target = src[index]
	if index == length-1 {
		return src[:index], target, nil
	} else if index == 0 {
		return src[1:], target, nil
	}
	return append(src[:index], src[index+1:]...), target, nil
}

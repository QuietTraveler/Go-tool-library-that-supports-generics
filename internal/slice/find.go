package slice

import "gtool/internal/slice/doc"

// Find the first element that satisfies the condition
// if no exit, return false
func Find[T any](src []T, match doc.MatchFunc[T]) (T, bool) {
	for _, val := range src {
		if match(val) {
			return val, true
		}
	}
	return *new(T), false
}

func FindAll[T any](src []T, match doc.MatchFunc[T]) []T {
	// We think that there should be a small number of qualified elements
	//So it will be divided by 8
	//That is, in the case of triggering expansion,
	//it will be the same as the original capacity three times at most
	//+1 is to ensure that there is at least one element

	res := make([]T, 0, len(src)>>3+1)
	for _, val := range src {
		if match(val) {
			res = append(res, val)
		}
	}
	return res
}

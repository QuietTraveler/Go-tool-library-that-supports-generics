package slice

import (
	"gtool/internal/slice"
	"gtool/internal/slice/doc"
)

func Find[T any](src []T, matchFunc doc.MatchFunc[T]) (T, bool) {
	return slice.Find[T](src, matchFunc)
}

func FindAll[T any](src []T, match doc.MatchFunc[T]) []T {
	return slice.FindAll[T](src, match)
}

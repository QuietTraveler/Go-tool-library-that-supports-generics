package doc

// Match type, used to provide user-defined functions
type MatchFunc[T any] func(src T) bool

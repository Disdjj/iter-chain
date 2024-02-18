package core

type Iter[T any] interface {
	// Next return the next element in the iteration
	Next() (value *T, ok bool)
}

package slice

func New[T, MapT any]() *Iterator[T, MapT] {
	return &Iterator[T, MapT]{
		inner: &FromIterable[T]{},
	}
}

func From[T any, MapT any](slice []T) *Iterator[T, MapT] {
	if slice == nil {
		slice = []T{}
	}
	return &Iterator[T, MapT]{
		inner: &FromIterable[T]{
			slice: slice,
		},
	}
}

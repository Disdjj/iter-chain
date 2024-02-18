package slice

import "github.com/Disdjj/iter-chain/core"

type FilterIter[T any] struct {
	iterator core.Iter[T]
	filter   func(T) bool
}

func (f *FilterIter[T]) Next() (*T, bool) {
	for {
		if item, ok := f.iterator.Next(); ok {
			if f.filter(*item) {
				return item, true
			}
		} else {
			return nil, false
		}
	}
}

func (i *Iterator[T, any]) Filter(f func(T) bool) *Iterator[T, any] {
	return &Iterator[T, any]{
		inner: &FilterIter[T]{
			iterator: i,
			filter:   f,
		},
	}
}

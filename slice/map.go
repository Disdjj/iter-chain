package slice

import "github.com/Disdjj/iter-chain/core"

type MapIterator[A any, B any] struct {
	iterator core.Iter[A]
	mapFunc  func(*A) *B
}

func (m *MapIterator[A, B]) Next() (*B, bool) {
	for {
		if item, ok := m.iterator.Next(); ok {
			return m.mapFunc(item), true
		} else {
			return nil, false
		}
	}
}

func (i *Iterator[A, B]) Map(f func(*A) *B) *Iterator[B, any] {
	return &Iterator[B, any]{
		inner: &MapIterator[A, B]{
			iterator: i,
			mapFunc:  f,
		},
	}
}

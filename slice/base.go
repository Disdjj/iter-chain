package slice

import "github.com/Disdjj/iter-chain/core"

type FromIterable[T any] struct {
	slice []T  // the slice to iterate
	index int  // the current index
	done  bool // if the iteration is done
}

func (i *FromIterable[T]) Next() (*T, bool) {
	if i.done {
		return nil, false
	}

	if i.index >= len(i.slice) {
		i.done = true
		return nil, false
	}

	v := i.slice[i.index]
	i.index++
	return &v, true
}

// Iterator Ugly Generic type hint, but it works, and it's the only way to do it in Go if you want to use generics and chain method calls.
type Iterator[T any, MapT any] struct {
	inner core.Iter[T]
}

func (i *Iterator[T, _]) Next() (*T, bool) {
	return i.inner.Next()
}

func (i *Iterator[T, _]) Collect() []T {
	var result []T
	for {
		v, ok := i.Next()
		if !ok {
			break
		}
		result = append(result, *v)
	}
	return result
}

func (i *Iterator[T, _]) Count() int {
	count := 0
	for {
		_, ok := i.Next()
		if !ok {
			break
		}
		count++
	}
	return count
}

func (i *Iterator[T, _]) First() *T {
	v, ok := i.Next()
	if !ok {
		return nil
	}
	return v
}

func (i *Iterator[T, _]) Last() *T {
	var result *T
	for {
		v, ok := i.Next()
		if !ok {
			break
		}
		result = v
	}
	return result
}

func (i *Iterator[T, _]) Any() *T {
	return i.First()
}

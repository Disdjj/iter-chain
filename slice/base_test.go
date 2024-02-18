package slice

import (
	"reflect"
	"testing"
)

func TestFromIterable_Next(t *testing.T) {
	type testCase[T any] struct {
		name  string
		i     FromIterable[T]
		want  []T
		want1 bool
	}
	tests := []testCase[int]{
		{
			name: "only one",
			i: FromIterable[int]{
				slice: []int{1},
				index: 0,
				done:  false,
			},
			want:  []int{1},
			want1: false,
		},
		{
			name: "has many",
			i: FromIterable[int]{
				slice: []int{1, 2, 3, 4},
				index: 0,
				done:  false,
			},
			want:  []int{1, 2, 3, 4},
			want1: false,
		},
		{
			name: "empty",
			i: FromIterable[int]{
				slice: []int{},
				index: 0,
				done:  false,
			},
			want:  []int{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				for _, v := range tt.want {
					got, _ := tt.i.Next()
					if !reflect.DeepEqual(*got, v) {
						t.Errorf("Next() got = %v, want %v", *got, v)
					}
				}
			},
		)
	}
}

func TestIterator_Any(t *testing.T) {
	type testCase[T any, MapT any] struct {
		name string
		i    *Iterator[T, MapT]
		want *T
	}
	var (
		one   = 1
		two   = 2
		three = 3
	)
	tests := []testCase[int, any]{
		{
			name: "",
			i:    From[int, any]([]int{one, two, three}),
			want: &one,
		},
		{
			name: "",
			i:    From[int, any]([]int{two}),
			want: &two,
		},
		{
			name: "",
			i:    From[int, any]([]int{}),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.i.Any(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Any() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestIterator_Collect(t *testing.T) {
	type testCase[T any, MapT any] struct {
		name string
		i    *Iterator[T, MapT]
		want []T
	}
	var (
		one   = 1
		two   = 2
		three = 3
		empty []int
	)
	tests := []testCase[int, any]{
		{
			name: "",
			i:    From[int, any]([]int{one, two, three}),
			want: []int{one, two, three},
		},
		{
			name: "",
			i:    From[int, any]([]int{two}),
			want: []int{two},
		},
		{
			name: "",
			i:    From[int, any]([]int{}),
			want: empty,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.i.Collect(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Collect() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestIterator_Count(t *testing.T) {
	type testCase[T any, MapT any] struct {
		name string
		i    *Iterator[T, MapT]
		want int
	}
	var (
		one   = 1
		two   = 2
		three = 3
	)
	tests := []testCase[int, any]{
		{
			name: "",
			i:    From[int, any]([]int{one, two, three}),
			want: 3,
		},
		{
			name: "",
			i:    From[int, any]([]int{two}),
			want: 1,
		},
		{
			name: "",
			i:    From[int, any]([]int{}),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.i.Count(); got != tt.want {
					t.Errorf("Count() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestIterator_First(t *testing.T) {
	type testCase[T any, MapT any] struct {
		name string
		i    *Iterator[T, MapT]
		want *T
	}
	var (
		one   = 1
		two   = 2
		three = 3
	)
	tests := []testCase[int, any]{
		{
			name: "",
			i:    From[int, any]([]int{one, two, three}),
			want: &one,
		},
		{
			name: "",
			i:    From[int, any]([]int{two}),
			want: &two,
		},
		{
			name: "",
			i:    From[int, any]([]int{}),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.i.First(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("First() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestIterator_Last(t *testing.T) {
	type testCase[T any, MapT any] struct {
		name string
		i    *Iterator[T, MapT]
		want *T
	}
	var (
		one   = 1
		two   = 2
		three = 3
	)
	tests := []testCase[int, any]{
		{
			name: "",
			i:    From[int, any]([]int{one, two, three}),
			want: &three,
		},
		{
			name: "",
			i:    From[int, any]([]int{two}),
			want: &two,
		},
		{
			name: "",
			i:    From[int, any]([]int{}),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.i.Last(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Last() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

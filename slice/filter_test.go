package slice

import (
	"reflect"
	"testing"
)

func TestFilterIter_Next(t *testing.T) {
	type testCase[T any] struct {
		name  string
		f     FilterIter[T]
		want  *T
		want1 bool
	}
	var two = 2
	tests := []testCase[int]{
		{
			name: "Next with non-empty iterator and valid filter",
			f: FilterIter[int]{
				iterator: From[int, any]([]int{1, two, 3, 4, 5}),
				filter:   func(n int) bool { return n%2 == 0 },
			},
			want:  &two,
			want1: true,
		},
		{
			name: "Next with empty iterator",
			f: FilterIter[int]{
				iterator: From[int, any]([]int{}),
				filter:   func(n int) bool { return n%2 == 0 },
			},
			want:  nil,
			want1: false,
		},
		{
			name: "Next with non-empty iterator and no valid filter",
			f: FilterIter[int]{
				iterator: From[int, any]([]int{1, 3, 5}),
				filter:   func(n int) bool { return n%2 == 0 },
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, got1 := tt.f.Next()
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Next() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("Next() got1 = %v, want %v", got1, tt.want1)
				}
			},
		)
	}

	f := FilterIter[int]{
		iterator: From[int, any]([]int{1, two, 3, 4, 5}),
		filter:   func(n int) bool { return n%2 == 0 },
	}
	for {
		if v, ok := f.Next(); ok {
			println(*v)
		} else {
			break
		}
	}
}

//
// func TestIterator_Filter(t *testing.T) {
// 	type args[T any] struct {
// 		f func(T) bool
// 	}
// 	type testCase[T any, MapT any] struct {
// 		name string
// 		i    Iterator[T, MapT]
// 		args args[T]
// 		want *Iterator[T, MapT]
// 	}
// 	tests := []testCase[int, int]{
// 		{
// 			name: "Filter with non-empty iterator and valid filter",
// 			i:    From[int, any]([]int{1, 2, 3, 4, 5}),
// 			args: args[int]{f: func(n int) bool { return n%2 == 0 }},
// 			want: From[int, any]([]int{2, 4}),
// 		},
// 		{
// 			name: "Filter with empty iterator",
// 			i:    NewIterator([]int{}),
// 			args: args[int]{f: func(n int) bool { return n%2 == 0 }},
// 			want: NewIterator([]int{}),
// 		},
// 		{
// 			name: "Filter with non-empty iterator and no valid filter",
// 			i:    NewIterator([]int{1, 3, 5}),
// 			args: args[int]{f: func(n int) bool { return n%2 == 0 }},
// 			want: NewIterator([]int{}),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(
// 			tt.name, func(t *testing.T) {
// 				if got := tt.i.Filter(tt.args.f); !reflect.DeepEqual(got, tt.want) {
// 					t.Errorf("Filter() = %v, want %v", got, tt.want)
// 				}
// 			},
// 		)
// 	}
// }

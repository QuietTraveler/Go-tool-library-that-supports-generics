package slice

import (
	"github.com/stretchr/testify/assert" // testify
	"gtool/internal/errs"
	"testing"
)

func TestAdd(t *testing.T) {
	type args[T any] struct {
		src     []T
		element T
		index   int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "Add to beginning of non-empty slice",
			args: args[int]{
				src:     []int{1, 2, 3},
				element: 0,
				index:   0,
			},
			want:    []int{0, 1, 2, 3},
			wantErr: nil,
		},
		{
			name: "Add to middle of non-empty slice",
			args: args[int]{
				src:     []int{1, 2, 3},
				element: 4,
				index:   1,
			},
			want:    []int{1, 4, 2, 3},
			wantErr: nil,
		},
		{
			name: "Add to end of non-empty slice",
			args: args[int]{
				src:     []int{1, 2, 3},
				element: 4,
				index:   3,
			},
			want:    []int{1, 2, 3, 4},
			wantErr: nil,
		},
		{
			name: "Add to empty slice",
			args: args[int]{
				src:     []int{},
				element: 1,
				index:   0,
			},
			want:    []int{1},
			wantErr: nil,
		},
		{
			name: "Index out of range (too large)",
			args: args[int]{
				src:     []int{1, 2, 3},
				element: 4,
				index:   4,
			},
			want:    nil,
			wantErr: errs.NewErrIndexOutOfRange(3, 4),
		},
		{
			name: "Index out of range (negative)",
			args: args[int]{
				src:     []int{1, 2, 3},
				element: 4,
				index:   -1,
			},
			want:    nil,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name: "Add to slice with one element",
			args: args[int]{
				src:     []int{1},
				element: 2,
				index:   1,
			},
			want:    []int{1, 2},
			wantErr: nil,
		},
		{
			name: "Add zero to non-empty slice",
			args: args[int]{
				src:     []int{1, 2, 3},
				element: 0,
				index:   1,
			},
			want:    []int{1, 0, 2, 3},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Add(tt.args.src, tt.args.element, tt.args.index)
			assert.Equal(t, tt.wantErr, err, "Add() error: %v, wantErr: %v", err, tt.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, res, "Add() want: %v, res: %v", tt.want, res)
		})
	}
}

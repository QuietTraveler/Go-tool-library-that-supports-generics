package slice

import (
	"basic-go/gtool/internal/errs"
	"github.com/stretchr/testify/assert" // testify
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
			name: "index 0",
			args: args[int]{
				src:     []int{123, 100},
				element: 233,
				index:   0,
			},
			want: []int{233, 123, 100},
		},
		{
			name: "index middle",
			args: args[int]{
				src:     []int{123, 124, 125},
				element: 233,
				index:   1,
			},
			want: []int{123, 233, 124, 125},
		},
		{
			name: "index out of range",
			args: args[int]{
				src:     []int{123, 100},
				element: 233,
				index:   12,
			},
			wantErr: errs.NewErrIndexOutOfRange(2, 12),
		},
		{
			name: "index less than 0",
			args: args[int]{
				src:     []int{123, 100},
				element: 233,
				index:   -1,
			},
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name: "index last",
			args: args[int]{
				src:     []int{123, 100, 101, 102, 102, 102},
				element: 233,
				index:   5,
			},
			want: []int{123, 100, 101, 102, 102, 233, 102},
		},
		{
			name: "append on last",
			args: args[int]{
				src:     []int{123, 100, 101, 102, 102, 102},
				element: 233,
				index:   6,
			},
			want: []int{123, 100, 101, 102, 102, 102, 233},
		},
		{
			name: "index out of range",
			args: args[int]{
				src:     []int{123, 100, 101, 102, 102, 102},
				element: 233,
				index:   7,
			},
			wantErr: errs.NewErrIndexOutOfRange(6, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Add(tt.args.src, tt.args.element, tt.args.index)
			assert.Equal(t, err, tt.wantErr, "Add() error: %v, wantErr: %v")
			if err != nil {
				return
			}
			assert.Equal(t, res, tt.want, "Add() res: %v, want: %v")
		})
	}
}

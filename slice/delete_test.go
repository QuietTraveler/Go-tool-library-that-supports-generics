package slice

import (
	"basic-go/gtool/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDelete test
func TestDelete(t *testing.T) {
	type args[T any] struct {
		src   []T
		index int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "index -1",
			args: args[int]{
				src:   []int{123, 100},
				index: -1,
			},
			want:    []int{100},
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name: "index 0",
			args: args[int]{
				src:   []int{123, 100},
				index: 0,
			},
			want:    []int{100},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _, err := Delete(tt.args.src, tt.args.index)
			assert.Equal(t, tt.wantErr, err, "Delete() error: %v, wantErr: %v")
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, res, "Delete() want: %v, res: %v")
		})
	}
}

// TestFilterDelete test
func TestFilterDelete(t *testing.T) {
	type args[T any] struct {
		src             []T
		deleteCondition func(index int, src T) bool
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "Delete even numbers",
			args: args[int]{
				src: []int{1, 2, 3, 4, 5, 6},
				deleteCondition: func(index int, src int) bool {
					return src%2 == 0
				},
			},
			want:    []int{1, 3, 5},
			wantErr: nil,
		},
		{
			name: "Delete odd numbers",
			args: args[int]{
				src: []int{1, 2, 3, 4, 5, 6},
				deleteCondition: func(index int, src int) bool {
					return src%2 != 0
				},
			},
			want:    []int{2, 4, 6},
			wantErr: nil,
		},
		{
			name: "Delete numbers greater than 3",
			args: args[int]{
				src: []int{1, 2, 3, 4, 5, 6},
				deleteCondition: func(index int, src int) bool {
					return src > 3
				},
			},
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name: "Delete all elements",
			args: args[int]{
				src: []int{1, 2, 3, 4, 5},
				deleteCondition: func(index int, src int) bool {
					return true
				},
			},
			want:    []int{},
			wantErr: nil,
		},
		{
			name: "Delete no elements",
			args: args[int]{
				src: []int{1, 2, 3, 4, 5},
				deleteCondition: func(index int, src int) bool {
					return false
				},
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name: "Delete based on index (even indices)",
			args: args[int]{
				src: []int{1, 2, 3, 4, 5},
				deleteCondition: func(index int, src int) bool {
					return index%2 == 0
				},
			},
			want:    []int{2, 4},
			wantErr: nil,
		},
		{
			name: "Empty slice",
			args: args[int]{
				src: []int{},
				deleteCondition: func(index int, src int) bool {
					return true
				},
			},
			want:    []int{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := FilterDelete(tt.args.src, tt.args.deleteCondition)
			assert.Equal(t, tt.want, res, "FilterDelete() want: %v, res: %v")
		})
	}
}

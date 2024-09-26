package slice

import (
	"basic-go/gtool/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	type args[T any] struct {
		src   []T
		index int
	}
	type testCase[T any] struct {
		name       string
		args       args[T]
		want       []T
		wantTarget T
		wantErr    error
	}
	tests := []testCase[int]{
		{
			name: "Delete from middle",
			args: args[int]{
				src:   []int{1, 2, 3, 4, 5},
				index: 2,
			},
			want:       []int{1, 2, 4, 5},
			wantTarget: 3,
			wantErr:    nil,
		},
		{
			name: "Delete first element",
			args: args[int]{
				src:   []int{1, 2, 3},
				index: 0,
			},
			want:       []int{2, 3},
			wantTarget: 1,
			wantErr:    nil,
		},
		{
			name: "Delete last element",
			args: args[int]{
				src:   []int{1, 2, 3},
				index: 2,
			},
			want:       []int{1, 2},
			wantTarget: 3,
			wantErr:    nil,
		},
		{
			name: "Delete from empty slice",
			args: args[int]{
				src:   []int{},
				index: 0,
			},
			want:       nil,
			wantTarget: 0,
			wantErr:    errs.NewErrIndexOutOfRange(0, 0),
		},
		{
			name: "Index out of range (negative)",
			args: args[int]{
				src:   []int{1, 2, 3},
				index: -1,
			},
			want:       nil,
			wantTarget: 0,
			wantErr:    errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name: "Index out of range (too large)",
			args: args[int]{
				src:   []int{1, 2, 3},
				index: 3,
			},
			want:       nil,
			wantTarget: 0,
			wantErr:    errs.NewErrIndexOutOfRange(3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, target, err := Delete(tt.args.src, tt.args.index)
			assert.Equal(t, err, tt.wantErr, "Delete() error: %v, wantErr: %v")
			if err != nil {
				return
			}
			assert.Equal(t, target, tt.wantTarget, "Delete() target: %v, want: %v")
			assert.Equal(t, res, tt.want, "Add() res: %v, want: %v")
		})
	}
}

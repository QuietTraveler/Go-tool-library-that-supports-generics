package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShrink(t *testing.T) {
	testCases := []struct {
		name        string
		originCap   int
		enqueueLoop int
		expectCap   int
	}{
		{
			name:        "Less than minimum capacity",
			originCap:   32,
			enqueueLoop: 6,
			expectCap:   32,
		},
		{
			name:        "Less than maximum capacity, less than 1/4",
			originCap:   1000,
			enqueueLoop: 20,
			expectCap:   500,
		},
		{
			name:        "Less than maximum capacity, more than 1/4",
			originCap:   1000,
			enqueueLoop: 400,
			expectCap:   1000,
		},
		{
			name:        "more than maximum capacity，less than half",
			originCap:   3000,
			enqueueLoop: 60,
			expectCap:   1875,
		},
		{
			name:        "more than maximum capacity，more than half",
			originCap:   3000,
			enqueueLoop: 2000,
			expectCap:   3000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := make([]int, 0, tc.originCap)
			for i := 0; i < tc.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = Shrink[int](l)
			assert.Equal(t, tc.expectCap, cap(l), "expectCap: %v, resultCap: %v", tc.expectCap, cap(l))
		})
	}
}

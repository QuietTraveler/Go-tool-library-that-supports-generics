package slice

import (
	"github.com/stretchr/testify/assert"
	"gtool/internal/slice/doc"
	"testing"
)

// MatchFunc is a sample match function for testing
func isEven(n int) bool {
	return n%2 == 0
}

func TestFind(t *testing.T) {
	testCases := []struct {
		name   string
		src    []int
		match  doc.MatchFunc[int]
		expect int
		found  bool
	}{
		{
			name:   "Find first even number",
			src:    []int{1, 3, 4, 5},
			match:  isEven,
			expect: 4,
			found:  true,
		},
		{
			name:   "No even number found",
			src:    []int{1, 3, 5},
			match:  isEven,
			expect: 0, // Default value for int
			found:  false,
		},
		{
			name:   "Empty slice",
			src:    []int{},
			match:  isEven,
			expect: 0, // Default value for int
			found:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, found := Find(tc.src, tc.match)
			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.found, found)
		})
	}
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name   string
		src    []int
		match  doc.MatchFunc[int]
		expect []int
	}{
		{
			name:   "Find all even numbers",
			src:    []int{1, 2, 3, 4, 5, 6},
			match:  isEven,
			expect: []int{2, 4, 6},
		},
		{
			name:   "No even numbers found",
			src:    []int{1, 3, 5},
			match:  isEven,
			expect: []int{},
		},
		{
			name:   "Empty slice",
			src:    []int{},
			match:  isEven,
			expect: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindAll(tc.src, tc.match)
			assert.Equal(t, tc.expect, result)
		})
	}
}

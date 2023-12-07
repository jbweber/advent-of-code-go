package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jbweber/advent-of-code-go/internal"
	"github.com/stretchr/testify/require"
)

func TestPermutations(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		in := []string{"one", "two", "three"}
		out := [][]string{
			{"one", "two", "three"},
			{"one", "three", "two"},
			{"two", "three", "one"},
			{"two", "one", "three"},
			{"three", "one", "two"},
			{"three", "two", "one"},
		}
		result := internal.Permutations(in)
		require.Len(t, result, 6)
		require.ElementsMatch(t, out, result)
	})

	t.Run("int", func(t *testing.T) {
		in := []int{1, 2, 3}
		out := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 3, 1},
			{2, 1, 3},
			{3, 1, 2},
			{3, 2, 1},
		}
		result := internal.Permutations(in)
		require.Len(t, result, 6)
		require.ElementsMatch(t, out, result)
	})

}

func TestFindIntersection(t *testing.T) {
	testCases := []struct {
		name     string
		r1       internal.Range
		r2       internal.Range
		overlap  bool
		rOverlap internal.Range
	}{
		{
			name:     "not in",
			r1:       internal.Range{Start: 79, End: 93},
			r2:       internal.Range{Start: 98, End: 100},
			overlap:  false,
			rOverlap: internal.Range{},
		},
		{
			name:     "all in",
			r1:       internal.Range{Start: 79, End: 93},
			r2:       internal.Range{Start: 59, End: 98},
			overlap:  true,
			rOverlap: internal.Range{Start: 79, End: 93},
		},
		{
			name:     "starts outside, partial",
			r1:       internal.Range{Start: 58, End: 93},
			r2:       internal.Range{Start: 59, End: 98},
			overlap:  true,
			rOverlap: internal.Range{Start: 59, End: 93},
		},
		{
			name:     "ends outside, partial",
			r1:       internal.Range{Start: 79, End: 102},
			r2:       internal.Range{Start: 59, End: 98},
			overlap:  true,
			rOverlap: internal.Range{Start: 79, End: 98},
		},
		{
			name:     "center only, partial",
			r1:       internal.Range{Start: 1, End: 100},
			r2:       internal.Range{Start: 50, End: 75},
			overlap:  true,
			rOverlap: internal.Range{Start: 50, End: 75},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r, ok := internal.FindIntersection(testCase.r1, testCase.r2)

			assert.Equal(t, testCase.overlap, ok)
			assert.Equal(t, testCase.rOverlap, r)
		})
	}
}

func TestFindPartialIntersection(t *testing.T) {
	testCases := []struct {
		name     string
		r1       internal.Range
		r2       internal.Range
		overlap  bool
		rOverlap internal.Range
		pOverlap []internal.Range
	}{
		{
			name:     "not in",
			r1:       internal.Range{Start: 79, End: 93},
			r2:       internal.Range{Start: 98, End: 100},
			overlap:  false,
			rOverlap: internal.Range{},
			pOverlap: nil,
		},
		{
			name:     "all in",
			r1:       internal.Range{Start: 79, End: 93},
			r2:       internal.Range{Start: 59, End: 98},
			overlap:  true,
			rOverlap: internal.Range{Start: 79, End: 93},
			pOverlap: nil,
		},
		{
			name:     "starts outside, partial",
			r1:       internal.Range{Start: 58, End: 93},
			r2:       internal.Range{Start: 59, End: 98},
			overlap:  true,
			rOverlap: internal.Range{Start: 59, End: 93},
			pOverlap: []internal.Range{{Start: 58, End: 58}},
		},
		{
			name:     "ends outside, partial",
			r1:       internal.Range{Start: 79, End: 102},
			r2:       internal.Range{Start: 59, End: 98},
			overlap:  true,
			rOverlap: internal.Range{Start: 79, End: 98},
			pOverlap: []internal.Range{{Start: 99, End: 102}},
		},
		{
			name:     "center only, partial",
			r1:       internal.Range{Start: 1, End: 100},
			r2:       internal.Range{Start: 50, End: 75},
			overlap:  true,
			rOverlap: internal.Range{Start: 50, End: 75},
			pOverlap: []internal.Range{{Start: 1, End: 49}, {Start: 76, End: 100}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r, p, ok := internal.FindPartialIntersection(testCase.r1, testCase.r2)

			assert.Equal(t, testCase.overlap, ok)
			assert.Equal(t, testCase.rOverlap, r)
			assert.Equal(t, testCase.pOverlap, p)
		})
	}
}

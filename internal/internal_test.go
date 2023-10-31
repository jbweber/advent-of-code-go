package internal_test

import (
	"testing"

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

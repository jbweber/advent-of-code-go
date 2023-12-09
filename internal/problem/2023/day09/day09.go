package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jbweber/advent-of-code-go/internal"
)

func Execute(input string) (string, string, error) {
	result1, err := part1(input)
	if err != nil {
		return internal.ReturnError(err)
	}

	result2, err := part2(input)
	if err != nil {
		return internal.ReturnError(err)
	}

	return result1, result2, nil
}

func part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		n := parse1(line)
		total += n
	}

	return fmt.Sprint(total), nil
}

func parse1(input string) int {
	inputStr := strings.Fields(input)

	var nums []int

	for _, is := range inputStr {
		n, _ := strconv.Atoi(is)

		nums = append(nums, n)
	}

	seqs := [][]int{nums}
	for {

		var seq []int
		for i := 0; i < len(nums)-1; i++ {
			v := nums[i+1] - nums[i]
			seq = append(seq, v)
		}
		seqs = append(seqs, seq)
		nums = seq

		az := allZero(seq)
		if az {
			break
		}
	}

	reverse(seqs)

	next := 0
	for i := 1; i < len(seqs); i++ {
		v := seqs[i]
		p := v[len(v)-1]

		n := p + next
		next = n
	}

	return next
}

func allZero(in []int) bool {
	for _, i := range in {
		if i != 0 {
			return false
		}
	}

	return true
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func part2(input string) (string, error) {
	return "", nil
}

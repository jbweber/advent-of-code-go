package day07

import (
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
	return "", nil
}

func part2(input string) (string, error) {
	return "", nil
}

func parse(input string) []int {
	lines := strings.Split(input, "\n")
	if len(lines) > 1 {
		panic("too many lines")
	}

	var result []int
	for _, d := range strings.Split(lines[0], ",") {
		days, _ := strconv.Atoi(d)
		result = append(result, days)
	}

	return result
}

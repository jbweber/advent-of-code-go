package day17

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
	//containers := parse(input)

	return "", nil
}

func part2(input string) (string, error) {
	return "", nil
}

func parse(in string) []int {
	var result []int

	for _, i := range strings.Split(in, "\n") {
		v, _ := strconv.Atoi(i)
		result = append(result, v)
	}

	return result
}

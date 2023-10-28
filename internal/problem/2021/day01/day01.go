package day01

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
	increased := 0
	previous := -1

	for _, i := range strings.Split(input, "\n") {
		cur, err := strconv.Atoi(i)
		if err != nil {
			return "", err
		}

		if previous == -1 {
			previous = cur
			continue
		}

		if cur > previous {
			increased += 1
		}
		previous = cur
	}

	return strconv.Itoa(increased), nil
}

func part2(input string) (string, error) {
	inputSlice := make([]int, 0)
	for _, s := range strings.Split(input, "\n") {
		item, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		inputSlice = append(inputSlice, item)
	}

	length := len(inputSlice)
	previous := -1
	increased := 0
	for idx, i := range inputSlice {
		// no more measurements left
		if idx+1 >= length || idx+2 >= length {
			break
		}

		next := i + inputSlice[idx+1] + inputSlice[idx+2]

		if previous == -1 {
			previous = next
			continue
		}

		if next > previous {
			increased += 1
		}

		previous = next
	}

	return strconv.Itoa(increased), nil
}

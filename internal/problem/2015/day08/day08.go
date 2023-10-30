package day08

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
	x, y, _ := parse(input)

	fmt.Println(len(x))
	fmt.Println(len(y))

	return fmt.Sprintf("%d", len(x)-len(y)), nil
}

func part2(input string) (string, error) {
	x, _, y := parse(input)

	fmt.Println(len(x))
	fmt.Println(len(y))

	return fmt.Sprintf("%d", len(y)-len(x)), nil
}

func parse(input string) (string, string, string) {
	lines := strings.Split(input, "\n")
	literal := strings.ReplaceAll(input, "\n", "")

	memory := ""
	for _, line := range lines {
		l, err := strconv.Unquote(line)
		if err != nil {
			panic(err)
		}
		memory += l
	}

	encoded := ""
	for _, line := range lines {
		l := strconv.Quote(line)
		encoded += l
	}

	return literal, memory, encoded
}

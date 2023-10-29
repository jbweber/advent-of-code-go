package day01

import (
	"fmt"
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
	directions := parse(input)

	currentFloor := 0

	for _, direction := range directions {
		switch direction {
		case "(":
			currentFloor += 1
			break
		case ")":
			currentFloor -= 1
		default:
		}
	}

	return fmt.Sprintf("current floor: %d", currentFloor), nil
}

func part2(input string) (string, error) {
	directions := parse(input)

	currentFloor := 0
	entered := -1

	for idx, direction := range directions {
		switch direction {
		case "(":
			currentFloor += 1
			break
		case ")":
			currentFloor -= 1
		default:
		}

		if currentFloor == -1 {
			entered = idx + 1
			break
		}
	}

	return fmt.Sprintf("basement entered at: %d", entered), nil
}

func parse(input string) []string {
	return strings.Split(strings.TrimRight(strings.TrimLeft(input, "\n"), "\n"), "")
}

package day02

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
		id, ok := parse(line)
		if ok {
			total += id
		}
	}
	return fmt.Sprint(total), nil
}

func part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		cb := parse2(line)

		total += cb

	}
	return fmt.Sprint(total), nil
}

func parse(line string) (int, bool) {
	parts1 := strings.Split(line, ":")
	idstr := strings.Fields(parts1[0])[1]

	id, _ := strconv.Atoi(idstr)

	parts2 := strings.Split(parts1[1], ";")

	possible := true
	for _, i := range parts2 {
		cubes := strings.Split(i, ",")

		for _, cube := range cubes {
			cf := strings.Fields(cube)
			num, _ := strconv.Atoi(cf[0])
			if cf[1] == "red" {
				if num > 12 {
					possible = false
					break
				}
			}

			if cf[1] == "green" {
				if num > 13 {
					possible = false
					break
				}
			}

			if cf[1] == "blue" {
				if num > 14 {
					possible = false
					break
				}
			}
		}

		if !possible {
			break
		}
	}

	return id, possible
}

func parse2(line string) int {
	parts1 := strings.Split(line, ":")
	parts2 := strings.Split(parts1[1], ";")

	minRed := -1
	minGreen := -1
	minBlue := -1
	for _, i := range parts2 {
		cubes := strings.Split(i, ",")

		for _, cube := range cubes {
			cf := strings.Fields(cube)
			num, _ := strconv.Atoi(cf[0])
			if cf[1] == "red" {
				if minRed == -1 || num > minRed {
					minRed = num
				}
			}

			if cf[1] == "green" {
				if minGreen == -1 || num > minGreen {
					minGreen = num
				}
			}

			if cf[1] == "blue" {
				if minBlue == -1 || num > minBlue {
					minBlue = num
				}
			}
		}
	}

	cb := minRed * minGreen * minBlue
	return cb
}

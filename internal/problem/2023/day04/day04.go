package day04

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
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		total += parseLine1(line)
	}
	return fmt.Sprint(total), nil
}

func part2(input string) (string, error) {
	return "", nil
}

func parseLine1(line string) int {
	parts1 := strings.Split(line, " | ")
	parts2 := strings.Split(parts1[0], ":")
	//parts3 := strings.Fields(parts2[0])
	//id := parts3[1]
	winners := strings.Fields(parts2[1])
	numbers := strings.Fields(parts1[1])

	unions := map[string]int{}

	for _, v := range winners {
		unions[v] = 0
	}

	count := 0
	for _, v := range numbers {
		_, ok := unions[v]
		if ok {
			count += 1
		}
	}

	total := 0
	for i := 0; i < count; i++ {
		if total == 0 {
			total = 1
			continue
		}
		total = total * 2
	}
	return total
}

type card struct {
}

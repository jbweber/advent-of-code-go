package day06

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
	fishes := parse(input)

	for i := 0; i < 80; i++ {
		add := 0
		for _, fish := range fishes {
			r := fish.Simulate()
			if r {
				add += 1
			}
		}
		for j := 0; j < add; j++ {
			fishes = append(fishes, &lanternfish{days: 8})
		}
	}

	return fmt.Sprintf("%d", len(fishes)), nil
}

func part2(input string) (string, error) {
	fishes := parse2(input)

	for i := 0; i < 256; i++ {
		spawnNew := fishes[0]
		fishes[0] = fishes[1]
		fishes[1] = fishes[2]
		fishes[2] = fishes[3]
		fishes[3] = fishes[4]
		fishes[4] = fishes[5]
		fishes[5] = fishes[6]
		fishes[6] = fishes[7] + spawnNew
		fishes[7] = fishes[8]
		fishes[8] = spawnNew
	}

	total := 0
	for _, i := range fishes {
		total += i
	}

	return fmt.Sprintf("%d", total), nil
}

type lanternfish struct {
	days int
}

func (l *lanternfish) Simulate() bool {
	if l.days == 0 {
		l.days = 6
		return true
	}
	l.days -= 1

	return false
}

func parse(input string) []*lanternfish {
	lines := strings.Split(input, "\n")
	if len(lines) > 1 {
		panic("too many lines")
	}

	var result []*lanternfish
	for _, d := range strings.Split(lines[0], ",") {
		days, _ := strconv.Atoi(d)
		result = append(result, &lanternfish{days: days})
	}

	return result
}

func parse2(input string) [9]int {
	lines := strings.Split(input, "\n")
	if len(lines) > 1 {
		panic("too many lines")
	}

	result := [9]int{0, 0, 0, 0, 0, 0, 0, 0}

	for _, d := range strings.Split(lines[0], ",") {
		days, _ := strconv.Atoi(d)
		result[days] += 1
	}

	return result
}

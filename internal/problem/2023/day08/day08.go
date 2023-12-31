package day08

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

	d, p := parse1(input)

	turns := 0
	next := "AAA"
	for i := 0; i < len(d); i++ {
		turns += 1

		opts := p[next]

		if d[i] == "L" {
			next = opts.Left
		} else {
			next = opts.Right
		}

		if next == "ZZZ" {
			break
		}

		if i+1 == len(d) {
			i = -1
		}
	}

	return fmt.Sprint(turns), nil
}

func part2(input string) (string, error) {
	d, p := parse1(input)

	// find starts
	var nexts []string
	for k, _ := range p {
		if strings.HasSuffix(k, "A") {
			nexts = append(nexts, k)
		}
	}

	var turns []int
	var starts []string

	for k, _ := range p {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}

	for _, start := range starts {

		count := 0
		next := start
		for i := 0; i < len(d); i++ {
			count += 1

			opts := p[next]

			if d[i] == "L" {
				next = opts.Left
			} else {
				next = opts.Right
			}

			if strings.HasSuffix(next, "Z") {
				break
			}

			if i+1 == len(d) {
				i = -1
			}
		}

		turns = append(turns, count)
	}

	r := getLCM(turns)

	return fmt.Sprint(r), nil
}

func parse1(input string) ([]string, map[string]Instruction) {
	lines := strings.Split(input, "\n")

	directions := strings.Split(lines[0], "")

	paths := map[string]Instruction{}
	for _, line := range lines[2:] {
		parts1 := strings.Split(line, " = ")
		start := parts1[0]
		other := strings.TrimLeft(parts1[1], "(")
		other = strings.TrimRight(other, ")")
		parts2 := strings.Split(other, ", ")

		paths[start] = Instruction{Left: parts2[0], Right: parts2[1]}
	}

	return directions, paths
}

type Instruction struct {
	Left  string
	Right string
}

// https://en.wikipedia.org/wiki/Greatest_common_divisor
// Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// https://en.wikipedia.org/wiki/Least_common_multiple
func lcm(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}

	return a * b / gcd(a, b)
}

func getLCM(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}

	return result
}

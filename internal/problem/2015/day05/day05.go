package day05

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
	words := strings.Split(input, "\n")

	count := 0
	for _, word := range words {
		threeVowels := false
		repeats := false
		noBadLetters := false

		vowels := countVowels(word)
		if vowels >= 3 {
			threeVowels = true
		}

		repeats = anyRepeats(word)

		noBadLetters = !anyBanned(word)

		if threeVowels && repeats && noBadLetters {
			count += 1
		}
	}

	return fmt.Sprintf("nice strings: %d", count), nil
}

func part2(input string) (string, error) {

	words := strings.Split(input, "\n")

	count := 0
	for _, word := range words {
		a := false
		b := false

		a = multiRepeats(word)
		b = repeatOneBetween(word)

		if a && b {
			count += 1
		}
	}

	return fmt.Sprintf("nice strings: %d", count), nil
}

func countVowels(input string) int {
	vowels := map[string]int{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0}
	for _, c := range input {
		_, ok := vowels[string(c)]
		if ok {
			vowels[string(c)] += 1
		}
	}

	return vowels["a"] + vowels["e"] + vowels["i"] + vowels["o"] + vowels["u"]
}

func anyRepeats(input string) bool {
	previous := ""
	for _, c := range input {
		curr := string(c)
		if previous == curr {
			return true
		}
		previous = curr
	}

	return false
}

func anyBanned(input string) bool {
	if strings.Contains(input, "ab") {
		return true
	}

	if strings.Contains(input, "cd") {
		return true
	}

	if strings.Contains(input, "pq") {
		return true
	}

	if strings.Contains(input, "xy") {
		return true
	}

	return false
}

func multiRepeats(input string) bool {
	for i := 0; i < len(input); i++ {
		// we're at the end
		if i == len(input)-1 {
			continue
		}

		p1 := string(input[i])
		p2 := string(input[i+1])
		p := p1 + p2

		fi := strings.LastIndex(input, p)

		// match or overlap or no matches
		if fi == i || fi == i+1 || fi == -1 {
			continue
		}

		//
		return true
	}

	return false
}

func repeatOneBetween(input string) bool {
	lngth := len(input)

	for i := 0; i < lngth; i++ {

		// can't get three letters
		if i >= lngth-2 {
			continue
		}

		p1 := string(input[i])
		p2 := string(input[i+1])
		p3 := string(input[i+2])
		p := p1 + p2 + p3

		if p[0] == p[2] {
			return true
		}
	}

	return false
}

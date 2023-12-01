package day01

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

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
	total := 0
	for _, line := range strings.Split(input, "\n") {
		f := 'f'
		l := 'l'
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) == true {
				if f == 'f' {
					f = rune(line[i])
				}
				l = rune(line[i])

			}
		}

		d, _ := strconv.Atoi(string(f) + string(l))
		total += d
	}
	return fmt.Sprint(total), nil
}

var nums = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func part2(input string) (string, error) {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		f := 'f'
		l := 'l'
		for i := 0; i < len(line); i++ {
			for k, v := range nums {
				if strings.HasPrefix(line[i:], k) {
					if f == 'f' {
						f = v
					}
					l = v
					continue
				}
			}
			if unicode.IsDigit(rune(line[i])) == true {
				if f == 'f' {
					f = rune(line[i])
				}
				l = rune(line[i])
			}
		}
		d, _ := strconv.Atoi(string(f) + string(l))
		total += d
	}

	return fmt.Sprint(total), nil
}

package day04

import (
	"crypto/md5"
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
	input = strings.Trim(input, "\n")

	count := 1
	for {
		iv := fmt.Sprintf("%s%d", input, count)
		mv := fmt.Sprintf("%x", md5.Sum([]byte(iv)))
		if strings.HasPrefix(mv, "00000") {
			break
		}

		count += 1
	}

	return fmt.Sprintf("first is %d", count), nil
}

func part2(input string) (string, error) {
	input = strings.Trim(input, "\n")

	count := 1
	for {
		iv := fmt.Sprintf("%s%d", input, count)
		mv := fmt.Sprintf("%x", md5.Sum([]byte(iv)))
		if strings.HasPrefix(mv, "000000") {
			break
		}

		count += 1
	}

	return fmt.Sprintf("first is %d", count), nil
}

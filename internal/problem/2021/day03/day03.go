package day03

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
	// [][]string{}

	var dataTable [][]string

	for _, line := range strings.Split(input, "\n") {
		dataTable = append(dataTable, strings.Split(line, ""))
	}

	size := len(dataTable[0])

	gamma := make([]int, size)
	epsilon := make([]int, size)

	ones := make([]int, size)
	zeroes := make([]int, size)
	for _, data := range dataTable {
		for idx, i := range data {
			if i == "0" {
				zeroes[idx] += 1
			}

			if i == "1" {
				ones[idx] += 1
			}
		}
	}

	for i := 0; i < size; i++ {
		if ones[i] > zeroes[i] {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	var gammaStr, epsilonStr string

	for _, i := range gamma {
		gammaStr = gammaStr + strconv.Itoa(i)
	}

	for _, i := range epsilon {
		epsilonStr = epsilonStr + strconv.Itoa(i)
	}

	gammaR, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		return "", err
	}

	epsilonR, err := strconv.ParseInt(epsilonStr, 2, 32)
	if err != nil {
		return "", err
	}

	result := strconv.Itoa(int(gammaR * epsilonR))

	return result, nil
}

func part2(input string) (string, error) {
	// [][]string{}

	lines := strings.Split(input, "\n")

	var oxygen int64
	var scrubber int64
	var err error

	pos := 0
	zeroCount := 0

	var lines0 []string
	var lines1 []string
	filtered := lines

	for {
		for _, line := range filtered {
			if string(line[pos]) == "0" {
				zeroCount += 1
				lines0 = append(lines0, line)
			} else {
				lines1 = append(lines1, line)
			}
		}

		pos += 1

		if len(lines1) >= len(lines0) {
			filtered = lines1
		} else {
			filtered = lines0
		}

		lines0 = lines0[:0]
		lines1 = lines1[:0]

		if len(filtered) == 1 {
			oxygen, err = strconv.ParseInt(filtered[0], 2, 64)
			if err != nil {
				return "", err
			}
			break
		}
	}

	pos = 0
	zeroCount = 0
	lines0 = lines0[:0]
	lines1 = lines1[:0]
	filtered = lines

	for {
		for _, line := range filtered {
			if string(line[pos]) == "0" {
				zeroCount += 1
				lines0 = append(lines0, line)
			} else {
				lines1 = append(lines1, line)
			}
		}

		pos += 1

		if len(lines1) >= len(lines0) {
			filtered = lines0
		} else {
			filtered = lines1
		}

		lines0 = lines0[:0]
		lines1 = lines1[:0]

		if len(filtered) == 1 {
			scrubber, err = strconv.ParseInt(filtered[0], 2, 64)
			if err != nil {
				return "", err
			}
			break
		}
	}

	return strconv.Itoa(int(scrubber * oxygen)), nil
}

//func Execute(input string) (string, string, error) {
//	result1, err := part1(input)
//	if err != nil {
//		return internal.ReturnError(err)
//	}
//
//	result2, err := part2(input)
//	if err != nil {
//		return internal.ReturnError(err)
//	}
//
//	return result1, result2, nil
//}
//
//func part1(input string) (string, error) {
//	return "", nil
//}
//
//func part2(input string) (string, error) {
//	return "", nil
//}

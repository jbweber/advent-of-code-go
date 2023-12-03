package day03

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
	data := strings.Split(input, "\n")
	total := 0
	for i := 0; i < len(data); i++ {
		adjacent := false
		inDigit := false
		num := ""
		for j := 0; j < len(data[i]); j++ {
			if unicode.IsDigit(rune(data[i][j])) {
				inDigit = true
				num += string(data[i][j])

				// check adjacent
				if !adjacent {

					adjacent = IsAdjacent(data, j, i)
				}
				if j+1 != len(data[i]) {
					continue
				}
			}

			if inDigit == true {
				fmt.Println(num, adjacent)

				if adjacent {
					v, _ := strconv.Atoi(num)
					total += v
				}

				adjacent = false
				inDigit = false
				num = ""
			}
		}
	}

	return fmt.Sprint(total), nil
}

func part2(input string) (string, error) {
	return "", nil
}

func IsAdjacent(data []string, x, y int) bool {
	// x,y+1
	if y < len(data)-1 {
		val := data[y+1][x]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	// x-1,y
	if x > 0 {
		val := data[y][x-1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	// x+1,y
	if x < len(data[y])-1 {
		val := data[y][x+1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	// x,y-1
	if y > 0 {
		val := data[y-1][x]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	// x+1,y+1
	if y < len(data)-1 && x < len(data[y])-1 {
		val := data[y+1][x+1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	// x-1,y+1
	if y < len(data)-1 && x > 0 {
		val := data[y+1][x-1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	// x+1,y-1
	if y > 0 && x < len(data[y])-1 {
		val := data[y-1][x+1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	// x-1,y-1
	if y > 0 && x > 0 {
		val := data[y-1][x-1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			fmt.Println(string(val))
			return true
		}
	}

	return false
}

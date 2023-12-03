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

					adjacent, _, _, _ = IsAdjacent(data, j, i)
				}
				if j+1 != len(data[i]) {
					continue
				}
			}

			if inDigit == true {

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
	data := strings.Split(input, "\n")
	gears := map[string][]int{}
	for i := 0; i < len(data); i++ {
		adjacent := false
		inDigit := false
		gear := false
		gearX := -1
		gearY := -1
		num := ""
		for j := 0; j < len(data[i]); j++ {
			if unicode.IsDigit(rune(data[i][j])) {
				inDigit = true
				num += string(data[i][j])

				// check adjacent
				if !adjacent {
					adjacent, gear, gearX, gearY = IsAdjacent(data, j, i)

				}
				if j+1 != len(data[i]) {
					continue
				}
			}

			if inDigit == true {
				if adjacent && gear {
					v, _ := strconv.Atoi(num)
					si, ok := gears[fmt.Sprintf("%d-%d", gearX, gearY)]
					if !ok {
						gears[fmt.Sprintf("%d-%d", gearX, gearY)] = []int{v}
					} else {
						gears[fmt.Sprintf("%d-%d", gearX, gearY)] = append(si, v)
					}
				}

				adjacent = false
				inDigit = false
				gear = false
				gearX = -1
				gearY = -1
				num = ""
			}
		}
	}

	total := 0
	for _, v := range gears {
		if len(v) == 2 {
			total += (v[0] * v[1])
		}
	}

	return fmt.Sprint(total), nil
}

func IsAdjacent(data []string, x, y int) (bool, bool, int, int) {
	// x,y+1
	if y < len(data)-1 {
		val := data[y+1][x]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x, y + 1
		}
	}

	// x-1,y
	if x > 0 {
		val := data[y][x-1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x - 1, y
		}
	}

	// x+1,y
	if x < len(data[y])-1 {
		val := data[y][x+1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x + 1, y
		}
	}

	// x,y-1
	if y > 0 {
		val := data[y-1][x]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x, y - 1
		}
	}

	// x+1,y+1
	if y < len(data)-1 && x < len(data[y])-1 {
		val := data[y+1][x+1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x + 1, y + 1
		}
	}

	// x-1,y+1
	if y < len(data)-1 && x > 0 {
		val := data[y+1][x-1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x - 1, y + 1
		}
	}

	// x+1,y-1
	if y > 0 && x < len(data[y])-1 {
		val := data[y-1][x+1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x + 1, y - 1
		}
	}

	// x-1,y-1
	if y > 0 && x > 0 {
		val := data[y-1][x-1]
		if !unicode.IsDigit(rune(val)) && val != '.' {
			return true, val == '*', x - 1, y - 1
		}
	}

	return false, false, 0, 0
}

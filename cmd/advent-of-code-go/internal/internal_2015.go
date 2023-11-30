package internal

import (
	"fmt"
	"log"

	year2015day01 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day01"
	year2015day02 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day02"
	year2015day03 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day03"
	year2015day04 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day04"
	year2015day05 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day05"
	year2015day06 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day06"
	year2015day07 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day07"
	year2015day08 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day08"
	year2015day09 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day09"
	year2015day10 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day10"
	year2015day12 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day12"
	year2015day14 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day14"
	year2015day16 "github.com/jbweber/advent-of-code-go/internal/problem/2015/day16"
)

func Execute2015(input string, year, day int) (result1 string, result2 string, err error) {

	switch fmt.Sprintf("%d-%d", year, day) {
	case "2015-1":
		result1, result2, err = year2015day01.Execute(input)
		break
	case "2015-2":
		result1, result2, err = year2015day02.Execute(input)
		break
	case "2015-3":
		result1, result2, err = year2015day03.Execute(input)
		break
	case "2015-4":
		result1, result2, err = year2015day04.Execute(input)
		break
	case "2015-5":
		result1, result2, err = year2015day05.Execute(input)
		break
	case "2015-6":
		result1, result2, err = year2015day06.Execute(input)
		break
	case "2015-7":
		result1, result2, err = year2015day07.Execute(input)
		break
	case "2015-8":
		result1, result2, err = year2015day08.Execute(input)
		break
	case "2015-9":
		result1, result2, err = year2015day09.Execute(input)
		break
	case "2015-10":
		result1, result2, err = year2015day10.Execute(input)
		break
	case "2015-12":
		result1, result2, err = year2015day12.Execute(input)
		break
	case "2015-14":
		result1, result2, err = year2015day14.Execute(input)
		break
	case "2015-16":
		result1, result2, err = year2015day16.Execute(input)
		break
	default:
		log.Fatalf("unknown year / day combo %d-%d", year, day)
	}

	return
}

package internal

import (
	"fmt"
	"log"

	year2023day01 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day01"
	year2023day02 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day02"
	year2023day03 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day03"
	year2023day04 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day04"
	year2023day05 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day05"
	year2023day06 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day06"
	year2023day07 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day07"
	year2023day08 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day08"
	year2023day09 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day09"
	year2023day10 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day10"
	year2023day12 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day12"
	year2023day13 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day13"
	year2023day14 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day14"
	year2023day15 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day15"
	year2023day16 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day16"
	year2023day17 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day17"
	year2023day18 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day18"
	year2023day19 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day19"
	year2023day20 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day20"
	year2023day21 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day21"
	year2023day22 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day22"
	year2023day23 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day23"
	year2023day24 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day24"
	year2023day25 "github.com/jbweber/advent-of-code-go/internal/problem/2023/day25"
)

func Execute2023(input string, year, day int) (result1 string, result2 string, err error) {

	switch fmt.Sprintf("%d-%d", year, day) {
	case "2023-1":
		result1, result2, err = year2023day01.Execute(input)
		break
	case "2023-2":
		result1, result2, err = year2023day02.Execute(input)
		break
	case "2023-3":
		result1, result2, err = year2023day03.Execute(input)
		break
	case "2023-4":
		result1, result2, err = year2023day04.Execute(input)
		break
	case "2023-5":
		result1, result2, err = year2023day05.Execute(input)
		break
	case "2023-6":
		result1, result2, err = year2023day06.Execute(input)
		break
	case "2023-7":
		result1, result2, err = year2023day07.Execute(input)
		break
	case "2023-8":
		result1, result2, err = year2023day08.Execute(input)
		break
	case "2023-9":
		result1, result2, err = year2023day09.Execute(input)
		break
	case "2023-10":
		result1, result2, err = year2023day10.Execute(input)
		break
	case "2023-12":
		result1, result2, err = year2023day12.Execute(input)
		break
	case "2023-13":
		result1, result2, err = year2023day13.Execute(input)
		break
	case "2023-14":
		result1, result2, err = year2023day14.Execute(input)
		break
	case "2023-15":
		result1, result2, err = year2023day15.Execute(input)
		break
	case "2023-16":
		result1, result2, err = year2023day16.Execute(input)
		break
	case "2023-17":
		result1, result2, err = year2023day17.Execute(input)
		break
	case "2023-18":
		result1, result2, err = year2023day18.Execute(input)
		break
	case "2023-19":
		result1, result2, err = year2023day19.Execute(input)
		break
	case "2023-20":
		result1, result2, err = year2023day20.Execute(input)
		break
	case "2023-21":
		result1, result2, err = year2023day21.Execute(input)
		break
	case "2023-22":
		result1, result2, err = year2023day22.Execute(input)
		break
	case "2023-23":
		result1, result2, err = year2023day23.Execute(input)
		break
	case "2023-24":
		result1, result2, err = year2023day24.Execute(input)
		break
	case "2023-25":
		result1, result2, err = year2023day25.Execute(input)
		break
	default:
		log.Fatalf("unknown year / day combo %d-%d", year, day)
	}

	return
}

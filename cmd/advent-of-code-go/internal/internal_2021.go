package internal

import (
	"fmt"
	"log"

	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day01"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day02"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day03"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day04"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day05"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day06"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day07"
)

func Execute2021(input string, year, day int) (result1 string, result2 string, err error) {
	switch fmt.Sprintf("%d-%d", year, day) {
	case "2021-1":
		result1, result2, err = day01.Execute(input)
		break
	case "2021-2":
		result1, result2, err = day02.Execute(input)
		break
	case "2021-3":
		result1, result2, err = day03.Execute(input)
		break
	case "2021-4":
		result1, result2, err = day04.Execute(input)
		break
	case "2021-5":
		result1, result2, err = day05.Execute(input)
		break
	case "2021-6":
		result1, result2, err = day06.Execute(input)
		break
	case "2021-7":
		result1, result2, err = day07.Execute(input)
		break
	default:
		log.Fatalf("unknown year / day combo %d-%d", year, day)
	}

	return
}

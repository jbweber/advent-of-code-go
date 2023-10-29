package main

import (
	"fmt"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day06"
	"log"
	"os"
	"strconv"

	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day01"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day02"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day03"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day04"
	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day05"
	"github.com/jbweber/advent-of-code-go/internal/probleminputs"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("3 arguments required, only received %d", len(os.Args)-1)
	}

	inputType := os.Args[1]
	inputYear := os.Args[2]
	inputDay := os.Args[3]

	year, err := strconv.Atoi(inputYear)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to convert year from string, %w", err))
	}

	day, err := strconv.Atoi(inputDay)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to convert day from string, %w", err))
	}

	var input string
	switch inputType {
	case "input":
		input, err = probleminputs.GetInput(year, day)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "sample":
		input, err = probleminputs.GetSample(year, day)
		if err != nil {
			log.Fatal(err)
		}
		break
	default:
		log.Fatalf("unknown input type %s, must be input or sample\n", inputType)
	}

	execute(input, year, day)
}

func execute(input string, year, day int) {
	var result1, result2 string
	var err error

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
	default:
		log.Fatalf("unknown year / day combo %d-%d", year, day)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("First Result: %s", result1)
	log.Printf("Second Result: %s", result2)
}

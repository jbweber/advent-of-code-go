package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jbweber/advent-of-code-go/cmd/advent-of-code-go/internal"
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

	switch year {
	case 2015:
		result1, result2, err = internal.Execute2015(input, year, day)
		break
	case 2021:
		result1, result2, err = internal.Execute2021(input, year, day)
		break
	case 2023:
		result1, result2, err = internal.Execute2023(input, year, day)
		break
	default:
		err = fmt.Errorf("unknown year / day combo %d-%d", year, day)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("First Result: %s", result1)
	log.Printf("Second Result: %s", result2)
}

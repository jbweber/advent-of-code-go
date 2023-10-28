package main

import (
	"log"

	"github.com/jbweber/advent-of-code-go/internal/problem/2021/day01"
	"github.com/jbweber/advent-of-code-go/internal/probleminputs"
)

func main() {
	input, err := probleminputs.GetInput(2021, 1)
	if err != nil {
		log.Fatal(err)
	}

	result1, result2, err := day01.Execute(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("First Result: %s", result1)
	log.Printf("Second Result: %s", result2)
}

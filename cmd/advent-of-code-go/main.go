package main

import (
	"fmt"

	"github.com/jbweber/advent-of-code-go/internal/probleminputs"
)

func main() {
	file, err := probleminputs.GetSample(2021, 1)
	if err != nil {
		return
	}

	fmt.Print(file)
}

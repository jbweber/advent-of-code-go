package day06

import (
	"fmt"

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
	//times := []int{7, 15, 30}
	//distances := []int{9, 40, 200}
	times := []int{47, 84, 74, 67}
	distances := []int{207, 1394, 1209, 1014}

	wins := 1
	for i := 0; i < len(times); i++ {
		t := times[i]
		d := distances[i]

		count := 0
		for bs := 0; bs < t; bs++ {
			gs := t - bs
			s := bs
			td := gs * s
			if td > d {
				count += 1
			}
		}
		if count > 0 {
			wins *= count
		}
	}

	return fmt.Sprint(wins), nil
}

func part2(input string) (string, error) {
	//times := []int{71530}
	//distances := []int{940200}
	times := []int{47847467}
	distances := []int{207139412091014}

	wins := 1
	for i := 0; i < len(times); i++ {
		t := times[i]
		d := distances[i]

		count := 0
		for bs := 0; bs < t; bs++ {
			gs := t - bs
			s := bs
			td := gs * s
			if td > d {
				count += 1
			}
		}
		if count > 0 {
			wins *= count
		}
	}

	return fmt.Sprint(wins), nil
}

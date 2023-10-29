package day03

import (
	"fmt"
	"strings"

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
	directions := parse(input)

	houses := map[string]int{"0-0": 0}

	x := 0
	y := 0

	for _, direction := range directions {
		switch direction {
		case ">":
			x += 1
		case "<":
			x -= 1
		case "^":
			y += 1
		case "v":
			y -= 1
		default:
		}

		key := fmt.Sprintf("%d-%d", x, y)

		cur, ok := houses[key]
		if !ok {
			cur = 0
		}

		houses[key] = cur + 1
	}

	return fmt.Sprintf("%d houses got a gift", len(houses)), nil
}

func part2(input string) (string, error) {
	directions := parse(input)

	housesSanta := map[string]int{"0-0": 0}
	housesRobot := map[string]int{"0-0": 0}

	xSanta := 0
	ySanta := 0

	xRobot := 0
	yRobot := 0

	actor := 0

	for _, direction := range directions {
		switch direction {
		case ">":
			if actor == 0 {
				xSanta += 1
			} else {
				xRobot += 1
			}
		case "<":
			if actor == 0 {
				xSanta -= 1
			} else {
				xRobot -= 1
			}
		case "^":
			if actor == 0 {
				ySanta += 1
			} else {
				yRobot += 1
			}
		case "v":
			if actor == 0 {
				ySanta -= 1
			} else {
				yRobot -= 1
			}
		default:
		}

		if actor == 0 {
			key := fmt.Sprintf("%d-%d", xSanta, ySanta)
			cur, ok := housesSanta[key]
			if !ok {
				cur = 0
			}

			housesSanta[key] = cur + 1
			actor = 1
		} else {
			key := fmt.Sprintf("%d-%d", xRobot, yRobot)
			cur, ok := housesRobot[key]
			if !ok {
				cur = 0
			}

			housesRobot[key] = cur + 1
			actor = 0
		}
	}

	for k, v := range housesRobot {
		housesSanta[k] = v
	}

	return fmt.Sprintf("%d houses got a gift", len(housesSanta)), nil
}

func parse(input string) []string {
	input = strings.Trim(input, "\n")
	directions := strings.Split(input, "")
	return directions
}
